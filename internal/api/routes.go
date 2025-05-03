package api

import (
	"github.com/Karaulkin/fio-rest-api/internal/api/handlers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// SetupRoutes настраивает маршруты API
func SetupRoutes(e *echo.Echo, userHandler *handlers.UserHandler) {
	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Префикс для API
	api := e.Group("/api/v1")

	// Хэндлеры пользователей
	api.GET("/users", userHandler.GetUsers)          // Получить всех пользователей с фильтрами и пагинацией
	api.POST("/users", userHandler.CreateUser)       // Добавить нового пользователя
	api.DELETE("/users/:id", userHandler.DeleteUser) // Удалить пользователя по ID
	api.PUT("/users/:id", userHandler.UpdateUser)    // Обновить пользователя по ID

	// Health-check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
}
