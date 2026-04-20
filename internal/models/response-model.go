package models

const StatusSuccess string = "success"
const StatusError string = "error"

type APIResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
