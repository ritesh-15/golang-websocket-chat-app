package utils

type ApiError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewApiError(success bool, message string) *ApiError {
	return &ApiError{Success: success, Message: message}
}
