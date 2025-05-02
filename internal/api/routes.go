package api

import (
	"github.com/Karaulkin/fio-rest-api/internal/api/handlers"
	"github.com/Karaulkin/fio-rest-api/internal/config"
	"github.com/labstack/echo/v4"
)

// SetupRoutes настраивает маршруты API
func SetupRoutes(e *echo.Echo, cfg *config.Config, userHandler *handlers.UserHandler) {
	// TODO:
}
