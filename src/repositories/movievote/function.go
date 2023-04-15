package movievoterepository

import (
	"context"

	"movie-festival-app/src/models"

	"gorm.io/gorm"
)

func (db *movieVoteRepository) Create(ctx context.Context, props models.MovieVote) error {
	query := db.mysql.WithContext(ctx)
	if err := query.Create(&props).Error; err != nil {
		return err
	}

	return nil
}

type GetOneProps struct {
	ID      *int
	IDMovie *int
	IDUser  *int
}

func (db *movieVoteRepository) GetOne(ctx context.Context, props GetOneProps) (record *models.MovieVote, err error) {
	query := db.mysql

	if props.ID != nil {
		query = query.Where("id = ?", props.ID)
	}

	if props.IDMovie != nil {
		query = query.Where("id_movie = ?", props.IDMovie)
	}

	if props.IDUser != nil {
		query = query.Where("id_user = ?", props.IDUser)
	}

	if err := query.WithContext(ctx).First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return record, nil
}

type UpdateProps struct {
	ID     int
	Fields []string
	Data   models.MovieVote
}

func (db *movieVoteRepository) Update(ctx context.Context, props UpdateProps) error {
	query := db.mysql.WithContext(ctx).
		Select(props.Fields).
		Where("id = ?", props.ID)

	if err := query.Updates(&props.Data).Error; err != nil {
		return err
	}

	return nil
}
