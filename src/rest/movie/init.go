package moviecontroller

import movieusecase "movie-festival-app/src/usecase/movie"

type movieController struct {
	movieusecase movieusecase.MovieUsecase
}

func NewMovieController(movieusecase movieusecase.MovieUsecase) *movieController {
	return &movieController{movieusecase: movieusecase}
}
