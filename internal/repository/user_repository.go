package repository

import "github.com/Karaulkin/fio-rest-api/internal/models"

type UserRepository struct {
	db *DB
}

func NewUsersRepository(db *DB) *UserRepository {
	return &UserRepository{db}
}

// TODO: Для получения данных с различными фильтрами и пагинацией

// CreateUser для добавления новых людей в формате
func (u *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	panic("implement me")
}

// FindByUserId для удаления по индификатору
func (u *UserRepository) DeleteUserById(id int64) (models.User, error) {
	panic("implement me")
}

// EditeByUser для изменения сущности
func (u *UserRepository) UpdateUser(response models.UserResponse) (models.User, error) {
	panic("implement me")
}
