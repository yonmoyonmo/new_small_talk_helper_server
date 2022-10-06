package handler

import (
	"log"
	"net/http"

	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
)

func TestHandler(resWriter http.ResponseWriter, req *http.Request) {
	var newSuggestion model.Sugguestion = model.Sugguestion{}
	newSuggestion = *newSuggestion.InitTimeNow()
	log.Printf("%v, %T", newSuggestion, newSuggestion)
}
