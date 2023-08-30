package repository

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository/repoerrors"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

type UsersRepository struct {
	store *store.Store
}

func NewUsersRepository(store *store.Store) *UsersRepository {
	return &UsersRepository{store}
}

func (r *UsersRepository) Create(u *entity.Users) error {
	err := r.store.Db.QueryRow("INSERT INTO users (id) values ($1) RETURNING id", &u.Id).Scan(&u.Id)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_pkey\"" {
			return repoerrors.ErrAlreadyExists
		}
		return err
	}
	return nil
}

func (r *UsersRepository) GetNumOfRandom(n int) ([]*entity.Users, error) {
	us := make([]*entity.Users, 0)
	rows, err := r.store.Db.Query("SELECT * FROM (SELECT DISTINCT id FROM users) AS ee ORDER BY RANDOM() LIMIT $1", n)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := &entity.Users{}
		err := rows.Scan(&u.Id)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	if len(us) == 0 {
		return nil, repoerrors.ErrNotFound
	}
	return us, nil
}

func (r *UsersRepository) GetCount() (int, error) {
	var count int
	if err := r.store.Db.QueryRow("SELECT count(*) FROM users").Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
