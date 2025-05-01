package main

import (
	"github.com/Karaulkin/fio-rest-api/internal/config"
	"github.com/Karaulkin/fio-rest-api/internal/utils"
	"log/slog"
)

func main() {
	// TODO: Загружаем конфигурацию
	cfg := config.MustLoad()

	log := utils.SetupLogger(cfg.Env)
	log.Info("Starting server", slog.String("env", cfg.Env))
}
