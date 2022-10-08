package main

import (
	"log"
	"net/http"

	"github.com/yonmoyonmo/new_small_talk_helper_server/router"
)

func main() {
	log.Println("new SMTHP server")
	mux := router.InitializeRouter()
	// db := dbconn.GetDBConnector()

	// var connectionTest string
	// err := db.QueryRow("SELECT sugguestion_text FROM sugguestion WHERE id =12").Scan(&connectionTest)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(connectionTest)
	// print("success")
	log.Fatal(http.ListenAndServe(":8888", mux))
}
