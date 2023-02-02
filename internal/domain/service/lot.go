package service

import (
	"price-placements-service/internal/domain/entity"
)

type LotRepository interface {
	Get(url string) (lots []entity.Lot, err error)
	Validate(url string) (results []string, err error)
}

type LotService struct {
	repo LotRepository
}

func NewLotService(repo LotRepository) *LotService {
	return &LotService{repo: repo}
}

func (ls LotService) Get(url string) (lots []entity.Lot, err error) {
	return ls.repo.Get(url)
}

func (ls LotService) Validate(url string) (results []string, err error) {
	return ls.repo.Validate(url)
}
