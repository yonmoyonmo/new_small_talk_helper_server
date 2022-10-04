package main

import (
	"log"

	dbconn "github.com/yonmoyonmo/new_small_talk_helper_server/dbConn"
)

func main() {
	log.Println("new SMTHP server")
	dbconn.TestLog("sibal")
	db := dbconn.GetDBConnector()

	var connectionTest string
	err := db.QueryRow("SELECT sugguestion_text FROM sugguestion WHERE id =12").Scan(&connectionTest)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(connectionTest)
	print("success")
}
