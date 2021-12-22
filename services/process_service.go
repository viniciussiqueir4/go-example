package services

import (
	"kyc/dtos"
	"kyc/interfaces"
	"kyc/models"
)

type ProcessService struct {
	repository interfaces.IProcessRepository
}

func (s *ProcessService) CreateProcess(data dtos.ProcessDTO) (models.Process, error) {
	process := models.NewProcess(data.UserId)

	return s.repository.Add(*process)
}

func NewProcessService(r interfaces.IProcessRepository) *ProcessService {
	return &ProcessService{
		repository: r,
	}
}
