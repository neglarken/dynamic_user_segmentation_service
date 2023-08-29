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
