package model

type SimpleResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
