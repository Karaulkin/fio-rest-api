package repository

import (
	"fmt"
	"github.com/Karaulkin/fio-rest-api/internal/models"
	pg "github.com/Karaulkin/fio-rest-api/internal/repository/postgres"
)

type UserRepository struct {
	db *pg.DB
}

func NewUsersRepository(db *pg.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) GetUsers(name string, page, pageSize int) ([]models.User, error) {
	offset := (page - 1) * pageSize
	var users []models.User

	query := `
        SELECT id, name, surname, patronymic, age, gender, nationality
        FROM users
        WHERE ($1 = '' OR name ILIKE '%'  $1  '%')
        ORDER BY id
        LIMIT $2 OFFSET $3;
    `

	rows, err := u.db.Query(query, name, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID, &user.Name, &user.Surname,
			&user.Patronymic, &user.Age, &user.Gender, &user.Nationality,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) GetUser(id int64) (models.User, error) {
	var user models.User

	query := `
		SELECT * FROM users WHERE id = $1;
	`

	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality)

	if err != nil {
		return models.User{}, fmt.Errorf("failed to query user: %w", err)
	}

	return user, nil
}

// CreateUser для добавления новых людей в формате (Корректное сообщение обогатить)
func (u *UserRepository) CreateUser(user *models.User) error {
	if user.Name == "" || user.Surname == "" {
		return fmt.Errorf("name and surname are required")
	}

	query := `
        INSERT INTO users (name, surname, patronymic, age, gender, nationality)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id;
    `

	err := u.db.QueryRow(query, user.Name, user.Surname, user.Patronymic, user.Age, user.Gender, user.Nationality).
		Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

// FindByUserId для удаления по индификатору
func (u *UserRepository) DeleteUserById(id int64) error {
	query := `
		DELETE FROM users WHERE id = $1;
	`
	result, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}

// UpdateUserобновляет пользоваткля
func (u *UserRepository) UpdateUser(user *models.User) error {
	query := `
        UPDATE users
        SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, nationality = $6
        WHERE id = $7;
    `

	result, err := u.db.Exec(query,
		user.Name, user.Surname, user.Patronymic,
		user.Age, user.Gender, user.Nationality,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", user.ID)
	}

	return nil
}
