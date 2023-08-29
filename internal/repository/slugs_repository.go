package repository

import (
	"database/sql"
	"fmt"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository/repoerrors"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

type SlugsRepository struct {
	store *store.Store
}

func NewSlugsRepository(store *store.Store) *SlugsRepository {
	return &SlugsRepository{store}
}

func (r *SlugsRepository) Create(s *entity.Slugs) error {
	err := r.store.Db.QueryRow(
		"INSERT INTO slugs (title) values ($1) RETURNING id, title",
		&s.Title,
	).Scan(&s.Id, &s.Title)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "pq: duplicate key value violates unique constraint \"slugs_title_key\"" {
			return repoerrors.ErrAlreadyExists
		}
		return err
	}
	return nil
}

func (r *SlugsRepository) GetSlugIdBySlugTitle(title string) (*entity.Slugs, error) {
	slug := &entity.Slugs{}
	if err := r.store.Db.QueryRow("SELECT id FROM slugs WHERE title = $1", title).Scan(&slug.Id); err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.ErrNotFound
		}
		return nil, err
	}
	return slug, nil
}

func (r *SlugsRepository) GetById(id int) (*entity.Slugs, error) {
	slug := &entity.Slugs{}
	if err := r.store.Db.QueryRow("SELECT id, title FROM slugs WHERE id = $1", id).Scan(&slug.Id, &slug.Title); err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.ErrNotFound
		}
		return nil, err
	}
	return slug, nil
}

func (r *SlugsRepository) Delete(s *entity.Slugs) error {
	err := r.store.Db.QueryRow("DELETE FROM slugs WHERE title = $1", s.Title).Scan(&s.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return repoerrors.ErrNotFound
		}
		return err
	}
	return nil
}
