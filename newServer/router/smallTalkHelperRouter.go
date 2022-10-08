package router

import (
	"log"
	"net/http"

	"github.com/yonmoyonmo/new_small_talk_helper_server/handler"
)

func InitializeRouter() *http.ServeMux {
	log.Println("initializing router...")
	mux := http.NewServeMux()
	//routes
	mux.HandleFunc("/test", handler.TestHandler)
	//random sugguestion
	smthpBaseURL := "/api/sugguestion/small-talk-helper"
	mux.HandleFunc(smthpBaseURL+"/random", handler.RandomSuggHandler)
	//
	log.Println("router is ready... i think...")
	return mux
}
