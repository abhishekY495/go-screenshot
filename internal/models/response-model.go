package models

const StatusSuccess string = "success"
const StatusFailure string = "failure"

type APIResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
