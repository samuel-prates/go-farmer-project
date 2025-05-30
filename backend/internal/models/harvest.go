// internal/models/harvest.go
package models

import (
	"errors"
	"time"
)

type Harvest struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Year      int       `json:"year" gorm:"not null"`
	Culture   string    `json:"culture" gorm:"not null"`
	FarmID    *uint     `json:"farm_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HarvestCultureCount represents the count of harvests by type
type HarvestCultureCount struct {
	Culture string `json:"culture"`
	Count   int    `json:"count"`
}

func (c *Harvest) Validate() error {
	if c.Year <= 0 {
		return errors.New("ano da safra é obrigatório")
	}

	if c.Culture == "" {
		return errors.New("tipo de cultivo é obrigatório")
	}

	return nil
}
