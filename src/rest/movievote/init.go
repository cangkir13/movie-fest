package movievotecontroller

import (
	movieusecase "movie-festival-app/src/usecase/movie"
	movievoteusecase "movie-festival-app/src/usecase/movievote"
)

type movieVoteController struct {
	movieusecase     movieusecase.MovieUsecase
	movievoteusecase movievoteusecase.MovieVoteUsecase
}

func NewMovieVoteController(
	movieusecase movieusecase.MovieUsecase,
	movievoteusecase movievoteusecase.MovieVoteUsecase,
) *movieVoteController {
	return &movieVoteController{
		movieusecase:     movieusecase,
		movievoteusecase: movievoteusecase,
	}
}
