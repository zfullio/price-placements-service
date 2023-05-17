package service

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
)

type FeedRepository interface {
	Get(spreadsheetID string, developer string) (feeds []entity.Feed, err error)
}

type FeedService struct {
	repo   FeedRepository
	logger *zerolog.Logger
}

func NewFeedService(repo FeedRepository, logger *zerolog.Logger) *FeedService {
	serviceLogger := logger.With().Str("service", "feed").Logger()

	return &FeedService{
		repo:   repo,
		logger: &serviceLogger,
	}
}

func (fs FeedService) Get(spreadsheetId string, developer string) ([]entity.Feed, error) {
	fs.logger.Trace().Msg("Get")

	feeds, err := fs.repo.Get(spreadsheetId, developer)
	if err != nil {
		return nil, fmt.Errorf("can`t get feeds: %w", err)
	}

	return feeds, nil
}

func (fs FeedService) Validate(spreadsheetId string, developer string) ([]entity.Feed, error) {
	fs.logger.Trace().Msg("Validate")

	feeds, err := fs.repo.Get(spreadsheetId, developer)
	if err != nil {
		return nil, fmt.Errorf("can`t get feeds: %w", err)
	}

	return feeds, nil
}
