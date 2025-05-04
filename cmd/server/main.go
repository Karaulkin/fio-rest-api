// @title FIO REST API
// @version 1.0
// @description API для обработки ФИО и обогащения данными

// @host localhost:8080
// @BasePath /api/v1

// @schemes http
package main

import (
	"context"
	_ "github.com/Karaulkin/fio-rest-api/docs"
	"github.com/Karaulkin/fio-rest-api/internal/api"
	"github.com/Karaulkin/fio-rest-api/internal/api/handlers"
	"github.com/Karaulkin/fio-rest-api/internal/config"
	customMiddleware "github.com/Karaulkin/fio-rest-api/internal/midleware"
	"github.com/Karaulkin/fio-rest-api/internal/repository"
	"github.com/Karaulkin/fio-rest-api/internal/repository/postgres"
	"github.com/Karaulkin/fio-rest-api/internal/service"
	"github.com/Karaulkin/fio-rest-api/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	log := utils.SetupLogger(cfg)
	log.Info("Starting server", slog.String("log", cfg.Log.Level))

	// Инициализация базы данных
	log.Debug("Init database")
	db, err := postgres.NewDB(cfg)
	if err != nil {
		log.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	// Запуск миграций
	if err := postgres.RunMigrations(db.DB, "./migrations", log); err != nil {
		log.Error("Failed to run migrations: %v", err)
		os.Exit(1)
	}

	// Инициализация репозиториев
	log.Debug("Up user repo service")
	userRepo := repository.NewUsersRepository(db, log)

	// Инициализация сервисов
	log.Debug("Up user service")
	userService := service.NewServiceUser(userRepo, log)

	// Инициализация обработчиков
	log.Debug("Up user handler")
	userHandler := handlers.NewUserHandler(userService, log)

	// Инициализация Echo
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Инициализация валидатора
	e.Validator = customMiddleware.NewValidator()

	// Настройка маршрутов
	log.Debug("Init routes")
	api.SetupRoutes(e, userHandler)

	// Запуск сервера
	go func() {
		if err := e.Start(":" + cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			log.Error("Failed to start server: %v", err)
			os.Exit(1)
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
		log.Error("Failed to shutdown server: %v", err)
		os.Exit(1)
	}

}
