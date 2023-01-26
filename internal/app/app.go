package app

import (
	"checkphones/internal/adapters/avito"
	"checkphones/internal/adapters/cian"
	"checkphones/internal/adapters/domclick"
	"checkphones/internal/adapters/gs"
	"checkphones/internal/adapters/realty"
	"checkphones/internal/config"
	"checkphones/internal/domain/entity"
	"checkphones/internal/domain/service"
	"checkphones/internal/domain/usecase/phone"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"strings"
)

func Run(ctx context.Context, cfg *config.Config) {
	ghSrv, err := sheets.NewService(ctx, option.WithCredentialsFile(cfg.GS.ServiceKeyPath))
	if err != nil {
		fmt.Println("не могу инициализировать google sheets")
	}
	ghPhoneRepo := gs.NewPhoneRepository(*ghSrv)
	ghFeedRepo := gs.NewFeedRepository(*ghSrv)

	phSrv := service.NewPhoneService(ghPhoneRepo)
	feedSrv := service.NewFeedService(ghFeedRepo)

	feeds, err := feedSrv.Get(cfg.SpreadsheetID)
	if err != nil {
		fmt.Printf("не могу получить настройки фидов: %s", err)
	}
	for _, feed := range feeds {
		result := make([]string, 0)
		fmt.Printf("Проверка %s. Url: %s\n", feed.Placement, feed.Url)
		switch feed.Placement {
		case entity.Realty:
			lotRepo := realty.NewLotRepository()
			lotSrv := service.NewLotService(lotRepo)
			phUsecase := phone.NewPhoneUseCase(phSrv, lotSrv)
			result, err = phUsecase.CheckNumbers(cfg.SpreadsheetID, feed.Url, feed.Placement)
			if err != nil {
				fmt.Printf("ошибка сопоставления: %s", err)
			}
		case entity.Cian:
			lotRepo := cian.NewLotRepository()
			lotSrv := service.NewLotService(lotRepo)
			phUsecase := phone.NewPhoneUseCase(phSrv, lotSrv)
			result, err = phUsecase.CheckNumbers(cfg.SpreadsheetID, feed.Url, feed.Placement)
			if err != nil {
				fmt.Printf("ошибка сопоставления: %s", err)
			}
		// TODO исправить сущность, когда разберусь
		case entity.NotFound:
			lotRepo := avito.NewLotRepository()
			lotSrv := service.NewLotService(lotRepo)
			phUsecase := phone.NewPhoneUseCase(phSrv, lotSrv)
			result, err = phUsecase.CheckNumbers(cfg.SpreadsheetID, feed.Url, feed.Placement)
			if err != nil {
				fmt.Printf("ошибка сопоставления: %s", err)
			}
		case entity.Domclick:
			lotRepo := domclick.NewLotRepository()
			lotSrv := service.NewLotService(lotRepo)
			phUsecase := phone.NewPhoneUseCase(phSrv, lotSrv)
			result, err = phUsecase.CheckNumbers(cfg.SpreadsheetID, feed.Url, feed.Placement)
			if err != nil {
				fmt.Printf("ошибка сопоставления: %s", err)
			}
		default:
			fmt.Printf("Я не умею обрабатывать данный тип фида: {%s}, проверьте настройки", feed.Placement)
		}
		if len(result) > 0 {
			fmt.Println(strings.Join(result, "\n"))
		}
	}

}
