package repository

import "errors"

var (
	UserNotFoundError      = errors.New("user not found")
	UserAlreadyExistsError = errors.New("user already exists")
)
