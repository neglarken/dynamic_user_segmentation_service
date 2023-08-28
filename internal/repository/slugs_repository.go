package repository

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

type SlugsRepository struct {
	store *store.Store
}

func NewSlugsRepository(store *store.Store) *SlugsRepository {
	return &SlugsRepository{store}
}

func (r *SlugsRepository) Create(title string) error {
	return nil
}

func (r *SlugsRepository) Delete(title string) error {
	return nil
}
