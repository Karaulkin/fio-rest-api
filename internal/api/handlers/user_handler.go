package handlers

import "github.com/Karaulkin/fio-rest-api/internal/service"

type UserHandler struct {
	serviceUser *service.ServiceUser
}

// NewUserHandler создает новый обработчик для работы с пользователями
func NewUserHandler(serviceUser *service.ServiceUser) *UserHandler {
	return &UserHandler{
		serviceUser: serviceUser,
	}
}

// TODO: handlers
