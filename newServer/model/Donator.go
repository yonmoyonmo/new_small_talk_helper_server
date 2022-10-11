package model

type Donator struct {
	Id          int    `json:"id" db:"id"`
	DonatorName string `json:"donator_name" db:"donator_name"`
	Password    string `json:"password" db:"password"`
}

type DonationMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
