package models

const DeviceDesktop string = "desktop"
const DeviceMobile string = "mobile"

type ScreenshotRequest struct {
	Url    string `json:"url"`
	Device string `json:"device"`
}
