package usercontroller

import userusecase "movie-festival-app/src/usecase/user"

type userController struct {
	userusecase userusecase.UserUsecase
}

func NewUserController(userusecase userusecase.UserUsecase) *userController {
	return &userController{userusecase: userusecase}
}
