package service

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository"
)

type SlugsUsersService struct {
	slugsUsersRepo repository.SlugsUsersRepository
}

func NewSlugsUsersService(slugsUsersRepo repository.SlugsUsersRepository) *SlugsUsersService {
	return &SlugsUsersService{slugsUsersRepo: slugsUsersRepo}
}

func (s *SlugsUsersService) Add(title []string, id int) error {
	err := s.slugsUsersRepo.Add(title, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SlugsUsersService) Get(id int) ([]*entity.SlugsUsers, error) {
	slugs, err := s.slugsUsersRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return slugs, nil
}
