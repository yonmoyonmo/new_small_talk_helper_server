package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yonmoyonmo/new_small_talk_helper_server/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	log.Println("new small talk helper server")
	mux := router.InitializeRouter()
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(port, mux))
}
