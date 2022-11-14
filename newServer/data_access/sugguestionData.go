package dataaccess

import (
	"log"

	"github.com/yonmoyonmo/new_small_talk_helper_server/dbconn"
	"github.com/yonmoyonmo/new_small_talk_helper_server/model"
)

func GetRandomSuggData() model.Sugguestion {
	db := dbconn.GetDBConnector()
	var rds model.Sugguestion

	var whereCondition string = "WHERE sugguestion_type != 'love36'"
	var randomSuggSQL string = "SELECT * FROM sugguestion " + whereCondition + " ORDER BY RAND() LIMIT 1"

	err := db.QueryRow(randomSuggSQL).Scan(&rds.Id, &rds.SugguestionType, &rds.SuggustionText, &rds.CountLike, &rds.CreatedAt)
	if err != nil {
		log.Println("GetRandomSuggData error")
		log.Println(err)
		db.Close()
		return rds
	}
	db.Close()
	return rds
}

func GetLove36Data() []model.Sugguestion {
	db := dbconn.GetDBConnector()
	var resultList []model.Sugguestion

	var love36SQL string = "SELECT * FROM sugguestion WHERE sugguestion_type = 'love36'"

	rows, err := db.Query(love36SQL)
	if err != nil {
		log.Println("GetLove36Data error")
		log.Println(err)
		db.Close()
		return resultList
	}
	for rows.Next() {
		s := model.Sugguestion{}
		err := rows.Scan(&s.Id, &s.SugguestionType, &s.SuggustionText, &s.CountLike, &s.CreatedAt)
		if err != nil {
			log.Printf("could not scan row: %v", err)
			db.Close()
			return resultList
		}
		resultList = append(resultList, s)
	}
	db.Close()
	return resultList
}

func GetToptenData() []model.Sugguestion {
	db := dbconn.GetDBConnector()
	var resultList []model.Sugguestion

	var toptenSQL string = "SELECT *  FROM sugguestion ORDER BY count_likes DESC LIMIT 10;"

	rows, err := db.Query(toptenSQL)
	if err != nil {
		log.Println("GetToptenData")
		log.Println(err)
		db.Close()
		return resultList
	}
	for rows.Next() {
		s := model.Sugguestion{}
		err := rows.Scan(&s.Id, &s.SugguestionType, &s.SuggustionText, &s.CountLike, &s.CreatedAt)
		if err != nil {
			log.Printf("could not scan row: %v", err)
			db.Close()
			return resultList
		}
		resultList = append(resultList, s)
	}
	db.Close()
	return resultList
}

func GetFavoriteData(ids string) []model.Sugguestion {
	db := dbconn.GetDBConnector()
	var resultList []model.Sugguestion
	var whereCondition string = "WHERE id IN (" + ids + ")"

	var favoriteSQL string = "SELECT * FROM sugguestion " + whereCondition
	rows, err := db.Query(favoriteSQL)
	if err != nil {
		log.Println("GetToptenData")
		log.Println(err)
		db.Close()
		return resultList
	}
	for rows.Next() {
		s := model.Sugguestion{}
		err := rows.Scan(&s.Id, &s.SugguestionType, &s.SuggustionText, &s.CountLike, &s.CreatedAt)
		if err != nil {
			log.Printf("could not scan row: %v", err)
			db.Close()
			return resultList
		}
		resultList = append(resultList, s)
	}
	db.Close()
	return resultList
}

func UpdateLikes(suggId int, likeValue int) bool {
	db := dbconn.GetDBConnector()
	log.Println(suggId, likeValue)

	_, err := db.Exec("UPDATE sugguestion SET count_likes = count_likes + ? WHERE id= ?", likeValue, suggId)
	if err != nil {
		log.Println("UpdateLikes : something went wrong")
		log.Println(err)
		db.Close()
		return false
	}
	db.Close()
	return true
}

func InsertNewUserSugg(userName string, text string) bool {
	db := dbconn.GetDBConnector()
	log.Println(userName, text)
	_, err := db.Exec("INSERT INTO user_sugguestion(user_name, text) value (?, ?)", userName, text)
	if err != nil {
		log.Println("InsertNewUserSugg : something went wrong")
		log.Println(err)
		db.Close()
		return false
	}
	db.Close()
	return true
}
