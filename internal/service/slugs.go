package service

import (
	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository"
)

type SlugsService struct {
	slugsRepo repository.SlugsRepository
}

func NewSlugsService(slugsRepo repository.SlugsRepository) *SlugsService {
	return &SlugsService{slugsRepo: slugsRepo}
}

func (s *SlugsService) Create(slugs *entity.Slugs) error {
	err := s.slugsRepo.Create(slugs)
	if err != nil {
		return err
	}
	return nil
}

func (s *SlugsService) Delete(slugs *entity.Slugs) error {
	err := s.slugsRepo.Delete(slugs)
	if err != nil {
		return err
	}
	return nil
}

func (s *SlugsService) GetSlugIdBySlugTitle(title string) (*entity.Slugs, error) {
	su, err := s.slugsRepo.GetSlugIdBySlugTitle(title)
	if err != nil {
		return nil, err
	}
	return su, nil
}

func (s *SlugsService) GetTitleById(id int) (*entity.Slugs, error) {
	su, err := s.slugsRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return su, nil
}
