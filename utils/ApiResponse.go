package utils

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewApiResponse(success bool, message string, data interface{}) *ApiResponse {
	return &ApiResponse{Success: success, Message: message, Data: data}
}
