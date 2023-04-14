package movieusecase

import (
	"context"

	"movie-festival-app/src/models"
	movierepository "movie-festival-app/src/repositories/movie"
)

type movieUsecase struct {
	movierepository movierepository.MovieRepository
}

func NewMovieUsecase(movierepository movierepository.MovieRepository) *movieUsecase {
	return &movieUsecase{movierepository: movierepository}
}

type MovieUsecase interface {
	GetList(ctx context.Context, props GetListProps) (records []*models.Movie, err error)
	CreateNew(ctx context.Context, props models.Movie) error
	UpdateMovie(ctx context.Context, props UpdateMovieProps) error
}
