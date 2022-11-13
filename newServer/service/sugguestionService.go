package service

import (
	"log"
	"strings"

	dataaccess "github.com/yonmoyonmo/new_small_talk_helper_server/data_access"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
)

func RandomSuggService() model.Sugguestion {
	log.Println("I'm random sugg service")
	result := dataaccess.GetRandomSuggData()
	return result
}

func Love36Service() []model.Sugguestion {
	result := dataaccess.GetLove36Data()
	return result
}

func ToptenService() []model.Sugguestion {
	result := dataaccess.GetToptenData()
	return result
}

func FavoriteService(idList model.FavoriteIdList) []model.Sugguestion {
	var idsString string = strings.Join(idList.FavoriteIdList, ",")
	if idsString == "" {
		log.Println("no favorites")
		return nil
	}
	result := dataaccess.GetFavoriteData(idsString)
	return result
}

func ApplyLikes(likes model.Liikes) bool {
	result := dataaccess.UpdateLikes(likes.SugguestionId, likes.LikeValue)
	return result
}

func CreateNewUserSugg(newUserSugg model.UserSugguestion) bool {
	return dataaccess.InsertNewUserSugg(newUserSugg.UserName, newUserSugg.Text)
}
