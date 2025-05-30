// internal/services/dashboard_service.go
package services

import (
	"github.com/samuel-prates/farm-project/backend/internal/models"
	"github.com/samuel-prates/farm-project/backend/internal/repository"
)

type DashboardData struct {
	TotalFarms int     `json:"totalFarms"`
	TotalArea  float64 `json:"totalArea"`
}

type AreaDistribution struct {
	AgricultureArea float64 `json:"arableArea"`
	VegetationArea  float64 `json:"vegetationArea"`
}

type DashboardService struct {
	farmRepo    *repository.FarmRepository
	harvestRepo *repository.HarvestRepository
}

func NewDashboardService(farmRepo *repository.FarmRepository, harvestRepo *repository.HarvestRepository) *DashboardService {
	return &DashboardService{
		farmRepo:    farmRepo,
		harvestRepo: harvestRepo,
	}
}

func (s *DashboardService) GetDashboardData() (*DashboardData, error) {
	totalFarms, err := s.farmRepo.Count()
	if err != nil {
		return nil, err
	}

	totalArea, err := s.farmRepo.SumTotalArea()
	if err != nil {
		return nil, err
	}

	return &DashboardData{
		TotalFarms: totalFarms,
		TotalArea:  totalArea,
	}, nil
}

func (s *DashboardService) GetFarmsByState() ([]models.StateCount, error) {
	return s.farmRepo.CountByState()
}

func (s *DashboardService) GetHarvestTypes() ([]models.HarvestCultureCount, error) {
	return s.harvestRepo.CountByType()
}

func (s *DashboardService) GetAreaDistribution() (*AreaDistribution, error) {
	agricultureArea, err := s.farmRepo.SumAgricultureArea()
	if err != nil {
		return nil, err
	}

	vegetationArea, err := s.farmRepo.SumVegetationArea()
	if err != nil {
		return nil, err
	}

	return &AreaDistribution{
		AgricultureArea: agricultureArea,
		VegetationArea:  vegetationArea,
	}, nil
}
