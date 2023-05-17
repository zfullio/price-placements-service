package policy

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
	"github.com/zfullio/price-placements-service/internal/domain/service"
	"github.com/zfullio/price-placements-service/internal/domain/usecase/phone"
	"github.com/zfullio/price-placements-service/internal/repository/avito"
	"github.com/zfullio/price-placements-service/internal/repository/cian"
	"github.com/zfullio/price-placements-service/internal/repository/domclick"
	"github.com/zfullio/price-placements-service/internal/repository/realty"
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

func (f FeedPolicy) CheckPhonesAll(ctx context.Context, spreadsheetID string, developer string) ([]entity.CheckResult, error) {
	f.logger.Trace().Msg("CheckPhonesAll")

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID, developer)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}

		mapPlacements := make(map[entity.Placement]bool, 0)
		for _, feed := range feedsAll {
			if _, ok := mapPlacements[feed.Placement]; !ok {
				mapPlacements[feed.Placement] = true
			}
		}

		resultsAll := make([]entity.CheckResult, 0)
		resultsCh := make(chan []entity.CheckResult, len(entity.PlacementsAll))
		eg, gCtx := errgroup.WithContext(ctx)

		for placement := range mapPlacements {
			placement := placement

			eg.Go(func() error {
				select {
				case <-gCtx.Done():
					return gCtx.Err()
				default:
					err := f.CheckPhonesPlacement(gCtx, feedsAll, spreadsheetID, placement, resultsCh)
					if err != nil {
						return fmt.Errorf("can't check %s: %w", string(placement), err)
					}
				}

				return nil
			})
		}

		if err := eg.Wait(); err != nil {
			return nil, fmt.Errorf("ошибка проверки площадок: %w", err)
		}

		for i := 0; i < len(mapPlacements); i++ {
			resultsAll = append(resultsAll, <-resultsCh...)
		}

		return resultsAll, err
	}
}

func (f FeedPolicy) CheckPhonesRealty(ctx context.Context, spreadsheetID string, developer string) ([]entity.CheckResult, error) {
	f.logger.Trace().Msg("CheckPhonesRealty")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID, developer)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}

		resultsCh := make(chan []entity.CheckResult, 1)

		err = f.CheckPhonesPlacement(ctx, feedsAll, spreadsheetID, entity.Realty, resultsCh)
		if err != nil {
			return nil, fmt.Errorf("ошибка проверки Realty: %w", err)
		}

		result := <-resultsCh

		return result, err
	}
}

func (f FeedPolicy) CheckPhonesCian(ctx context.Context, spreadsheetID string, developer string) ([]entity.CheckResult, error) {
	f.logger.Trace().Msg("CheckPhonesCian")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID, developer)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}

		errCh := make(chan error)
		resultsCh := make(chan []entity.CheckResult)

		go func() {
			err := f.CheckPhonesPlacement(ctx, feedsAll, spreadsheetID, entity.Cian, resultsCh)
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

func (f FeedPolicy) CheckPhonesAvito(ctx context.Context, spreadsheetID string, developer string) ([]entity.CheckResult, error) {
	f.logger.Trace().Msg("CheckPhonesAvito")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID, developer)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}

		errCh := make(chan error)
		resultsCh := make(chan []entity.CheckResult)

		go func() {
			err := f.CheckPhonesPlacement(ctx, feedsAll, spreadsheetID, entity.Avito, resultsCh)
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

func (f FeedPolicy) CheckPhonesDomclick(ctx context.Context, spreadsheetID string, developer string) ([]entity.CheckResult, error) {
	f.logger.Trace().Msg("CheckPhonesDomclick")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		feedsAll, err := f.feedService.Get(spreadsheetID, developer)
		if err != nil {
			return nil, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
		}

		errCh := make(chan error)
		resultsCh := make(chan []entity.CheckResult)

		go func() {
			err := f.CheckPhonesPlacement(ctx, feedsAll, spreadsheetID, entity.DomClick, resultsCh)
			if err != nil {
				errCh <- err
			}
			errCh <- nil
		}()

		err = <-errCh
		if err != nil {
			return nil, fmt.Errorf("ошибка проверки DomClick: %w", err)
		}

		return <-resultsCh, err
	}
}

func (f FeedPolicy) CheckPhonesPlacement(ctx context.Context, feedsAll []entity.Feed, spreadsheetID string, placement entity.Placement, resultsCh chan []entity.CheckResult) error {
	f.logger.Trace().Str("type", string(placement)).Msg("CheckPhonesPlacement")
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		filteredFeeds := filterByPlacement(feedsAll, placement)

		lotSrv, err := f.initLotServiceByPlacement(placement)
		if err != nil {
			return err
		}

		phUseCase := phone.NewPhoneUseCase(f.phoneService, lotSrv, f.logger)

		results := make([]entity.CheckResult, 0)

		for _, feed := range filteredFeeds {
			result, err := phUseCase.CheckNumbers(ctx, spreadsheetID, feed.Developer, feed.URL, feed.Placement)
			if err != nil {
				results = append(results, entity.CheckResult{
					Developer: feed.Developer,
					Placement: feed.Placement,
					Base:      feed.BaseFeed,
					URL:       feed.URL,
					Message:   err.Error(),
					Status:    entity.Error,
				})

				continue
			}

			if len(result) > 0 {
				results = append(results, result...)
			} else {
				results = append(results, entity.CheckResult{
					Developer: feed.Developer,
					Placement: feed.Placement,
					Base:      feed.BaseFeed,
					URL:       feed.URL,
					Message:   "warnings not found",
					Status:    entity.OK,
				})
			}
		}
		resultsCh <- results

		return nil
	}
}

