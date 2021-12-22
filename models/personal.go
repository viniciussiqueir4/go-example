package models

import (
	"time"

	"gorm.io/gorm"
)

type Personal struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	FullName   string         `json:"fullName"`
	MotherName string         `json:"motherName"`
	Cpf        string         `json:"cpf"`
	BirthDate  time.Time      `json:"birthDate"`
	SocialName string         `json:"socialName"`
	CreatedAt  time.Time      `json:"created"`
	UpdatedAt  time.Time      `json:"updated"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted"`
}
