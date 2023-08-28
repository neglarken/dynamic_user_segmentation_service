package repository

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

type SlugsUsersRepository struct {
	store *store.Store
}

func NewSlugsUsersRepository(store *store.Store) *SlugsUsersRepository {
	return &SlugsUsersRepository{store}
}

func (r *SlugsUsersRepository) Add(title []string, id int) error {
	return nil
}

func (r *SlugsUsersRepository) Get(id int) ([]*entity.SlugsUsers, error) {
	return make([]*entity.SlugsUsers, 1), nil
}
