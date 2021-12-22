package migrations

import (
	"kyc/models"

	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {

	db.AutoMigrate(models.Process{},
		models.Address{},
		models.Personal{})
}
