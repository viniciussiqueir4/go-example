package interfaces

import "kyc/models"

type IProcessRepository interface {
	Add(p models.Process) (models.Process, error)
	// Delete(id uint) (models.Process, error)
	// Update(p models.Process) error
	// GetById(id uint) (models.Process, error)
}
