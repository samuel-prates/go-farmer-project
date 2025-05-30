// internal/api/handlers/adapters.go
package handlers

import (
	"github.com/samuel-prates/farm-project/backend/internal/models"
	"github.com/samuel-prates/farm-project/backend/internal/services"
)

// ServiceAdapter adapts the real services to our interfaces for testing
type ServiceAdapter struct {
	FarmerService    *services.FarmerService
	DashboardService *services.DashboardService
}

// NewServiceAdapter creates a new ServiceAdapter
func NewServiceAdapter(farmerService *services.FarmerService, dashboardService *services.DashboardService) *ServiceAdapter {
	return &ServiceAdapter{
		FarmerService:    farmerService,
		DashboardService: dashboardService,
	}
}

// FarmerServiceAdapter adapts the real FarmerService to our FarmerServiceInterface
type FarmerServiceAdapter struct {
	service *services.FarmerService
}

// NewFarmerServiceAdapter creates a new FarmerServiceAdapter
func NewFarmerServiceAdapter(service *services.FarmerService) FarmerServiceInterface {
	return &FarmerServiceAdapter{service: service}
}

// Create implements FarmerServiceInterface
func (a *FarmerServiceAdapter) Create(farmer *models.Farmer) (*models.Farmer, error) {
	return a.service.Create(farmer)
}

// Update implements FarmerServiceInterface
func (a *FarmerServiceAdapter) Update(farmer *models.Farmer) (*models.Farmer, error) {
	return a.service.Update(farmer)
}

// Delete implements FarmerServiceInterface
func (a *FarmerServiceAdapter) Delete(id uint) error {
	return a.service.Delete(id)
}

// GetByID implements FarmerServiceInterface
func (a *FarmerServiceAdapter) GetByID(id uint) (*models.Farmer, error) {
	return a.service.GetByID(id)
}

// GetAll implements FarmerServiceInterface
func (a *FarmerServiceAdapter) GetAll(params models.PaginationParams) (models.PaginatedResult, error) {
	return a.service.GetAll(params)
}

// DashboardServiceAdapter adapts the real DashboardService to our DashboardServiceInterface
type DashboardServiceAdapter struct {
	service *services.DashboardService
}

// NewDashboardServiceAdapter creates a new DashboardServiceAdapter
func NewDashboardServiceAdapter(service *services.DashboardService) DashboardServiceInterface {
	return &DashboardServiceAdapter{service: service}
}

// GetDashboardData implements DashboardServiceInterface
func (a *DashboardServiceAdapter) GetDashboardData() (*DashboardData, error) {
	data, err := a.service.GetDashboardData()
	if err != nil {
		return nil, err
	}
	return &DashboardData{
		TotalFarms: data.TotalFarms,
		TotalArea:  data.TotalArea,
	}, nil
}

// GetFarmsByState implements DashboardServiceInterface
func (a *DashboardServiceAdapter) GetFarmsByState() ([]models.StateCount, error) {
	return a.service.GetFarmsByState()
}

// GetHarvestTypes implements DashboardServiceInterface
func (a *DashboardServiceAdapter) GetHarvestTypes() ([]models.HarvestCultureCount, error) {
	return a.service.GetHarvestTypes()
}

// GetAreaDistribution implements DashboardServiceInterface
func (a *DashboardServiceAdapter) GetAreaDistribution() (*AreaDistribution, error) {
	data, err := a.service.GetAreaDistribution()
	if err != nil {
		return nil, err
	}
	return &AreaDistribution{
		AgricultureArea: data.AgricultureArea,
		VegetationArea:  data.VegetationArea,
	}, nil
}
