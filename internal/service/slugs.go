package service

import "github.com/neglarken/dynamic_user_segmentation_service/internal/repository"

type SlugsService struct {
	slugsRepo repository.SlugsRepository
}

func NewSlugsService(slugsRepo repository.SlugsRepository) *SlugsService {
	return &SlugsService{slugsRepo: slugsRepo}
}

func (s *SlugsService) Create(title string) error {
	err := s.slugsRepo.Create(title)
	if err != nil {
		return err
	}
	return nil
}

func (s *SlugsService) Delete(title string) error {
	err := s.slugsRepo.Create(title)
	if err != nil {
		return err
	}
	return nil
}
