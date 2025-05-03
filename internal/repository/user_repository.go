package repository

import (
	"github.com/Karaulkin/fio-rest-api/internal/models"
	pg "github.com/Karaulkin/fio-rest-api/internal/repository/postgres"
)

type UserRepository struct {
	db *pg.DB
}

// TODO:установить уровень изоляции commited read

func NewUsersRepository(db *pg.DB) *UserRepository {
	return &UserRepository{db}
}

// TODO: Для получения данных с различными фильтрами и пагинацией
func (u *UserRepository) GetUsers(name string, page, pageSize, total int) ([]*models.User, error) {
	var users []*models.User
	/*
		select * from users
		where name = "name"

	*/
}

// CreateUser для добавления новых людей в формате (Корректное сообщение обогатить)
func (u *UserRepository) CreateUser(name, surname, patronymic string) error {
	var createdUser *models.User = &models.User{
		Name:       name,
		Surname:    surname,
		Patronymic: patronymic,
	}

	/*
		INSERT INTO users (name, surname, patronymic, age, gender, nationality)
		VALUES (?, ?, ?, ?, ?, ?)
	*/

	// TODO: err if field is empty
}

// FindByUserId для удаления по индификатору
func (u *UserRepository) DeleteUserById(id int64) error {
	var user models.User
	/*
		delete from users
			where id = ?
	*/
	panic("implement me")
}

// EditeByUser для изменения сущности
func (u *UserRepository) UpdateUser(user models.User) error {
	/*
			UPDATE users
				SET name = ?,
		    		surname = ?,
					patronymic = ?,
					age = ?,
					gender = ?,
					nationality = ?
		[WHERE id = ?]

	*/

	panic("implement me")
}
