// internal/repository/harvest_repository.go
package repository

import (
	"github.com/samuel-prates/farm-project/backend/internal/models"
	"gorm.io/gorm"
)

type HarvestRepository struct {
	db *gorm.DB
}

func NewHarvestRepository(db *gorm.DB) *HarvestRepository {
	return &HarvestRepository{db: db}
}

func (r *HarvestRepository) Create(harvest *models.Harvest) (*models.Harvest, error) {
	if err := r.db.Create(harvest).Error; err != nil {
		return nil, err
	}
	return harvest, nil
}

func (r *HarvestRepository) Update(harvest *models.Harvest) (*models.Harvest, error) {
	if err := r.db.Save(harvest).Error; err != nil {
		return nil, err
	}
	return harvest, nil
}

func (r *HarvestRepository) Delete(id uint) error {
	return r.db.Delete(&models.Harvest{}, id).Error
}

func (r *HarvestRepository) GetByID(id uint) (*models.Harvest, error) {
	var harvest models.Harvest
	if err := r.db.First(&harvest, id).Error; err != nil {
		return nil, err
	}
	return &harvest, nil
}

func (r *HarvestRepository) GetAll() ([]models.Harvest, error) {
	var harvests []models.Harvest
	if err := r.db.Find(&harvests).Error; err != nil {
		return nil, err
	}
	return harvests, nil
}

// Method for dashboard
func (r *HarvestRepository) CountByType() ([]models.HarvestCultureCount, error) {
	var results []models.HarvestCultureCount
	if err := r.db.Model(&models.Harvest{}).
		Select("type, COUNT(*) as count").
		Group("type").
		Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
