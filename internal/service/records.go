package service

import (
	"time"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository"
)

type RecordsService struct {
	recordsRepo repository.RecordsRepository
}

func NewRecordsService(recordsRepo repository.RecordsRepository) *RecordsService {
	return &RecordsService{recordsRepo: recordsRepo}
}

func (s *RecordsService) Create(rec *entity.Records) error {
	err := s.recordsRepo.Create(rec)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecordsService) GetToCsv(time time.Time) error {
	err := s.recordsRepo.GetToCsv(time)
	if err != nil {
		return err
	}
	return nil
}
