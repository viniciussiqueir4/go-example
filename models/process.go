package models

import (
	"time"

	"gorm.io/gorm"
)

type Process struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	UserId     int            `json:"userId"`
	AgreeAt    time.Time      `json:"agreeAt"`
	PersonalId uint           `json:"personalId"`
	Personal   Personal       `gorm:"foreignKey:PersonalId"`
	AddressId  uint           `json:"addressId"`
	Address    Address        `gorm:"foreignKey:AddressId"`
	CreatedAt  time.Time      `json:"created"`
	UpdatedAt  time.Time      `json:"updated"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func NewProcess(userId int) *Process {
	return &Process{
		UserId: userId,
	}
}
