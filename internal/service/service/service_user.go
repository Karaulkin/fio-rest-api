package service

import (
	"github.com/Karaulkin/fio-rest-api/internal/models"
	"github.com/Karaulkin/fio-rest-api/internal/repository"
)

type User interface {
	GetUsers(name string, page, pageSize int) ([]*models.User, error)
	CreateUser(name, surname, patronymic string) error
	DeleteUserById(id int64) error
	UpdateUser(user models.User) error
}

type UserService struct {
	userRepo *repository.UserRepository
}

func NewServiceUser(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	panic("implement me")
}

func (s *UserService) DeleteUserById(userId int64) (models.User, error) {
	panic("implement me")
}

func (s *UserService) UpdateProfileUser(response models.UserResponse) (models.User, error) {
	panic("implement me")
}
