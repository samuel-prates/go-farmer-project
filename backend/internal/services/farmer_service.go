// internal/services/farmer_service.go
package services

import (
	"github.com/samuel-prates/farm-project/backend/internal/models"
	"github.com/samuel-prates/farm-project/backend/internal/repository"
)

type FarmerService struct {
	repo *repository.FarmerRepository
}

func NewFarmerService(repo *repository.FarmerRepository) *FarmerService {
	return &FarmerService{repo: repo}
}

func (s *FarmerService) Create(farmer *models.Farmer) (*models.Farmer, error) {
	return s.repo.Create(farmer)
}

func (s *FarmerService) Update(farmer *models.Farmer) (*models.Farmer, error) {
	return s.repo.Update(farmer)
}

func (s *FarmerService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *FarmerService) GetByID(id uint) (*models.Farmer, error) {
	return s.repo.GetByID(id)
}

func (s *FarmerService) GetAll(params models.PaginationParams) (models.PaginatedResult, error) {
	// Set default values if not provided
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}

	farmers, total, err := s.repo.GetAll(params)
	if err != nil {
		return models.PaginatedResult{}, err
	}

	return models.NewPaginatedResult(farmers, total, params), nil
}
