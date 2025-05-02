package main

import (
	"context"
	"github.com/Karaulkin/fio-rest-api/internal/api/handlers"
	"github.com/Karaulkin/fio-rest-api/internal/config"
	"github.com/Karaulkin/fio-rest-api/internal/repository"
	"github.com/Karaulkin/fio-rest-api/internal/service"
	"github.com/Karaulkin/fio-rest-api/internal/utils"
	"github.com/labstack/echo/v4"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	logger := utils.SetupLogger(cfg)
	logger.Info("Starting server", slog.String("log", cfg.Log.Level))

	// Инициализация базы данных
	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Запуск миграций
	if err := repository.RunMigrations(db.DB, "./migrations", logger); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Инициализация репозиториев
	userRepo := repository.NewUsersRepository(db)

	// Инициализация сервисов
	userService := service.NewServiceUser(userRepo)

	// Инициализация обработчиков
	userHandler := handlers.NewUserHandler(userService)

	// Инициализация Echo
	e := echo.New()
	// e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//e.Use(middleware.CORS())

	// Инициализация валидатора
	//e.Validator = customMiddleware.NewValidator()

	// Настройка маршрутов
	//api.SetupRoutes(e, cfg, userHandler, accountHandler, txHandler, otpHandler)

	// Запуск сервера
	go func() {
		if err := e.Start(":" + cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Ожидание сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.Timeout*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown server: %v", err)
	}

}
