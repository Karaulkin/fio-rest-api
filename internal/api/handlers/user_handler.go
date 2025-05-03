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

type UserHandler struct {
	userService *service.UserService
	log         *slog.Logger
}

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
// @Failure 500 {object} echo.HTTPError
// @Router /users [get]
func (uh *UserHandler) GetUsers(c echo.Context) error {
	name := c.QueryParam("name")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
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

// CreateUser Создание нового пользователя
// @Summary Добавить пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserCreateRequest true "Данные пользователя"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users [post]
func (uh *UserHandler) CreateUser(c echo.Context) error {
	var req models.UserCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	user := &models.User{
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
	}

	created, err := uh.userService.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, created)
}

// DeleteUser Удалить пользователя по ID
// @Summary Удалить пользователя
// @Tags users
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users/{id} [delete]
func (uh *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	deletedUser, err := uh.userService.DeleteById(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, deletedUser)
}

// UpdateUser Обновить данные пользователя
// @Summary Обновить пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body models.User true "Обновлённые данные"
// @Success 200 {object} models.User
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
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

	return c.JSON(http.StatusOK, updated)
}
