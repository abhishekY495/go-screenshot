package utils

import (
	"encoding/json"
	"go-screenshot/internal/models"
	"net/http"
)

func JsonError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(APIResponse(models.StatusError, message, nil))
}
