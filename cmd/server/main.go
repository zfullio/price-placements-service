package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/app/server"
	"github.com/zfullio/price-placements-service/internal/config"
	"os"
	"runtime/debug"
	"time"
)

const appName = "Price placements service"

func main() {
	var fileConfig = flag.String("f", "config.yml", "configuration file")
	var useEnv = flag.Bool("env", false, "use environment variables")
	var trace = flag.Bool("trace", false, "switch trace logging")
	flag.Parse()

	buildInfo, _ := debug.ReadBuildInfo()

	var logger zerolog.Logger
	if *trace {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Str("go_version", buildInfo.GoVersion).
			Logger()
		logger.Info().Msg("Logging level = Trace")
	} else {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()
	}

	if !*useEnv {
		logger.Info().Msgf("configuration file: %s", *fileConfig)
	} else {
		logger.Info().Msg("configuration from ENV")
	}

	cfg, err := config.NewServerConfig(*fileConfig, *useEnv)
	if err != nil {
		logger.Fatal().Err(err).Msg("Ошибка в файле настроек")
	}

	telegramService, err := telegram.New(cfg.TG.Token)
	if err != nil {
		logger.Fatal().Err(err).Msg("Ошибка в сервисе: Telegram")
	}
	telegramService.AddReceivers(cfg.Chat)

	appNotify := notify.New()
	appNotify.UseServices(telegramService)

	if !cfg.IsEnabled {
		notify.Disable(appNotify)
	}

	ctx := context.Background()
	a := server.NewApp(ctx, &logger, *cfg, appNotify)

	err = a.Run(ctx)
	if err != nil {
		a.Logger.Err(err).Msgf("%s |Ошибка приложения", appName)
		_ = a.Notify.Send(ctx, appName, fmt.Sprintf("Ошибка приложения. Err: %s", err))
	}
}
