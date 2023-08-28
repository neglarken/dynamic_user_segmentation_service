package repository

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

type UsersRepository struct {
	store *store.Store
}

func NewUsersRepository(store *store.Store) *UsersRepository {
	return &UsersRepository{store}
}

func (r *UsersRepository) Create(id string) error {
	return nil
}
