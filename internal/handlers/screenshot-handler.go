package handlers

import (
	"encoding/json"
	"go-screenshot/internal/models"
	"go-screenshot/internal/utils"
	"net/http"
)

func HandleScreenshot(w http.ResponseWriter, r *http.Request) {
	// validate request method
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(utils.APIResponse(models.StatusFailure, "method not allowed", nil))
		return
	}

	// Query Paramas
	deviceStr := r.URL.Query().Get("device")
	urlStr := r.URL.Query().Get("url")

	// default device to desktop
	if deviceStr == "" {
		deviceStr = models.DeviceDesktop
	}

	// validate url
	if !utils.ValidateURL(urlStr) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.APIResponse(models.StatusFailure, "invalid url", nil))
		return
	}

	// validate device
	if deviceStr != models.DeviceDesktop && deviceStr != models.DeviceMobile {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.APIResponse(models.StatusFailure, "invalid device: supported devices include "+models.DeviceDesktop+" and "+models.DeviceMobile, nil))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.APIResponse(models.StatusSuccess, "screenshot taken successfully for "+urlStr+", device: "+deviceStr, nil))
}
