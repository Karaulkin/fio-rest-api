package service

import (
	"github.com/Karaulkin/fio-rest-api/internal/models"
	"github.com/Karaulkin/fio-rest-api/internal/repository"
)

type User interface {
	// TODO: Для получения данных с различными фильтрами и пагинацией
	CreateUser(user *models.User) (*models.User, error)
	DeleteUserById(id int64) (models.User, error)
	UpdateUser(response models.UserResponse) (models.User, error)
}

type UserService struct {
	userRepo *repository.UserRepository
}

func NewServiceUser(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// TODO:сделать сервисы
// TODO:сделать с паганациями и фильтрами

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	panic("implement me")
}

func (s *UserService) DeleteUserById(userId int64) (models.User, error) {
	panic("implement me")
}

func (s *UserService) UpdateProfileUser(response models.UserResponse) (models.User, error) {
	panic("implement me")
}
