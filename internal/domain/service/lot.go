package service

import (
	"checkphones/internal/domain/entity"
)

type LotRepository interface {
	Get(url string) (lots []entity.Lot, err error)
}

type lotService struct {
	repo LotRepository
}

func NewLotService(repo LotRepository) *lotService {
	return &lotService{repo: repo}
}

func (ls lotService) Get(url string) (lots []entity.Lot, err error) {
	return ls.repo.Get(url)
}
