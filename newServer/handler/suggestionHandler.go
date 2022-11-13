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

func FavoriteHandler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var favoriteIdList model.FavoriteIdList
		json.NewDecoder(req.Body).Decode(&favoriteIdList)
		var favoriteSuggList []model.Sugguestion = service.FavoriteService(favoriteIdList)
		if favoriteSuggList == nil {
			var noFav model.Sugguestion
			noFav.SugguestionType = "error"
			noFav.SuggustionText = "등록된 즐겨찾기가 없습니다."
			resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(resWriter).Encode(noFav)
		} else {
			resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(resWriter).Encode(favoriteSuggList)
		}
	}
}

func LikesHandler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var likes model.Liikes
		json.NewDecoder(req.Body).Decode(&likes)
		var result bool = service.ApplyLikes(likes)
		resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(resWriter).Encode(result)
	}
}

func UserSugguestionHandler(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var newUserSugg model.UserSugguestion = model.UserSugguestion{}
		newUserSugg = *newUserSugg.InitTimeNow()
		json.NewDecoder(req.Body).Decode(&newUserSugg)
		log.Println(newUserSugg)
		result := service.CreateNewUserSugg(newUserSugg)
		resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(resWriter).Encode(result)
	}
}
