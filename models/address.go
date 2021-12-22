package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Street          string         `json:"Street"`
	ZipCode         string         `json:"zipCode"`
	StreetAditional string         `json:"streetAditional"`
	Number          string         `json:"number"`
	Neighborhood    string         `json:"neighborhood"`
	City            string         `json:"city"`
	State           string         `json:"state"`
	CreatedAt       time.Time      `json:"created"`
	UpdatedAt       time.Time      `json:"updated"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted"`
}
