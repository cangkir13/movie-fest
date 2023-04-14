package userrepository

import (
	"context"

	"movie-festival-app/src/models"

	"gorm.io/gorm"
)

func (db *userRepositories) GetOne(ctx context.Context, props string) (record *models.User, err error) {
	query := db.mysql.Where("username = ?", props)
	if errService := query.WithContext(ctx).First(&record).Error; errService != nil {
		if errService == gorm.ErrRecordNotFound {
			err = nil
			record = nil
			return
		}

		err = errService
		return
	}

	return
}

type CreateProps struct {
	Data models.User
}

func (db *userRepositories) Create(ctx context.Context, props CreateProps) error {
	query := db.mysql

	if err := query.WithContext(ctx).Create(&props.Data).Error; err != nil {
		return err
	}
	return nil
}
