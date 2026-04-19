package utils

import "go-screenshot/internal/models"

func APIResponse(status string, message string, data any) models.RootResponse {
	return models.RootResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
