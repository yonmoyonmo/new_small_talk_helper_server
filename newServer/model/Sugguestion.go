package model

import (
	"log"
	"time"
)

type Sugguestion struct {
	id              int
	sugguestionType string
	suggustionText  string
	countLike       int
	createdAt       time.Time
}

func (suggestion Sugguestion) InitTimeNow() *Sugguestion {
	newSuggestion := Sugguestion{}
	newSuggestion.createdAt = time.Now()
	log.Printf("%T", newSuggestion)
	return &newSuggestion
}
