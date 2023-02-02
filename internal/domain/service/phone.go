package service

import (
	"price-placements-service/internal/domain/entity"
)

type PhoneRepository interface {
	Get(spreadsheetID string) (phones []entity.Phone, err error)
}

type PhoneService struct {
	repo PhoneRepository
}

func NewPhoneService(repo PhoneRepository) *PhoneService {
	return &PhoneService{repo: repo}
}

func (p PhoneService) Get(spreadsheetID string) (phones []entity.Phone, err error) {
	return p.repo.Get(spreadsheetID)
}
