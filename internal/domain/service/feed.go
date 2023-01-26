package service

import "checkphones/internal/domain/entity"

type FeedRepository interface {
	Get(spreadsheetID string) (feeds []entity.Feed, err error)
}

type feedService struct {
	repo FeedRepository
}

func NewFeedService(repo FeedRepository) *feedService {
	return &feedService{repo: repo}
}

func (fs feedService) Get(spreadsheetId string) (feeds []entity.Feed, err error) {
	return fs.repo.Get(spreadsheetId)
}
