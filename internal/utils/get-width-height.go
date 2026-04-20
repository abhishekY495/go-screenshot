package utils

import "go-screenshot/internal/models"

func GetWidthHeight(deviceStr string) (int64, int64) {
	if deviceStr == models.DeviceMobile {
		return MobileWidth, MobileHeight
	}
	return DesktopWidth, DesktopHeight
}
