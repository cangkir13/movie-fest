package movievoterepository

import (
	"context"

	"movie-festival-app/src/models"

	"gorm.io/gorm"
)

type movieVoteRepository struct {
	mysql *gorm.DB
}

func NewMovieVoteRepositories(db *gorm.DB) *movieVoteRepository {
	return &movieVoteRepository{mysql: db}
}

type MovieVoteRepository interface {
	Create(ctx context.Context, props models.MovieVote) error
	GetOne(ctx context.Context, props GetOneProps) (record *models.MovieVote, err error)
	Update(ctx context.Context, props UpdateProps) error
}
