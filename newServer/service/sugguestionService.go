package service

import (
	"log"

	dataaccess "github.com/yonmoyonmo/new_small_talk_helper_server/data_access"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
)

func RandomSuggService() model.Sugguestion {
	log.Println("I'm random sugg serv")
	result := dataaccess.RandomSuggData()
	return result
}
