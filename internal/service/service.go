package service

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository"
)

type Users interface {
	Create(u *entity.Users) error
}

type Slugs interface {
	Create(slugs *entity.Slugs) error
	Delete(slugs *entity.Slugs) error
	GetSlugIdBySlugTitle(title string) (*entity.Slugs, error)
	GetTitleById(id int) (*entity.Slugs, error)
}

type SlugsUsers interface {
	Add(su *entity.SlugsUsers) error
	GetSlugIdsByUserId(id int) ([]*entity.SlugsUsers, error)
	Delete(su *entity.SlugsUsers) error
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
