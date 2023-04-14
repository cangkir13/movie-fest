package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey:autoIncrement"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	RoleID    int       `json:"role_id" gorm:"DEFAULT:2"`
	IsAllow   bool      `json:"is_allow" gorm:"DEFAULT:true"`
	CreatedAT time.Time `json:"created_at" gorm:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAT time.Time `json:"updated_at" gorm:"DEFAULT:CURRENT_TIMESTAMP"`
}
