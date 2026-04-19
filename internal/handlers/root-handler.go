package handlers

import (
	"encoding/json"
	"go-screenshot/internal/models"
	"go-screenshot/internal/utils"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.APIResponse(models.StatusSuccess, "server is live", nil))
}
