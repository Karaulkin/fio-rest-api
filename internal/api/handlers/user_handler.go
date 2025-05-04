package handlers

import (
	"errors"
	"github.com/Karaulkin/fio-rest-api/internal/models"
	"github.com/Karaulkin/fio-rest-api/internal/service"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
)

const (
	minPage  = 1
	sizePage = 10
)

type UserHandler struct {
	userService *service.UserService
	log         *slog.Logger
}

// TODO: debug slice

// NewUserHandler создает новый обработчик для работы с пользователями
func NewUserHandler(userService *service.UserService, log *slog.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		log:         log.WithGroup("handler"),
	}
}

// GetUsers Получить список пользователей
// @Summary Получить пользователей
// @Tags users
// @Param name query string false "Имя для фильтрации"
// @Param page query int false "Номер страницы"
// @Param page_size query int false "Размер страницы"
// @Success 200 {object} models.UsersResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /users [get]
func (uh *UserHandler) GetUsers(c echo.Context) error {
	name := c.QueryParam("name")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	if page <= 0 {
		page = minPage
	}
	if pageSize <= 0 {
		pageSize = sizePage
	}

	users, err := uh.userService.GetUsers(name, page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, models.UsersResponse{
		Users:    users,
		Page:     page,
		PageSize: pageSize,
	})
}

// CreateUser godoc
// @Summary Добавить пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserCreateRequest true "Данные пользователя"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /users [post]
func (uh *UserHandler) CreateUser(c echo.Context) error {
	var req models.UserCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	user := models.User{
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
	}

	created, err := uh.userService.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, models.UserResponse{
		Name:        created.Name,
		Surname:     created.Surname,
		Patronymic:  created.Patronymic,
		Age:         created.Age,
		Gender:      created.Gender,
		Nationality: created.Nationality,
	})
}

// DeleteUser godoc
// @Summary Удалить пользователя
// @Tags users
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 404 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /users/{id} [delete]
func (uh *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	deleted, err := uh.userService.DeleteById(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, models.UserResponse{
		Name:        deleted.Name,
		Surname:     deleted.Surname,
		Patronymic:  deleted.Patronymic,
		Age:         deleted.Age,
		Gender:      deleted.Gender,
		Nationality: deleted.Nationality,
	})
}

// UpdateUser godoc
// @Summary Обновить пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body models.UserResponse true "Обновлённые данные"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 404 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /users/{id} [put]
func (uh *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	var input models.User
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}

	input.ID = id
	updated, err := uh.userService.UpdateProfile(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, models.UserResponse{
		Name:        updated.Name,
		Surname:     updated.Surname,
		Patronymic:  updated.Patronymic,
		Age:         updated.Age,
		Gender:      updated.Gender,
		Nationality: updated.Nationality,
	})
}
