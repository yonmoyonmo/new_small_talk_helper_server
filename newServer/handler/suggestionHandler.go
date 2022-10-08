package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
	"github.com/yonmoyonmo/new_small_talk_helper_server/service"
)

func RandomSuggHandler(resWriter http.ResponseWriter, req *http.Request) {
	log.Println("Im random handler")
	switch req.Method {
	case http.MethodGet:
		var randomSugg model.Sugguestion = service.RandomSuggService()
		resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(resWriter).Encode(randomSugg)
	}
}

func Love36Handler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		var love36List []model.Sugguestion = service.Love36Service()
		resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(resWriter).Encode(love36List)
	}
}

func ToptenHandler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		var toptenList []model.Sugguestion = service.ToptenService()
		resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(resWriter).Encode(toptenList)
	}
}
