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
		log.Printf("%v", randomSugg)
		json.NewEncoder(resWriter).Encode(randomSugg)
	}

}
