package service

import (
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
)

type FeedRepository interface {
	Get(spreadsheetID string) (feeds []entity.Feed, err error)
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

func (fs FeedService) Get(spreadsheetId string) (feeds []entity.Feed, err error) {
	fs.logger.Trace().Msg("Get")

	return fs.repo.Get(spreadsheetId)
}

func (fs FeedService) Validate(spreadsheetId string) (feeds []entity.Feed, err error) {
	fs.logger.Trace().Msg("Validate")

	return fs.repo.Get(spreadsheetId)
}
