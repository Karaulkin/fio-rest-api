package service

import (
	"errors"
	"fmt"
	"github.com/Karaulkin/fio-rest-api/internal/client"
	"github.com/Karaulkin/fio-rest-api/internal/models"
	"github.com/Karaulkin/fio-rest-api/internal/repository"
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
	log      *slog.Logger //пционально для дебаг слоя о
}

func NewServiceUser(userRepo *repository.UserRepository, log *slog.Logger) *UserService {
	return &UserService{userRepo: userRepo,
		log: log.WithGroup("service"),
	}
}

func (s *UserService) Create(user *models.User) (models.UserResponse, error) {
	if user.Name == "" || user.Surname == "" {
		return models.UserResponse{}, ErrInvalidInput
	}

	enrichment, err := client.Enrich(user.Name)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to enrich data: %w", err)
	}

	user.Age = enrichment.Age
	user.Gender = enrichment.Gender
	user.Nationality = enrichment.Nationality

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to create user: %w", err)
	}

	return models.UserResponse{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}, nil
}

func (s *UserService) DeleteById(userId int64) (models.UserResponse, error) {
	user, err := s.userRepo.GetUser(userId)

	err = s.userRepo.DeleteUserById(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return models.UserResponse{}, ErrNotFound
		}
		return models.UserResponse{}, fmt.Errorf("failed to delete user: %w", err)
	}

	return models.UserResponse{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}, nil
}

func (s *UserService) UpdateProfile(u models.User) (models.UserResponse, error) {
	err := s.userRepo.UpdateUser(&u)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return models.UserResponse{}, ErrNotFound
		}
		return models.UserResponse{}, fmt.Errorf("failed to update user: %w", err)
	}

	return models.UserResponse{
		Name:        u.Name,
		Surname:     u.Surname,
		Patronymic:  u.Patronymic,
		Age:         u.Age,
		Gender:      u.Gender,
		Nationality: u.Nationality,
	}, nil
}

func (s *UserService) GetUsers(name string, page, pageSize int) ([]models.User, error) {
	if page < 1 || pageSize < 1 {
		return nil, ErrInvalidInput
	}

	users, err := s.userRepo.GetUsers(name, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}
