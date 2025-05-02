package utils

import (
	"github.com/Karaulkin/fio-rest-api/internal/config"
	"log/slog"
	"os"
)

var (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Log.Level {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
