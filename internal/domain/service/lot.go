package service

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
)

type LotRepository interface {
	Get(url string) (lots []entity.Lot, err error)
	Validate(url string) (results []string, err error)
}

type LotService struct {
	repo   LotRepository
	logger *zerolog.Logger
}

func NewLotService(repo LotRepository, logger *zerolog.Logger) *LotService {
	serviceLogger := logger.With().Str("service", "lot").Logger()
	return &LotService{
		repo:   repo,
		logger: &serviceLogger,
	}
}

func (ls LotService) Get(url string) (lots []entity.Lot, err error) {
	ls.logger.Trace().Msg("Get")

	lots, err = ls.repo.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can`t get lots: %w", err)
	}

	return lots, err
}

func (ls LotService) Validate(url string) (results []string, err error) {
	ls.logger.Trace().Msg("Validate")

	results, err = ls.repo.Validate(url)
	if err != nil {
		return nil, fmt.Errorf("can`t validate lots: %w", err)
	}

	return results, err
}
