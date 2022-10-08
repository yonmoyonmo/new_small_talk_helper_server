package service

import (
	"log"

	dataaccess "github.com/yonmoyonmo/new_small_talk_helper_server/data_access"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
)

func RandomSuggService() model.Sugguestion {
	log.Println("I'm random sugg service")
	result := dataaccess.RandomSuggData()
	return result
}

func Love36Service() []model.Sugguestion {
	result := dataaccess.Love36Data()
	return result
}

func ToptenService() []model.Sugguestion {
	result := dataaccess.ToptenData()
	return result
}
