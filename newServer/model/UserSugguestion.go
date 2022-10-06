package model

import (
	"time"
)

type UserSugguestion struct {
	id        int
	text      string
	userName  string
	createdAt time.Time
}

func (suggestion UserSugguestion) initTimeNow() *UserSugguestion {
	newSuggestion := UserSugguestion{}
	newSuggestion.createdAt = time.Now()
	return &newSuggestion
}
