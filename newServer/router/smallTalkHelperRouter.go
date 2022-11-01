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
	//small talk helper service
	smthpBaseURL := "/api/sugguestion/small-talk-helper"
	mux.HandleFunc(smthpBaseURL+"/random", handler.RandomSuggHandler)
	mux.HandleFunc(smthpBaseURL+"/love36", handler.Love36Handler)
	mux.HandleFunc(smthpBaseURL+"/topten", handler.ToptenHandler)
	mux.HandleFunc(smthpBaseURL+"/favorite", handler.FavoriteHandler)
	mux.HandleFunc(smthpBaseURL+"/likes", handler.LikesHandler)
	mux.HandleFunc(smthpBaseURL+"/users-sugguestion", handler.UserSugguestionHandler)
	//small talk helper donator
	donatorBase := "/api/donator"
	mux.HandleFunc(donatorBase+"/register", handler.RegisterNewDonatorHandler)
	mux.HandleFunc(donatorBase+"/check", handler.CheckDonatorHandler)
	//admin
	mux.HandleFunc("/api/admin/login", handler.AdminLoginHandler)
	mux.HandleFunc("/api/admin/register", handler.AdminRegisterHandler)
	//admin sugguestion management
	mux.HandleFunc("/api/sugguestion/register/multiple", handler.RegisterSugguestionsHandler)
	mux.HandleFunc("/api/sugguestion/update", handler.UpdateSugguestionHandler)
	mux.HandleFunc("/api/sugguestion/delete", handler.DeleteSugguestionHandler)
	mux.HandleFunc("/api/sugguestion/list", handler.GetSugguestionListHandler)
	mux.HandleFunc("/api/sugguestion/usersugguestions", handler.GetUserSugguestionListHandler)
	//done
	log.Println("routing done")
	return mux
}
