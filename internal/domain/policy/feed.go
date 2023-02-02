package policy

import (
	"fmt"
	"price-placements-service/internal/adapters/avito"
	"price-placements-service/internal/adapters/cian"
	"price-placements-service/internal/adapters/domclick"
	"price-placements-service/internal/adapters/realty"
	"price-placements-service/internal/domain/entity"
	"price-placements-service/internal/domain/service"
	"price-placements-service/internal/domain/usecase/phone"
)

type FeedPolicy struct {
	feedService  service.FeedService
	lotService   service.LotService
	phoneService service.PhoneService
}

func NewFeedPolicy(feedService service.FeedService, phoneService service.PhoneService) *FeedPolicy {
	return &FeedPolicy{feedService: feedService, phoneService: phoneService}
}

func (f FeedPolicy) CheckPhonesAll(spreadsheetID string) ([]string, error) {
	feedsAll, err := f.feedService.Get(spreadsheetID)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения настроек фидов: %w", err)
	}
	resultsAll := make([]string, 0)
	resultsRealty := f.checkPhonesPlacement(feedsAll, spreadsheetID, entity.Realty)
	resultsCian := f.checkPhonesPlacement(feedsAll, spreadsheetID, entity.Cian)
	resultsAvito := f.checkPhonesPlacement(feedsAll, spreadsheetID, entity.Avito)
	resultsDomclick := f.checkPhonesPlacement(feedsAll, spreadsheetID, entity.Domclick)
	resultsAll = append(resultsAll, resultsRealty...)
	resultsAll = append(resultsAll, resultsCian...)
	resultsAll = append(resultsAll, resultsAvito...)
	resultsAll = append(resultsAll, resultsDomclick...)
	return resultsAll, err
}

func (f FeedPolicy) CheckPhonesRealty(spreadsheetID string) ([]string, error) {
	feedsAll, err := f.feedService.Get(spreadsheetID)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения настроек фидов: %w", err)
	}
	results := f.checkPhonesPlacement(feedsAll, spreadsheetID, entity.Realty)
	return results, err
}

func (f FeedPolicy) CheckPhonesCian(spreadsheetID string) ([]string, error) {
	feedsAll, err := f.feedService.Get(spreadsheetID)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения настроек фидов: %w", err)
	}
	results := f.checkPhonesPlacement(feedsAll, spreadsheetID, entity.Cian)
	return results, err
}

func (f FeedPolicy) CheckPhonesAvito(spreadsheetID string) ([]string, error) {
	feedsAll, err := f.feedService.Get(spreadsheetID)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения настроек фидов: %w", err)
	}
	results := f.checkPhonesPlacement(feedsAll, spreadsheetID, entity.Avito)
	return results, err
}

func (f FeedPolicy) CheckPhonesDomclick(spreadsheetID string) ([]string, error) {
	feedsAll, err := f.feedService.Get(spreadsheetID)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения настроек фидов: %w", err)
	}
	results := f.checkPhonesPlacement(feedsAll, spreadsheetID, entity.Domclick)
	return results, err
}

func (f FeedPolicy) checkPhonesPlacement(feedsAll []entity.Feed, spreadsheetID string, placement entity.Placement) (results []string) {
	filteredFeeds := filterByPlacement(feedsAll, placement)
	phUsecase := phone.NewPhoneUseCase(nil, nil)
	switch placement {
	case entity.Realty:
		lotRepo := realty.NewLotRepository()
		lotSrv := service.NewLotService(lotRepo)
		phUsecase = phone.NewPhoneUseCase(f.phoneService, lotSrv)
	case entity.Cian:
		lotRepo := cian.NewLotRepository()
		lotSrv := service.NewLotService(lotRepo)
		phUsecase = phone.NewPhoneUseCase(f.phoneService, lotSrv)
	case entity.Avito:
		lotRepo := avito.NewLotRepository()
		lotSrv := service.NewLotService(lotRepo)
		phUsecase = phone.NewPhoneUseCase(f.phoneService, lotSrv)
	case entity.Domclick:
		lotRepo := domclick.NewLotRepository()
		lotSrv := service.NewLotService(lotRepo)
		phUsecase = phone.NewPhoneUseCase(f.phoneService, lotSrv)
	}
	for _, feed := range filteredFeeds {
		result := make([]string, 0)
		result, err := phUsecase.CheckNumbers(spreadsheetID, feed.Url, feed.Placement)
		if err != nil {
			results = append(results, fmt.Sprintf("%s: %s ошибка сопоставления: %s", feed.Placement, feed.Url, err))
			continue
		}
		if len(result) > 0 {
			results = append(results, result...)
		} else {
			results = append(results, fmt.Sprintf("%s: %s Все в порядке", feed.Placement, feed.Url))
		}
	}
	return results
}

func filterByPlacement(feeds []entity.Feed, placement entity.Placement) []entity.Feed {
	var result []entity.Feed
	for _, feed := range feeds {
		if feed.Placement == placement {
			result = append(result, feed)
		}
	}
	return result
}

func (f FeedPolicy) ValidateFeed(feedUrl string, placement entity.Placement) (results []string, err error) {
	lotSrv := service.NewLotService(nil)

	switch placement {
	case entity.Realty:
		lotRepo := realty.NewLotRepository()
		lotSrv = service.NewLotService(lotRepo)
	case entity.Cian:
		lotRepo := cian.NewLotRepository()
		lotSrv = service.NewLotService(lotRepo)
	case entity.Avito:
		lotRepo := avito.NewLotRepository()
		lotSrv = service.NewLotService(lotRepo)
	case entity.Domclick:
		lotRepo := domclick.NewLotRepository()
		lotSrv = service.NewLotService(lotRepo)
	}

	results, err = lotSrv.Validate(feedUrl)
	return results, err
}

func (f FeedPolicy) ValidateFeedAll(spreadsheetID string) (results []string, err error) {
	feedsAll, err := f.feedService.Get(spreadsheetID)
	if err != nil {
		return results, fmt.Errorf("ошибка получения настроек фидов: %w", err)
	}

	resultsAll := make([]string, 0)
	for _, feed := range feedsAll {
		results, err := f.ValidateFeed(feed.Url, feed.Placement)
		if err != nil {
			return results, err
		}

		if len(results) > 0 {
			resultsAll = append(resultsAll, results...)
		} else {
			resultsAll = append(resultsAll, fmt.Sprintf("%s: %s Все в порядке", feed.Placement, feed.Url))
		}
	}
	return resultsAll, err
}
