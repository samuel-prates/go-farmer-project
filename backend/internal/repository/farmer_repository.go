// internal/repository/farmer_repository.go
package repository

import (
	"github.com/samuel-prates/farm-project/backend/internal/models"
	"gorm.io/gorm"
)

type FarmerRepository struct {
	db *gorm.DB
}

func NewFarmerRepository(db *gorm.DB) *FarmerRepository {
	return &FarmerRepository{db: db}
}

func (r *FarmerRepository) Create(farmer *models.Farmer) (*models.Farmer, error) {
	if err := r.db.Create(farmer).Error; err != nil {
		return nil, err
	}
	return farmer, nil
}

func (r *FarmerRepository) Update(farmer *models.Farmer) (*models.Farmer, error) {
	if err := r.db.Where("farmer_id = ?", farmer.ID).Delete(&models.Farm{}).Error; err != nil {
		return nil, err
	}

	if err := r.db.Save(farmer).Error; err != nil {
		return nil, err
	}
	return farmer, nil
}

func (r *FarmerRepository) Delete(id uint) error {
	return r.db.Delete(&models.Farmer{}, id).Error
}

func (r *FarmerRepository) GetByID(id uint) (*models.Farmer, error) {
	var farmer models.Farmer
	if err := r.db.Preload("Farms.Harvests").Preload("Farms").First(&farmer, id).Error; err != nil {
		return nil, err
	}
	return &farmer, nil
}

func (r *FarmerRepository) GetAll(params models.PaginationParams) ([]models.Farmer, int64, error) {
	var farmers []models.Farmer
	var total int64

	// Count total records
	if err := r.db.Model(&models.Farmer{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (params.Page - 1) * params.Limit
	if err := r.db.Offset(offset).Limit(params.Limit).Preload("Farms.Harvests").Preload("Farms").Find(&farmers).Error; err != nil {
		return nil, 0, err
	}

	return farmers, total, nil
}
