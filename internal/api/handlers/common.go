package handlers

// ErrorResponse представляет ответ с ошибкой
type ErrorResponse struct {
	Message string `json:"message"`
}

// SuccessResponse представляет успешный ответ
type SuccessResponse struct {
	Message string `json:"message"`
}
