package model

type Adimin struct {
	Id        int    `json:"id" db:"id"`
	AdminName string `json:"admin_name" db:"admin_name"`
	Password  string `json:"password" db:"password"`
}
