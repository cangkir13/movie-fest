package movierepository

import (
	"context"

	"movie-festival-app/src/models"

	"gorm.io/gorm"
)

type movieRepository struct {
	mysql *gorm.DB
}

func NewMovieRepositories(db *gorm.DB) *movieRepository {
	return &movieRepository{mysql: db}
}

type MovieRepository interface {
	List(ctx context.Context, props ListProps) (records []*models.Movie, err error)
	Create(ctx context.Context, props models.Movie) error
	Update(ctx context.Context, props UpdateProps) error
	GetOne(ctx context.Context, ID int) (record *models.Movie, err error)
}
