package postgres

import (
	"kyc/database"
	"kyc/models"
)

type ProcessRepository struct {
}

func (s *ProcessRepository) Add(p models.Process) (models.Process, error) {
	db := database.GetDatabase()
	err := db.Create(&p).Error

	if err != nil {
		return models.Process{}, err
	}

	return p, nil
}
