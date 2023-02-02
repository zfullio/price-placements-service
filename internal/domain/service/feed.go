package service

import (
	"github.com/zfullio/price-placements-service/internal/domain/entity"
)

type FeedRepository interface {
	Get(spreadsheetID string) (feeds []entity.Feed, err error)
}

type FeedService struct {
	repo FeedRepository
}

func NewFeedService(repo FeedRepository) *FeedService {
	return &FeedService{repo: repo}
}

func (fs FeedService) Get(spreadsheetId string) (feeds []entity.Feed, err error) {
	return fs.repo.Get(spreadsheetId)
}

func (fs FeedService) Validate(spreadsheetId string) (feeds []entity.Feed, err error) {
	return fs.repo.Get(spreadsheetId)
}
