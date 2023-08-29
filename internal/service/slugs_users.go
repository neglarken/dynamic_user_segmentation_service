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

func (s *SlugsUsersService) Add(su *entity.SlugsUsers) error {
	err := s.slugsUsersRepo.Add(su)
	if err != nil {
		return err
	}
	return nil
}

func (s *SlugsUsersService) GetSlugIdsByUserId(id int) ([]*entity.SlugsUsers, error) {
	slugs, err := s.slugsUsersRepo.GetSlugIdsByUserId(id)
	if err != nil {
		return nil, err
	}
	return slugs, nil
}

func (s *SlugsUsersService) Delete(su *entity.SlugsUsers) error {
	err := s.slugsUsersRepo.Delete(su)
	if err != nil {
		return err
	}
	return nil
}
