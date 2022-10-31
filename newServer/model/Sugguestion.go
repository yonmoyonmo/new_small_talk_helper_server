package model

import (
	"time"
)

type Sugguestion struct {
	Id              int       `json:"id" db:"id"`
	SugguestionType string    `json:"sugguestionText" db:"sugguestion_text"`
	SuggustionText  string    `json:"type" db:"sugguestion_type"`
	CountLike       int       `json:"count_likes" db:"count_likes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

type Sugguestions []Sugguestion

func (s *Sugguestion) InitTimeNow() {
	s.CreatedAt = time.Now()
}
