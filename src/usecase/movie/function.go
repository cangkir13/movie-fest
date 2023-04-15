package movieusecase

import (
	"context"
	"fmt"

	"movie-festival-app/src/models"
	movierepository "movie-festival-app/src/repositories/movie"
)

type GetListProps struct {
	Page        int
	Limit       int
	Title       *string
	Description *string
	Artists     *string
	Genres      *string
}

func (repo *movieUsecase) GetList(ctx context.Context, props GetListProps) (records []*models.Movie, err error) {
	records, err = repo.movierepository.List(ctx, movierepository.ListProps(props))

	if err != nil {
		return
	}

	return
}

func (repo *movieUsecase) CreateNew(ctx context.Context, props models.Movie) error {
	createdata := repo.movierepository.Create(ctx, props)
	if createdata != nil {
		return createdata
	}

	return nil
}

type UpdateMovieProps struct {
	ID     int
	Fields []string
	Data   models.Movie
}

func (repo *movieUsecase) UpdateMovie(ctx context.Context, props UpdateMovieProps) error {
	// check movie id
	findMovie, err := repo.movierepository.GetOne(ctx, props.ID)
	if err != nil {
		return err
	}

	if findMovie == nil {
		return fmt.Errorf("movie not found")
	}

	// custom inputs data
	var (
		fields     []string
		updateData models.Movie
	)

	if props.Data.Title != "" {
		updateData.Title = props.Data.Title
		fields = append(fields, "title")
	}

	if props.Data.Description != "" {
		updateData.Description = props.Data.Description
		fields = append(fields, "description")
	}

	if props.Data.Genres != "" {
		updateData.Genres = props.Data.Genres
		fields = append(fields, "genres")
	}

	if props.Data.Url != "" {
		updateData.Url = props.Data.Url
		fields = append(fields, "url")
	}

	if props.Data.UpdatedBy != 0 {
		updateData.UpdatedBy = props.Data.UpdatedBy
		fields = append(fields, "updated_by")
	}

	if props.Data.Artists != "" {
		updateData.Artists = props.Data.Artists
		fields = append(fields, "artists")
	}

	propsUpdate := movierepository.UpdateProps{
		ID:     props.ID,
		Fields: fields,
		Data:   updateData,
	}

	err = repo.movierepository.Update(ctx, propsUpdate)
	if err != nil {
		return err
	}

	return nil
}
