package model

import (
	"time"
)

type Sugguestion struct {
	Id              int       `json:"id" db:"id"`
	SugguestionType string    `json:"sugguestion_type" db:"sugguestion_type"`
	SuggustionText  string    `json:"sugguestion_text" db:"sugguestion_text"`
	CountLike       int       `json:"count_likes" db:"count_likes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

func (s *Sugguestion) InitTimeNow() {
	s.CreatedAt = time.Now()
}
