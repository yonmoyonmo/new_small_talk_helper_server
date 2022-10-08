package main

import (
	"log"
	"net/http"

	"github.com/yonmoyonmo/new_small_talk_helper_server/router"
)

func main() {
	log.Println("new SMTHP server")
	mux := router.InitializeRouter()
	log.Println("Listing... at 8888")
	log.Fatal(http.ListenAndServe(":8888", mux))
}
