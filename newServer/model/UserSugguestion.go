package model

import (
	"time"
)

type UserSugguestion struct {
	Id        int       `json:"id" db:"id"`
	UserName  string    `json:"userName" db:"user_name"`
	Text      string    `json:"text" db:"text"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (s *UserSugguestion) InitTimeNow() *UserSugguestion {
	newSuggestion := UserSugguestion{}
	newSuggestion.CreatedAt = time.Now()
	return &newSuggestion
}
