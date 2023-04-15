package movievoteusecase

import (
	"context"
	"fmt"

	"movie-festival-app/src/models"
	movievoterepository "movie-festival-app/src/repositories/movievote"
)

type UserVoteProps struct {
	IDMovie int
	IDUser  int
	Vote    int
}

func (repo *movieVoteUsecase) UserVote(ctx context.Context, props UserVoteProps) error {
	findMovie, err := repo.movierepository.GetOne(ctx, props.IDMovie)
	if err != nil {
		return err
	}
	if findMovie == nil {
		return fmt.Errorf("movie not found")
	}

	findVoteuser, err := repo.movievotereposiotry.GetOne(ctx, movievoterepository.GetOneProps{
		IDMovie: &props.IDMovie,
		IDUser:  &props.IDUser,
	})
	if err != nil {
		return err
	}

	if findVoteuser != nil {
		return fmt.Errorf("cannot double vote whit this movie")
	}

	voteErr := repo.movievotereposiotry.Create(ctx, models.MovieVote{
		IDMovie:   props.IDMovie,
		IDUser:    props.IDUser,
		Vote:      props.Vote,
		CreatedBy: props.IDUser,
	})

	if voteErr != nil {
		return voteErr
	}

	return nil
}

type UserUnVoteProps struct {
	IDMovie int
	IDUser  int
	Vote    int
}

func (repo *movieVoteUsecase) UserUnVote(ctx context.Context, props UserUnVoteProps) error {
	findMovie, err := repo.movierepository.GetOne(ctx, props.IDMovie)
	if err != nil {
		return err
	}
	if findMovie == nil {
		return fmt.Errorf("movie not found")
	}

	findVoteuser, err := repo.movievotereposiotry.GetOne(ctx, movievoterepository.GetOneProps{
		IDMovie: &props.IDMovie,
		IDUser:  &props.IDUser,
	})
	if err != nil {
		return err
	}

	if findVoteuser == nil {
		return fmt.Errorf("vote first")
	}

	voteErr := repo.movievotereposiotry.Update(ctx, movievoterepository.UpdateProps{
		ID:     findVoteuser.ID,
		Fields: []string{"vote"},
		Data: models.MovieVote{
			Vote: props.Vote,
		},
	})

	if voteErr != nil {
		return voteErr
	}

	return nil
}
