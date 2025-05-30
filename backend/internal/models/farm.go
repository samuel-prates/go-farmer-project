// internal/models/farm.go
package models

import (
	"errors"
	"time"
)

type Farm struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"farmName" gorm:"not null"`
	City            string    `json:"city" gorm:"not null"`
	State           string    `json:"state" gorm:"not null"`
	TotalArea       float64   `json:"totalArea" gorm:"not null"`
	AgricultureArea float64   `json:"arableArea" gorm:"not null"`
	VegetationArea  float64   `json:"vegetationArea" gorm:"not null"`
	FarmerID        *uint     `json:"farmer_id"`
	Harvests        []Harvest `json:"harvests" gorm:"foreignKey:FarmID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (f *Farm) Validate() error {
	if f.Name == "" {
		return errors.New("nome da fazenda é obrigatório")
	}

	if f.City == "" {
		return errors.New("cidade é obrigatória")
	}

	if f.State == "" {
		return errors.New("estado é obrigatório")
	}

	if f.TotalArea <= 0 {
		return errors.New("área total deve ser maior que zero")
	}

	if f.AgricultureArea < 0 {
		return errors.New("área agrícola não pode ser negativa")
	}

	if f.VegetationArea < 0 {
		return errors.New("área de vegetação não pode ser negativa")
	}

	// Validação da soma das áreas
	if f.AgricultureArea+f.VegetationArea != f.TotalArea {
		return errors.New("a soma das áreas agrícola e de vegetação não pode ser maior ou menor que a área total")
	}

	return nil
}

// StateCount represents the count of farms by state
type StateCount struct {
	State string `json:"state"`
	Count int    `json:"count"`
}
