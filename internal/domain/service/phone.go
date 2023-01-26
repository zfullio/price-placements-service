package service

import (
	"checkphones/internal/domain/entity"
)

type PhoneRepository interface {
	Get(spreadsheetID string) (phones []entity.Phone, err error)
}

type phoneService struct {
	repo PhoneRepository
}

func NewPhoneService(repo PhoneRepository) *phoneService {
	return &phoneService{repo: repo}
}

func (p phoneService) Get(spreadsheetID string) (phones []entity.Phone, err error) {
	return p.repo.Get(spreadsheetID)
}
