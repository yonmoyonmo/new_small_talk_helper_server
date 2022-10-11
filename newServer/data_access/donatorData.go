package dataaccess

import (
	"log"

	"github.com/yonmoyonmo/new_small_talk_helper_server/dbconn"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
)

func InsertNewDonator(donatorName string, hashedPassword string) bool {
	db := dbconn.GetDBConnector()
	_, err := db.Exec("INSERT INTO donator(donator_name, password) value (?, ?)", donatorName, hashedPassword)
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func GetDonator(donatorName string) model.Donator {
	db := dbconn.GetDBConnector()
	var resultDonator model.Donator
	err := db.QueryRow("SELECT * FROM donator WHERE donator_name = ?", donatorName).Scan(&resultDonator.Id, &resultDonator.DonatorName, &resultDonator.Password)
	if err != nil {
		log.Println(err)
		resultDonator.Id = 0
		return resultDonator
	} else {
		log.Println(resultDonator)
		return resultDonator
	}
}
