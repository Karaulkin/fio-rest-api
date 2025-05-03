package handlers

import (
	"github.com/Karaulkin/fio-rest-api/internal/service/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	serviceUser *service.UserService
}

// NewUserHandler создает новый обработчик для работы с пользователями
func NewUserHandler(serviceUser *service.UserService) *UserHandler {
	return &UserHandler{
		serviceUser: serviceUser,
	}
}

// TODO: handlers
func (uh *UserHandler) RegisterUser(c echo.Context) error {
	panic("implement me")
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	panic("implement me")
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	panic("implement me")
}