func (f FeedPolicy) ValidateFeed(ctx context.Context, feedUrl string, placement entity.Placement) (results []entity.CheckResult, err error) {
	f.logger.Trace().Msg("ValidateFeed")

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		lotSrv, err := f.initLotServiceByPlacement(placement)
		if err != nil {
			return nil, err
		}

		rawResult, err := lotSrv.Validate(feedUrl)
		if err != nil {
			return results, err
		}

		results = prepareResult(rawResult, &entity.Feed{
			Developer: "",
			Placement: placement,
			AreaType:  "",
			Status:    entity.Unknown,
			URL:       feedUrl,
			BaseFeed:  entity.UnknownPlacement,
		})

		return results, nil
	}
}

func (f FeedPolicy) ValidateFeedAll(ctx context.Context, spreadsheetID string, developer string) (results []entity.CheckResult, err error) {
	f.logger.Trace().Msg("ValidateFeedAll")

	feedsAll, err := f.feedService.Get(spreadsheetID, developer)
	if err != nil {
		return results, fmt.Errorf("spreadsheet_id: %s|ошибка получения настроек фидов: %w", spreadsheetID, err)
	}

	resultsAll := make([]entity.CheckResult, 0)
	resultsCh := make(chan []entity.CheckResult, len(feedsAll))

	var eg, gCtx = errgroup.WithContext(ctx)

	for _, feed := range feedsAll {
		feed := feed

		eg.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
				rawResult, err := f.ValidateFeed(gCtx, feed.URL, feed.Placement)
				result := make([]entity.CheckResult, 0, len(rawResult))
				for _, item := range rawResult {
					item.Developer = feed.Developer
					item.Base = feed.BaseFeed
					result = append(result, entity.CheckResult{
						Developer: feed.Developer,
						Placement: feed.Placement,
						Base:      feed.BaseFeed,
						URL:       feed.URL,
						Message:   item.Message,
						Status:    item.Status,
					})
				}
				if err != nil {
					return err
				}
				resultsCh <- result

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

func (f FeedPolicy) initLotServiceByPlacement(placement entity.Placement) (service.LotService, error) {
	var lotRepo service.LotRepository

	var lotSrv *service.LotService

	switch placement {
	case entity.Realty:
		lotRepo = realty.NewLotRepository(f.logger)
		lotSrv = service.NewLotService(lotRepo, f.logger)
	case entity.Cian:
		lotRepo = cian.NewLotRepository(f.logger)
		lotSrv = service.NewLotService(lotRepo, f.logger)
	case entity.Avito:
		lotRepo = avito.NewLotRepository(f.logger)
		lotSrv = service.NewLotService(lotRepo, f.logger)
	case entity.DomClick:
		lotRepo = domclick.NewLotRepository(f.logger)
		lotSrv = service.NewLotService(lotRepo, f.logger)
	case entity.M2:
		lotRepo = realty.NewLotRepository(f.logger)
		lotSrv = service.NewLotService(lotRepo, f.logger)
	case entity.GdeTotDom:
		lotRepo = realty.NewLotRepository(f.logger)
		lotSrv = service.NewLotService(lotRepo, f.logger)
	case entity.NovostroyM:
		lotRepo = realty.NewLotRepository(f.logger)
		lotSrv = service.NewLotService(lotRepo, f.logger)
	case entity.CalugaHouse:
		lotRepo = realty.NewLotRepository(f.logger)
		lotSrv = service.NewLotService(lotRepo, f.logger)
	case entity.JCat:
		lotRepo = realty.NewLotRepository(f.logger)
		lotSrv = service.NewLotService(lotRepo, f.logger)
	default:
		return service.LotService{}, fmt.Errorf("invalid placement: %s", placement)
	}

	return *lotSrv, nil
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

func prepareResult(rawResult []string, feed *entity.Feed) (result []entity.CheckResult) {
	if len(rawResult) > 0 {
		for _, item := range rawResult {
			result = append(result, entity.CheckResult{
				Developer: feed.Developer,
				Placement: feed.Placement,
				Base:      feed.BaseFeed,
				URL:       feed.URL,
				Message:   item,
				Status:    entity.Warning,
			})
		}
	} else {
		result = append(result, entity.CheckResult{
			Developer: feed.Developer,
			Placement: feed.Placement,
			Base:      feed.BaseFeed,
			URL:       feed.URL,
			Message:   "warnings not found",
			Status:    entity.OK,
		})
	}

	return result
}
