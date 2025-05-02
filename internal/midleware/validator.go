package middleware

import (
	"github.com/go-playground/validator/v10"
)

// Validator представляет валидатор запросов
type Validator struct {
	validator *validator.Validate
}

// NewValidator создает новый валидатор
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate валидирует структуру
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
