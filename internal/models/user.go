package models

// User пользователь
type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

// UserCreateRequest для добавления новых юзеров
type UserCreateRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

// UserResponse для получения пользователя
type UserResponse struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

// UsersResponse для получения пользователей пагинацией и фильтрами
type UsersResponse struct {
	Users    []User `json:"users"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
