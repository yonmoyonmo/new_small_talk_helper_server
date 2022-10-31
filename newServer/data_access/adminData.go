package dataaccess

import (
	"log"

	"github.com/yonmoyonmo/new_small_talk_helper_server/dbconn"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
)

func GetAdmin(adminName string) model.Admin {
	db := dbconn.GetDBConnector()
	var resultAdmin model.Admin
	err := db.QueryRow("SELECT * FROM admin WHERE admin_name = ?", adminName).Scan(&resultAdmin.Id, &resultAdmin.AdminName, &resultAdmin.Password)
	if err != nil {
		log.Println("no matched admin name")
		log.Println(err)
		resultAdmin.Id = 0
		return resultAdmin
	} else {
		log.Println(resultAdmin)
		return resultAdmin
	}
}

func InsertNewAdmin(adminName string, password string) bool {
	db := dbconn.GetDBConnector()
	_, err := db.Exec("INSERT INTO admin(admin_name, password) value (?, ?)", adminName, password)
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func InsertNewSugguestion(newSugg model.Sugguestion) bool {
	db := dbconn.GetDBConnector()
	_, err := db.Exec("INSERT INTO sugguestion(sugguestion_text, sugguestion_type) value(?, ?)", newSugg.SuggustionText, newSugg.SugguestionType)
	if err != nil {
		log.Println("insert new sugguestion error")
		log.Println(err)
		return false
	} else {
		return true
	}
}
