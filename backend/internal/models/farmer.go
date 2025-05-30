// internal/models/farmer.go
package models

import (
	"errors"
	"time"
)

type Farmer struct {
	ID                    uint      `json:"id" gorm:"primaryKey"`
	FarmerName            string    `json:"farmerName" gorm:"column:name;not null"`
	FederalIdentification string    `json:"federalIdentification" gorm:"unique;not null"`
	Farms                 []Farm    `json:"farms" gorm:"foreignKey:FarmerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func (f *Farmer) Validate() error {
	if f.FarmerName == "" {
		return errors.New("nome do fazendeiro é obrigatório")
	}

	if f.FederalIdentification == "" {
		return errors.New("documento é obrigatório")
	}

	if len(f.FederalIdentification) != 11 && len(f.FederalIdentification) != 14 {
		return errors.New("CPF deve conter 11 dígitos / CNPJ deve conter 14 dígitos")
	}

	return nil
}
