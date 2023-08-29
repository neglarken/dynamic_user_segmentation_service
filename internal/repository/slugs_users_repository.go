package repository

import (
	"database/sql"
	"fmt"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository/repoerrors"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

type SlugsUsersRepository struct {
	store *store.Store
}

func NewSlugsUsersRepository(store *store.Store) *SlugsUsersRepository {
	return &SlugsUsersRepository{store}
}

func (r *SlugsUsersRepository) Add(su *entity.SlugsUsers) error {
	err := r.store.Db.QueryRow(
		"INSERT INTO slugs_users (slug_id, user_id) values ($1, $2) RETURNING slug_id, user_id",
		&su.SlugId,
		&su.UserId,
	).Scan(
		&su.SlugId,
		&su.UserId,
	)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"slugs_users_pkey\"" {
			fmt.Println(err)
			return repoerrors.ErrAlreadyExists
		}
		return err
	}
	return nil
}

func (r *SlugsUsersRepository) Delete(su *entity.SlugsUsers) error {
	err := r.store.Db.QueryRow(
		"DELETE FROM slugs_users WHERE slug_id = $1 AND user_id = $2 returning slug_id",
		&su.SlugId,
		&su.UserId,
	).Scan(&su.SlugId)
	if err != nil {
		if err == sql.ErrNoRows {
			return repoerrors.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *SlugsUsersRepository) GetSlugIdsByUserId(id int) ([]*entity.SlugsUsers, error) {
	sus := make([]*entity.SlugsUsers, 0)
	rows, err := r.store.Db.Query("SELECT slug_id, user_id FROM slugs_users WHERE user_id = $1", id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		su := &entity.SlugsUsers{}
		err := rows.Scan(&su.SlugId, &su.UserId)
		if err != nil {
			return nil, err
		}
		sus = append(sus, su)
	}
	if len(sus) == 0 {
		return nil, repoerrors.ErrNotFound
	}
	return sus, nil
}
