package models

import "time"

type MovieVote struct {
	ID        int       `json:"id" gorm:"primaryKey:autoIncrement"`
	IDMovie   int       `json:"id_movie"`
	IDUser    int       `json:"id_user"`
	Vote      int       `json:"vote"`
	CreatedBy int       `json:"created_by"`
	UpdatedBy int       `json:"updated_by"`
	CreatedAT time.Time `json:"created_at" gorm:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (MovieVote) TableName() string {
	return "movies_vote"
}
