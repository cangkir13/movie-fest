package userusecase

import (
	"context"
	"fmt"

	"movie-festival-app/src/helper"
	"movie-festival-app/src/models"
	userrepository "movie-festival-app/src/repositories/user"
)

type UserRegisterProps struct {
	Username string
	Password string
}

func (repo *userUsecase) UserRegister(ctx context.Context, props UserRegisterProps) error {
	user, err := repo.userrepository.GetOne(ctx, props.Username)
	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("account has been exist")
	}

	hash, _ := helper.HashPassword(props.Password)

	propsuser := userrepository.CreateProps{
		Data: models.User{
			Username: props.Username,
			Password: hash,
		},
	}

	newuser := repo.userrepository.Create(ctx, propsuser)

	if newuser != nil {
		return newuser
	}

	return nil
}

type UserLoginProps struct {
	Username string
	Password string
}

func (repo *userUsecase) UserLogin(ctx context.Context, props UserLoginProps) (*string, error) {
	account, err := repo.userrepository.GetOne(ctx, props.Username)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, fmt.Errorf("account not found")
	}

	if !account.IsAllow {
		return nil, fmt.Errorf("account not allow to access")
	}

	// decrypt password
	checkPass := helper.DecryptPassword(props.Password, account.Password)
	if !checkPass {
		return nil, fmt.Errorf("wrong password")
	}

	// service jwt
	token, err := helper.SignIn(helper.JWTClaim{
		UserID:   account.ID,
		Username: account.Username,
		RoleID:   account.RoleID,
	})
	if err != nil {
		return nil, err
	}

	return &token, nil
}
