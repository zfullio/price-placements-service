package policy

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/adapters/avito"
	"github.com/zfullio/price-placements-service/internal/adapters/cian"
	"github.com/zfullio/price-placements-service/internal/adapters/domclick"
	"github.com/zfullio/price-placements-service/internal/adapters/realty"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
	"github.com/zfullio/price-placements-service/internal/domain/service"
	"github.com/zfullio/price-placements-service/internal/domain/usecase/phone"
	"golang.org/x/sync/errgroup"
)

type FeedPolicy struct {
	feedService  service.FeedService
	lotService   service.LotService
	phoneService service.PhoneService
	logger       *zerolog.Logger
}

func NewFeedPolicy(feedService service.FeedService, phoneService service.PhoneService, logger *zerolog.Logger) *FeedPolicy {
	policyLogger := logger.With().Str("policy", "feed").Logger()
	return &FeedPolicy{
		feedService:  feedService,
		lotService:   service.LotService{},
		phoneService: phoneService,
		logger:       &policyLogger,
	}
}

func (f FeedPolicy) CheckPhonesAll(ctx context.Context, spreadsheetID string) ([]string, error) {
	f.logger.Trace().Msg("CheckPhonesAll")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}

		resultsAll := make([]string, 0)
		var placements = []entity.Placement{entity.Realty, entity.Cian, entity.Domclick, entity.Avito}
		resultsCh := make(chan []string, len(placements))
		eg, gCtx := errgroup.WithContext(ctx)

		for _, placement := range placements {
			placement := placement
			eg.Go(func() error {
				select {
				case <-gCtx.Done():
					return gCtx.Err()
				default:
					err := f.checkPhonesPlacement(gCtx, feedsAll, spreadsheetID, placement, resultsCh)
					if err != nil {
						return fmt.Errorf("can't check Realty: %w", err)
					}
				}
				return nil
			})
		}

		if err := eg.Wait(); err != nil {
			return nil, fmt.Errorf("ошибка проверки площадок: %w", err)
		}

		for i := 0; i < len(placements); i++ {
			resultsAll = append(resultsAll, <-resultsCh...)
		}

		return resultsAll, err
	}
}

func (f FeedPolicy) CheckPhonesRealty(ctx context.Context, spreadsheetID string) ([]string, error) {
	f.logger.Trace().Msg("CheckPhonesRealty")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}

		resultsCh := make(chan []string, 1)

		err = f.checkPhonesPlacement(ctx, feedsAll, spreadsheetID, entity.Realty, resultsCh)
		if err != nil {
			return nil, fmt.Errorf("ошибка проверки Realty: %w", err)
		}

		result := <-resultsCh
		return result, err
	}
}

func (f FeedPolicy) CheckPhonesCian(ctx context.Context, spreadsheetID string) ([]string, error) {
	f.logger.Trace().Msg("CheckPhonesCian")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}
		errCh := make(chan error)
		resultsCh := make(chan []string)
		go func() {
			err := f.checkPhonesPlacement(ctx, feedsAll, spreadsheetID, entity.Cian, resultsCh)
			if err != nil {
				errCh <- err
			}
			errCh <- nil
		}()
		err = <-errCh
		if err != nil {
			return nil, fmt.Errorf("ошибка проверки Cian: %w", err)
		}
		return <-resultsCh, err
	}
}

func (f FeedPolicy) CheckPhonesAvito(ctx context.Context, spreadsheetID string) ([]string, error) {
	f.logger.Trace().Msg("CheckPhonesAvito")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}
		errCh := make(chan error)
		resultsCh := make(chan []string)

		go func() {
			err := f.checkPhonesPlacement(ctx, feedsAll, spreadsheetID, entity.Avito, resultsCh)
			if err != nil {
				errCh <- err
			}
			errCh <- nil
		}()
		err = <-errCh
		if err != nil {
			return nil, fmt.Errorf("ошибка проверки Avito: %w", err)
		}
		return <-resultsCh, err
	}
}

func (f FeedPolicy) CheckPhonesDomclick(ctx context.Context, spreadsheetID string) ([]string, error) {
	f.logger.Trace().Msg("CheckPhonesDomclick")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}
		errCh := make(chan error)
		resultsCh := make(chan []string)

		go func() {
			err := f.checkPhonesPlacement(ctx, feedsAll, spreadsheetID, entity.Domclick, resultsCh)
			if err != nil {
				errCh <- err
			}
			errCh <- nil
		}()
		err = <-errCh
		if err != nil {
			return nil, fmt.Errorf("ошибка проверки Domclick: %w", err)
		}
		return <-resultsCh, err
	}
}

