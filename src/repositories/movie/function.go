package movierepository

import (
	"context"

	"movie-festival-app/src/models"

	"gorm.io/gorm"
)

type ListProps struct {
	Page        int
	Limit       int
	Title       *string
	Description *string
	Artists     *string
	Genres      *string
}

func (db *movieRepository) List(ctx context.Context, props ListProps) (records []*models.Movie, err error) {
	query := db.mysql

	if props.Limit != 0 || props.Page != 0 {
		offset := (props.Page - 1) * props.Limit
		query = query.Limit(props.Limit).Offset(offset)
	}

	if props.Title != nil {
		query = query.Where("title LIKE ?", "%"+*props.Title+"%")
	}

	if props.Artists != nil {
		query = query.Or("artists LIKE ?", "%"+*props.Artists+"%")
	}

	if props.Description != nil {
		query = query.Or("description LIKE ?", "%"+*props.Description+"%")
	}

	if props.Genres != nil {
		query = query.Or("genres LIKE ?", "%"+*props.Genres+"%")
	}

	// execute
	if err = query.WithContext(ctx).Debug().Find(&records).Error; err != nil {
		return
	}

	return
}

func (db *movieRepository) Create(ctx context.Context, props models.Movie) error {
	query := db.mysql.WithContext(ctx)
	if err := query.Create(&props).Error; err != nil {
		return err
	}

	return nil
}

type UpdateProps struct {
	ID     int
	Fields []string
	Data   models.Movie
}

func (db *movieRepository) Update(ctx context.Context, props UpdateProps) error {
	query := db.mysql.WithContext(ctx).
		Select(props.Fields).
		Where("id = ?", props.ID)

	if err := query.Updates(&props.Data).Error; err != nil {
		return err
	}

	return nil
}

func (db *movieRepository) GetOne(ctx context.Context, ID int) (record *models.Movie, err error) {
	query := db.mysql.Where("id = ?", ID)

	if err := query.WithContext(ctx).First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return record, nil
}
