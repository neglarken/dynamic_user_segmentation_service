package repository

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

type Users interface {
	Create(id string) error
}

type Slugs interface {
	Create(title string) error
	Delete(title string) error
}

type SlugsUsers interface {
	Add(title []string, id int) error
	Get(id int) ([]*entity.SlugsUsers, error)
}

type Repository struct {
	Users      *UsersRepository
	Slugs      *SlugsRepository
	SlugsUsers *SlugsUsersRepository
}

func NewRepository(store *store.Store) *Repository {
	return &Repository{
		Users:      NewUsersRepository(store),
		Slugs:      NewSlugsRepository(store),
		SlugsUsers: NewSlugsUsersRepository(store),
	}
}
