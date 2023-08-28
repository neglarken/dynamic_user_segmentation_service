package service

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository"
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

type Service struct {
	Users      Users
	Slugs      Slugs
	SlugsUsers SlugsUsers
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Users:      NewUsersService(*repo.Users),
		Slugs:      NewSlugsService(*repo.Slugs),
		SlugsUsers: NewSlugsUsersService(*repo.SlugsUsers),
	}
}
