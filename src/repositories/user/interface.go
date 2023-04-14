package userrepository

import (
	"context"

	"movie-festival-app/src/models"

	"gorm.io/gorm"
)

type userRepositories struct {
	mysql *gorm.DB
}

func NewUserRepositories(db *gorm.DB) *userRepositories {
	return &userRepositories{mysql: db}
}

type UserRepositories interface {
	GetOne(ctx context.Context, props string) (record *models.User, err error)
	Create(ctx context.Context, props CreateProps) error
}
