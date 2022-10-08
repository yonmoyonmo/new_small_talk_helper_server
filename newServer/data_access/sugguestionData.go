package dataaccess

import (
	"log"

	"github.com/yonmoyonmo/new_small_talk_helper_server/dbconn"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
)

func RandomSuggData() model.Sugguestion {
	db := dbconn.GetDBConnector()
	var rds model.Sugguestion
	var whereConditions string = "WHERE sugguestion_type != 'love36'"
	err := db.QueryRow("SELECT * FROM sugguestion "+whereConditions+" ORDER BY RAND() LIMIT 1").Scan(&rds.Id, &rds.SugguestionType, &rds.SuggustionText, &rds.CountLike, &rds.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rds)
	return rds
}
