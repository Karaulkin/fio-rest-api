package service

import (
	"github.com/Karaulkin/fio-rest-api/internal/models"
	"github.com/Karaulkin/fio-rest-api/internal/repository"
)

type User interface {
	// TODO: Для получения данных с различными фильтрами и пагинацией
	CreateUser(user *models.User) (*models.User, error)
	FindByUserId(id int64) (models.User, error)
	EditeByUser(response models.UserResponse) (models.User, error)
}

type ServiceUser struct {
	userRepo *repository.UserRepository
}

func NewServiceUser(userRepo *repository.UserRepository) *ServiceUser {
	return &ServiceUser{userRepo: userRepo}
}

// TODO:сделать сервис

func (s *ServiceUser) CreateUser(user *models.User) (*models.User, error) {
	panic("implement me")
}

func (s *ServiceUser) FindByUserId(userId int64) (models.User, error) {
	panic("implement me")
}

func (s *ServiceUser) EditeByUser(response models.UserResponse) (models.User, error) {
	panic("implement me")
}
