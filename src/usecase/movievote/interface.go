package movievoteusecase

import (
	"context"

	movierepository "movie-festival-app/src/repositories/movie"
	movievoterepository "movie-festival-app/src/repositories/movievote"
)

type movieVoteUsecase struct {
	movierepository     movierepository.MovieRepository
	movievotereposiotry movievoterepository.MovieVoteRepository
}

func NewMovieVoteUsecase(
	movierepository movierepository.MovieRepository,
	movievotereposiotry movievoterepository.MovieVoteRepository,
) *movieVoteUsecase {
	return &movieVoteUsecase{movierepository: movierepository, movievotereposiotry: movievotereposiotry}
}

type MovieVoteUsecase interface {
	UserVote(ctx context.Context, props UserVoteProps) error
	UserUnVote(ctx context.Context, props UserUnVoteProps) error
}
