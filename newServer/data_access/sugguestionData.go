package dataaccess

import (
	"log"

	"github.com/yonmoyonmo/new_small_talk_helper_server/dbconn"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
)

func RandomSuggData() model.Sugguestion {
	db := dbconn.GetDBConnector()
	var rds model.Sugguestion

	var whereCondition string = "WHERE sugguestion_type != 'love36'"
	var randomSuggSQL string = "SELECT * FROM sugguestion " + whereCondition + " ORDER BY RAND() LIMIT 1"

	err := db.QueryRow(randomSuggSQL).Scan(&rds.Id, &rds.SugguestionType, &rds.SuggustionText, &rds.CountLike, &rds.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	return rds
}

func Love36Data() []model.Sugguestion {
	db := dbconn.GetDBConnector()
	var resultList []model.Sugguestion

	var love36SQL string = "SELECT * FROM sugguestion WHERE sugguestion_type = 'love36'"

	rows, err := db.Query(love36SQL)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		s := model.Sugguestion{}
		err := rows.Scan(&s.Id, &s.SugguestionType, &s.SuggustionText, &s.CountLike, &s.CreatedAt)
		if err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		resultList = append(resultList, s)
	}
	return resultList
}

func ToptenData() []model.Sugguestion {
	db := dbconn.GetDBConnector()
	var resultList []model.Sugguestion

	var toptenSQL string = "SELECT *  FROM sugguestion ORDER BY count_likes DESC LIMIT 10;"

	rows, err := db.Query(toptenSQL)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		s := model.Sugguestion{}
		err := rows.Scan(&s.Id, &s.SugguestionType, &s.SuggustionText, &s.CountLike, &s.CreatedAt)
		if err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		resultList = append(resultList, s)
	}
	return resultList
}

func FavoriteData(ids string) []model.Sugguestion {
	db := dbconn.GetDBConnector()
	var resultList []model.Sugguestion
	var whereCondition string = "WHERE id IN (" + ids + ")"

	var favoriteSQL string = "SELECT * FROM sugguestion " + whereCondition
	rows, err := db.Query(favoriteSQL)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		s := model.Sugguestion{}
		err := rows.Scan(&s.Id, &s.SugguestionType, &s.SuggustionText, &s.CountLike, &s.CreatedAt)
		if err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		resultList = append(resultList, s)
	}
	return resultList
}
