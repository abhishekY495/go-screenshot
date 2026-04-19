package handlers

import (
	"encoding/json"
	"go-screenshot/internal/models"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.RootResponse{
		Status:  models.StatusSuccess,
		Message: "server is live",
	})
}
