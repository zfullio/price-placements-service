package main

import (
	"context"
	"github.com/zfullio/price-placements-service/internal/app"
	"github.com/zfullio/price-placements-service/internal/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig("config.yml")
	if err != nil {
		log.Fatalf("ошибка инициализации настроек: %s", err)
	}
	ctx := context.Background()
	a, err := app.NewApp(ctx, *cfg)
	if err != nil {
		log.Fatal("ошибка инициализации приложения")
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatal("ошибка приложения")
	}
}