func (f FeedPolicy) checkPhonesPlacement(ctx context.Context, feedsAll []entity.Feed, spreadsheetID string, placement entity.Placement, resultsCh chan []string) error {
	f.logger.Trace().Str("type", string(placement)).Msg("checkPhonesPlacement")
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		filteredFeeds := filterByPlacement(feedsAll, placement)
		phUseCase := phone.NewPhoneUseCase(nil, nil, f.logger)
		switch placement {
		case entity.Realty:
			lotRepo := realty.NewLotRepository(f.logger)
			lotSrv := service.NewLotService(lotRepo, f.logger)
			phUseCase = phone.NewPhoneUseCase(f.phoneService, lotSrv, f.logger)
		case entity.Cian:
			lotRepo := cian.NewLotRepository(f.logger)
			lotSrv := service.NewLotService(lotRepo, f.logger)
			phUseCase = phone.NewPhoneUseCase(f.phoneService, lotSrv, f.logger)
		case entity.Avito:
			lotRepo := avito.NewLotRepository(f.logger)
			lotSrv := service.NewLotService(lotRepo, f.logger)
			phUseCase = phone.NewPhoneUseCase(f.phoneService, lotSrv, f.logger)
		case entity.Domclick:
			lotRepo := domclick.NewLotRepository(f.logger)
			lotSrv := service.NewLotService(lotRepo, f.logger)
			phUseCase = phone.NewPhoneUseCase(f.phoneService, lotSrv, f.logger)
		}
		results := make([]string, 0)
		for _, feed := range filteredFeeds {
			result, err := phUseCase.CheckNumbers(ctx, spreadsheetID, feed.Url, feed.Placement)
			if err != nil {
				results = append(results, fmt.Sprintf("%s|%s|%s|Ошибка сопоставления: %s", feed.Developer, feed.Placement, feed.Url, err))
				continue
			}
			if len(result) > 0 {
				for _, item := range result {
					results = append(results, fmt.Sprintf("%s|%s|%s|%s", feed.Developer, feed.Placement, feed.Url, item))
				}
			} else {
				results = append(results, fmt.Sprintf("%s|%s|%s|Все в порядке", feed.Developer, feed.Placement, feed.Url))
			}
		}
		resultsCh <- results
		return nil
	}
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

func (f FeedPolicy) ValidateFeed(ctx context.Context, feedUrl string, placement entity.Placement) (results []string, err error) {
	f.logger.Trace().Msg("ValidateFeed")

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		lotSrv := service.NewLotService(nil, f.logger)

		switch placement {
		case entity.Realty:
			lotRepo := realty.NewLotRepository(f.logger)
			lotSrv = service.NewLotService(lotRepo, f.logger)
		case entity.Cian:
			lotRepo := cian.NewLotRepository(f.logger)
			lotSrv = service.NewLotService(lotRepo, f.logger)
		case entity.Avito:
			lotRepo := avito.NewLotRepository(f.logger)
			lotSrv = service.NewLotService(lotRepo, f.logger)
		case entity.Domclick:
			lotRepo := domclick.NewLotRepository(f.logger)
			lotSrv = service.NewLotService(lotRepo, f.logger)
		default:
			return results, fmt.Errorf("invalid placement: %s", placement)
		}

		results, err = lotSrv.Validate(feedUrl)
		if err != nil {
			return results, err
		}

		return results, nil
	}
}

func (f FeedPolicy) ValidateFeedAll(ctx context.Context, spreadsheetID string) (results []string, err error) {
	f.logger.Trace().Msg("ValidateFeedAll")

	feedsAll, err := f.feedService.Get(spreadsheetID)
	if err != nil {
		return results, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
	}

	resultsAll := make([]string, 0)
	resultsCh := make(chan []string, len(feedsAll))
	var eg, gCtx = errgroup.WithContext(ctx)
	for _, feed := range feedsAll {
		feed := feed
		eg.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
				result, err := f.ValidateFeed(gCtx, feed.Url, feed.Placement)
				if err != nil {
					return err
				}
				resultsCh <- prepareResult(result, &feed)

				return nil
			}
		})
	}
	eg.Go(func() error {
		for i := 0; i < len(feedsAll); i++ {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
				resultsAll = append(resultsAll, <-resultsCh...)
			}
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return resultsAll, err
	}

	return resultsAll, err
}

func prepareResult(rawResult []string, feed *entity.Feed) (result []string) {
	if len(rawResult) > 0 {
		for _, item := range result {
			result = append(result, fmt.Sprintf("%s|%s|%s|%s", feed.Developer, feed.Placement, feed.Url, item))
		}
	} else {
		result = append(result, fmt.Sprintf("%s|%s|%s|Все в порядке", feed.Developer, feed.Placement, feed.Url))
	}
	return result
}
