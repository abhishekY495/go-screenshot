package utils

import "go-screenshot/internal/models"

func GetWidthHeight(deviceStr string) (int64, int64) {
	if deviceStr == models.DeviceMobile {
		return 430, 932
	}
	return 1920, 1080
}
