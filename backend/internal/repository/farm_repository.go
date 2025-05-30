// internal/repository/farm_repository.go
package repository

import (
	"github.com/samuel-prates/farm-project/backend/internal/models"
	"gorm.io/gorm"
)

type FarmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) *FarmRepository {
	return &FarmRepository{db: db}
}

func (r *FarmRepository) Create(farm *models.Farm) (*models.Farm, error) {
	if err := r.db.Create(farm).Error; err != nil {
		return nil, err
	}
	return farm, nil
}

func (r *FarmRepository) Update(farm *models.Farm) (*models.Farm, error) {
	if err := r.db.Save(farm).Error; err != nil {
		return nil, err
	}
	return farm, nil
}

func (r *FarmRepository) Delete(id uint) error {
	return r.db.Delete(&models.Farm{}, id).Error
}

func (r *FarmRepository) GetByID(id uint) (*models.Farm, error) {
	var farm models.Farm
	if err := r.db.Preload("Harvests").First(&farm, id).Error; err != nil {
		return nil, err
	}
	return &farm, nil
}

func (r *FarmRepository) GetAll(params models.PaginationParams) ([]models.Farm, int64, error) {
	var farms []models.Farm
	var total int64

	// Count total records
	if err := r.db.Model(&models.Farm{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (params.Page - 1) * params.Limit
	if err := r.db.Offset(offset).Limit(params.Limit).Find(&farms).Error; err != nil {
		return nil, 0, err
	}

	return farms, total, nil
}

// Methods for dashboard
func (r *FarmRepository) Count() (int, error) {
	var count int64
	if err := r.db.Model(&models.Farm{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *FarmRepository) SumTotalArea() (float64, error) {
	var sum float64
	if err := r.db.Model(&models.Farm{}).Select("SUM(total_area)").Scan(&sum).Error; err != nil {
		return 0, err
	}
	return sum, nil
}

func (r *FarmRepository) CountByState() ([]models.StateCount, error) {
	var results []models.StateCount
	if err := r.db.Model(&models.Farm{}).
		Select("state, COUNT(*) as count").
		Group("state").
		Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *FarmRepository) SumAgricultureArea() (float64, error) {
	var sum float64
	if err := r.db.Model(&models.Farm{}).Select("SUM(agriculture_area)").Scan(&sum).Error; err != nil {
		return 0, err
	}
	return sum, nil
}

func (r *FarmRepository) SumVegetationArea() (float64, error) {
	var sum float64
	if err := r.db.Model(&models.Farm{}).Select("SUM(vegetation_area)").Scan(&sum).Error; err != nil {
		return 0, err
	}
	return sum, nil
}
