package service

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository"
)

type UsersService struct {
	usersRepo repository.UsersRepository
}

func NewUsersService(usersRepo repository.UsersRepository) *UsersService {
	return &UsersService{usersRepo: usersRepo}
}

func (s *UsersService) Create(u *entity.Users) error {
	err := s.usersRepo.Create(u)
	if err != nil {
		return err
	}
	return nil
}
