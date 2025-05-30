// internal/services/farm_service.go
package services

import (
	"github.com/samuel-prates/farm-project/backend/internal/models"
	"github.com/samuel-prates/farm-project/backend/internal/repository"
)

type FarmService struct {
	repo *repository.FarmRepository
}

func NewFarmService(repo *repository.FarmRepository) *FarmService {
	return &FarmService{repo: repo}
}

func (s *FarmService) Create(farm *models.Farm) (*models.Farm, error) {
	return s.repo.Create(farm)
}

func (s *FarmService) Update(farm *models.Farm) (*models.Farm, error) {
	return s.repo.Update(farm)
}

func (s *FarmService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *FarmService) GetByID(id uint) (*models.Farm, error) {
	return s.repo.GetByID(id)
}

func (s *FarmService) GetAll(params models.PaginationParams) (models.PaginatedResult, error) {
	// Set default values if not provided
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}

	farms, total, err := s.repo.GetAll(params)
	if err != nil {
		return models.PaginatedResult{}, err
	}

	return models.NewPaginatedResult(farms, total, params), nil
}
