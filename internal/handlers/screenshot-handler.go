package handlers

import (
	"fmt"
	"go-screenshot/internal/models"
	"go-screenshot/internal/utils"
	"net/http"
	"time"
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
	ttlStr := r.URL.Query().Get("ttl")

	// default device to desktop
	if deviceStr == "" {
		deviceStr = models.DeviceDesktop
	}

	// default ttl to 1 day
	if ttlStr == "" {
		ttlStr = "1d"
	}

	// validate ttl
	ttl, err := utils.ValidateTtl(ttlStr)
	if err != nil {
		utils.JsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	// validate url
	if !utils.ValidateURL(urlStr) {
		utils.JsonError(w, http.StatusBadRequest, "invalid url")
		return
	}

	// validate device
	if deviceStr != models.DeviceDesktop && deviceStr != models.DeviceMobile {
		utils.JsonError(w, http.StatusBadRequest, "invalid device - supported devices include "+models.DeviceDesktop+" and "+models.DeviceMobile)
		return
	}

	screenshotBytes, err := utils.CaptureScreenshot(models.ScreenshotRequest{
		Url:    urlStr,
		Device: deviceStr,
	})
	if err != nil {
		utils.JsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ttlSeconds := int(ttl.Seconds())

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", ttlSeconds))
	w.Header().Set("Expires", time.Now().Add(ttl).UTC().Format(http.TimeFormat))

	w.WriteHeader(http.StatusOK)
	w.Write(screenshotBytes)
}
