package models

import "time"

type Movie struct {
	ID          int       `json:"id" gorm:"primaryKey:autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Genres      string    `json:"genres"`
	Url         string    `json:"url"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
	CreatedAT   time.Time `json:"created_at" gorm:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at"`
}
