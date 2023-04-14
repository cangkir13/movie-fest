package userusecase

import (
	"context"

	userrepository "movie-festival-app/src/repositories/user"
)

type userUsecase struct {
	userrepository userrepository.UserRepositories
}

func NewUserUsecase(userrepository userrepository.UserRepositories) *userUsecase {
	return &userUsecase{userrepository: userrepository}
}

type UserUsecase interface {
	UserRegister(ctx context.Context, props UserRegisterProps) error
	UserLogin(ctx context.Context, props UserLoginProps) (*string, error)
}
