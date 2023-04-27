package service

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
)

type PhoneRepository interface {
	Get(spreadsheetID string, developer string) (phones []entity.Phone, err error)
}

type PhoneService struct {
	repo   PhoneRepository
	logger *zerolog.Logger
}

func NewPhoneService(repo PhoneRepository, logger *zerolog.Logger) *PhoneService {
	serviceLogger := logger.With().Str("service", "phone").Logger()

	return &PhoneService{
		repo:   repo,
		logger: &serviceLogger,
	}
}

func (p PhoneService) Get(spreadsheetID string, developer string) (phones []entity.Phone, err error) {
	p.logger.Trace().Msg("Get")

	phones, err = p.repo.Get(spreadsheetID, developer)
	if err != nil {
		return nil, fmt.Errorf("can`t get phones: %w", err)
	}

	return phones, err
}
