// pkg/database/database.go
package database

import (
	"github.com/samuel-prates/farm-project/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto Migrate the models
	err = db.AutoMigrate(&models.Farmer{}, &models.Farm{}, &models.Harvest{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
