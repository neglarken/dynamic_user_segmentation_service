package repository

import (
	"time"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

type Users interface {
	Create(u *entity.Users) error
	GetNumOfRandom(n int) ([]*entity.Users, error)
	GetCount() (int, error)
}

type Slugs interface {
	Create(s *entity.Slugs) error
	Delete(s *entity.Slugs) error
	GetSlugIdBySlugTitle(title string) (*entity.Slugs, error)
	GetById(id int) (*entity.Slugs, error)
}

type SlugsUsers interface {
	Add(su *entity.SlugsUsers) error
	GetSlugIdsByUserId(id int) ([]*entity.SlugsUsers, error)
	Delete(su *entity.SlugsUsers) error
}

type Records interface {
	Create(r *entity.Records) error
	GetToCsv(time time.Time) error
}

type Repository struct {
	Users      *UsersRepository
	Slugs      *SlugsRepository
	SlugsUsers *SlugsUsersRepository
	Records    *RecordsRepository
}

func NewRepository(store *store.Store) *Repository {
	return &Repository{
		Users:      NewUsersRepository(store),
		Slugs:      NewSlugsRepository(store),
		SlugsUsers: NewSlugsUsersRepository(store),
		Records:    NewRecordsRepository(store),
	}
}
