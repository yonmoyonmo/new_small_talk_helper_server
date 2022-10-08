package model

import (
	"log"
	"time"
)

type Sugguestion struct {
	Id              int
	SugguestionType string
	SuggustionText  string
	CountLike       int
	CreatedAt       time.Time
}

func (suggestion Sugguestion) InitTimeNow() *Sugguestion {
	newSuggestion := Sugguestion{}
	newSuggestion.CreatedAt = time.Now()
	log.Printf("%T", newSuggestion)
	return &newSuggestion
}
