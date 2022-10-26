package model

type Admin struct {
	Id        int    `json:"id" db:"id"`
	AdminName string `json:"admin_name" db:"admin_name"`
	Password  string `json:"password" db:"password"`
}

type AdminToken struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
