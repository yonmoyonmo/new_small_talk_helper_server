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
	//small talk helper
	smthpBaseURL := "/api/sugguestion/small-talk-helper"
	mux.HandleFunc(smthpBaseURL+"/random", handler.RandomSuggHandler)
	mux.HandleFunc(smthpBaseURL+"/love36", handler.Love36Handler)
	mux.HandleFunc(smthpBaseURL+"/topten", handler.ToptenHandler)
	mux.HandleFunc(smthpBaseURL+"/favorite", handler.FavoriteHandler)
	mux.HandleFunc(smthpBaseURL+"/likes", handler.LikesHandler)
	//admin
	//etc
	log.Println("router is ready... i think...")
	return mux
}
