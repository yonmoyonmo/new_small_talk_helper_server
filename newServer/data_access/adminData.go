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
		db.Close()
		return resultAdmin
	} else {
		log.Println(resultAdmin)
		db.Close()
		return resultAdmin
	}
}

func InsertNewAdmin(adminName string, password string) bool {
	db := dbconn.GetDBConnector()
	_, err := db.Exec("INSERT INTO admin(admin_name, password) value (?, ?)", adminName, password)
	if err != nil {
		log.Println(err)
		db.Close()
		return false
	} else {
		db.Close()
		return true
	}
}

func InsertNewSugguestion(newSugg model.Sugguestion) bool {
	db := dbconn.GetDBConnector()
	_, err := db.Exec("INSERT INTO sugguestion(sugguestion_text, sugguestion_type) value(?, ?)", newSugg.SuggustionText, newSugg.SugguestionType)
	if err != nil {
		log.Println("insert new sugguestion error")
		log.Println(err)
		log.Println(newSugg)
		db.Close()
		return false
	} else {
		db.Close()
		return true
	}
}

func UpdateSugguestion(targetSugg model.Sugguestion) bool {
	db := dbconn.GetDBConnector()
	_, err := db.Exec("UPDATE sugguestion SET sugguestion_text = ?,  sugguestion_type = ? WHERE id= ? ", targetSugg.SuggustionText, targetSugg.SugguestionType, targetSugg.Id)
	if err != nil {
		log.Println("update sugguestion error")
		log.Println(err)
		log.Println(targetSugg)
		db.Close()
		return false
	} else {
		db.Close()
		return true
	}
}

func DeleteSugguestion(targetSugg model.Sugguestion) bool {
	db := dbconn.GetDBConnector()
	_, err := db.Exec("DELETE FROM sugguestion WHERE id= ? ", targetSugg.Id)
	if err != nil {
		log.Println("delete sugguestion error")
		log.Println(err)
		log.Println(targetSugg)
		db.Close()
		return false
	} else {
		db.Close()
		return true
	}
}

func ReadSugguestionList(limit int, offset int) []model.Sugguestion {
	db := dbconn.GetDBConnector()
	var resultList []model.Sugguestion

	var suggList string = "SELECT *  FROM sugguestion ORDER BY id DESC LIMIT ? OFFSET ?;"

	rows, err := db.Query(suggList, limit, offset)
	if err != nil {
		log.Println("Select sugguestionserror")
		log.Println(err)
		db.Close()
		return resultList
	}
	for rows.Next() {
		s := model.Sugguestion{}
		err := rows.Scan(&s.Id, &s.SugguestionType, &s.SuggustionText, &s.CountLike, &s.CreatedAt)
		if err != nil {
			log.Println("Select sugguestions error")
			log.Println(err)
			db.Close()
			return resultList
		}
		resultList = append(resultList, s)
	}
	db.Close()
	return resultList
}

func ReadUserSugguestionList(limit int, offset int) []model.UserSugguestion {
	db := dbconn.GetDBConnector()
	var resultList []model.UserSugguestion

	var suggList string = "SELECT *  FROM user_sugguestion ORDER BY id DESC LIMIT ? OFFSET ?"

	rows, err := db.Query(suggList, limit, offset)
	if err != nil {
		log.Println("Select user sugguestions error")
		log.Println(err)
		db.Close()
		return resultList
	}
	for rows.Next() {
		s := model.UserSugguestion{}
		err := rows.Scan(&s.Id, &s.Text, &s.UserName, &s.CreatedAt)
		if err != nil {
			log.Println("Select user sugguestionserror")
			log.Println(err)
			db.Close()
			return resultList
		}
		resultList = append(resultList, s)
	}
	db.Close()
	return resultList
}
