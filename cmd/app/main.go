package main

import (
	"checkphones/internal/app"
	"checkphones/internal/config"
	"context"
	"log"
)

func main() {
	cfg, err := config.NewConfig("config.yml")
	if err != nil {
		log.Fatalln("ошибка инициализации настроек")
	}
	ctx := context.Background()
	app.Run(ctx, cfg)
}
