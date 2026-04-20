package handlers

import (
	"go-screenshot/internal/models"
	"go-screenshot/internal/utils"
	"net/http"
)

func HandleScreenshot(w http.ResponseWriter, r *http.Request) {
	// validate request method
	if r.Method != http.MethodGet {
		utils.JsonError(w, http.StatusMethodNotAllowed, "method not allowed")
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
		utils.JsonError(w, http.StatusBadRequest, "invalid url")
		return
	}

	// validate device
	if deviceStr != models.DeviceDesktop && deviceStr != models.DeviceMobile {
		utils.JsonError(w, http.StatusBadRequest, "invalid device: supported devices include "+models.DeviceDesktop+" and "+models.DeviceMobile)
		return
	}

	screenshotBytes, err := utils.CaptureScreenshot(urlStr, deviceStr)
	if err != nil {
		utils.JsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(screenshotBytes)
}
