package service

import (
	"errors"
	"fmt"
	"github.com/Karaulkin/fio-rest-api/internal/models"
	"github.com/Karaulkin/fio-rest-api/internal/repository"
	"github.com/Karaulkin/fio-rest-api/internal/utils/client"
	"log/slog"
	"strings"
)

var (
	ErrNotFound          = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidInput      = errors.New("invalid input")
)

type User interface {
	GetUsers(name string, page, pageSize int) ([]models.User, error)
	CreateUser(user *models.User) error
	DeleteUserById(id int64) error
	GetUser(id int64) (models.User, error)
	UpdateUser(user *models.User) error
}

type UserService struct {
	userRepo User
	log      *slog.Logger
}

func NewServiceUser(userRepo *repository.UserRepository, log *slog.Logger) *UserService {
	return &UserService{
		userRepo: userRepo,
		log:      log,
	}
}

func (s *UserService) Create(user models.User) (models.User, error) {
	const op = "service.Create"

	s.log.Info(op, "Creating user")

	if checkUserField(user.Name) != nil || checkUserField(user.Surname) != nil {
		return models.User{}, ErrInvalidInput
	}

	enrichment, err := client.Enrich(user.Name)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to enrich data: %w", err)
	}

	user.Age = enrichment.Age
	user.Gender = enrichment.Gender
	user.Nationality = enrichment.Nationality

	err = s.userRepo.CreateUser(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *UserService) DeleteById(userId int64) (models.User, error) {
	const op = "service.DeleteById"

	s.log.Info(op, "Deleting user")

	if userId <= 0 {
		return models.User{}, ErrInvalidInput
	}

	user, err := s.userRepo.GetUser(userId)

	err = s.userRepo.DeleteUserById(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return models.User{}, ErrNotFound
		}
		return models.User{}, fmt.Errorf("failed to delete user: %w", err)
	}

	return user, nil
}

func (s *UserService) UpdateProfile(user models.User) (models.User, error) {
	const op = "service.UpdateProfile"

	s.log.Info(op, "Updating user")

	if checkUserField(user.Name) != nil && checkUserField(user.Surname) != nil && checkUserField(user.Patronymic) != nil {
		return models.User{}, ErrInvalidInput
	}

	oldDataUser, err := s.userRepo.GetUser(user.ID)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	user = createUpdateUser(oldDataUser, user)

	err = s.userRepo.UpdateUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return models.User{}, ErrNotFound
		}
		return models.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUsers(name string, page, pageSize int) ([]models.User, error) {
	const op = "service.GetUsers"

	s.log.Info(op, "Getting users")

	if checkUserField(name) != nil {
		return []models.User{}, ErrInvalidInput
	}

	if page < 1 || pageSize < 1 {
		return nil, ErrInvalidInput
	}

	users, err := s.userRepo.GetUsers(name, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func createUpdateUser(newUser, oldUser models.User) models.User {
	var updatedUser = oldUser

	if len(newUser.Name) > 0 {
		updatedUser.Name = newUser.Name
	}

	if len(newUser.Surname) > 0 {
		updatedUser.Surname = newUser.Surname
	}

	if len(newUser.Patronymic) > 0 {
		updatedUser.Patronymic = newUser.Patronymic
	}

	if len(newUser.Gender) > 0 {
		updatedUser.Gender = newUser.Gender
	}

	if len(newUser.Nationality) > 0 {
		updatedUser.Nationality = newUser.Nationality
	}

	if newUser.Age > 0 {
		updatedUser.Age = newUser.Age
	}

	return updatedUser

}

func checkUserField(name string) error {
	if len(name) <= 0 {
		return errors.New("field is empty")
	}

	for _, char := range name {
		if char > '0' && char < '9' {
			return errors.New("field is invalid")
		}
	}

	return nil
}
