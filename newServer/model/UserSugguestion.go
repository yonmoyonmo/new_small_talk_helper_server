package model

import (
	"time"
)

type UserSugguestion struct {
	Id        int
	Text      string
	UserName  string
	CreatedAt time.Time
}

func (suggestion UserSugguestion) initTimeNow() *UserSugguestion {
	newSuggestion := UserSugguestion{}
	newSuggestion.CreatedAt = time.Now()
	return &newSuggestion
}
