package main

import (
	"context"
	"log"
	"price-placements-service/internal/app"
	"price-placements-service/internal/config"
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
