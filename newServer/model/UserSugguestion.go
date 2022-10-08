package model

import (
	"time"
)

type UserSugguestion struct {
	Id        int       `json:"id" db:"id"`
	Text      string    `json:"text" db:"text"`
	UserName  string    `json:"user_name" db:"user_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (suggestion UserSugguestion) initTimeNow() *UserSugguestion {
	newSuggestion := UserSugguestion{}
	newSuggestion.CreatedAt = time.Now()
	return &newSuggestion
}
