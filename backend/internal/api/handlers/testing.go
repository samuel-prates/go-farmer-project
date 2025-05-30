// internal/api/handlers/testing.go
package handlers

import (
	"github.com/samuel-prates/farm-project/backend/internal/models"
)

// DashboardData represents the data returned by the dashboard
type DashboardData struct {
	TotalFarms int     `json:"totalFarms"`
	TotalArea  float64 `json:"totalArea"`
}

// AreaDistribution represents the distribution of areas
type AreaDistribution struct {
	AgricultureArea float64 `json:"arableArea"`
	VegetationArea  float64 `json:"vegetationArea"`
}

// FarmerServiceInterface defines the interface for the FarmerService
// This is used for testing to allow mocking the service
type FarmerServiceInterface interface {
	Create(farmer *models.Farmer) (*models.Farmer, error)
	Update(farmer *models.Farmer) (*models.Farmer, error)
	Delete(id uint) error
	GetByID(id uint) (*models.Farmer, error)
	GetAll(params models.PaginationParams) (models.PaginatedResult, error)
}

// DashboardServiceInterface defines the interface for the DashboardService
// This is used for testing to allow mocking the service
type DashboardServiceInterface interface {
	GetDashboardData() (*DashboardData, error)
	GetFarmsByState() ([]models.StateCount, error)
	GetHarvestTypes() ([]models.HarvestCultureCount, error)
	GetAreaDistribution() (*AreaDistribution, error)
}

// Note: In a real project, we would ensure that the real services implement these interfaces.
// However, for testing purposes, we're using mock implementations directly.
