package usercontroller

import (
	"net/http"

	"movie-festival-app/src/helper"
	userusecase "movie-festival-app/src/usecase/user"

	"github.com/gin-gonic/gin"
)

type RegisterRequestJson struct {
	Username string `json:"username" binding:"required" example:"edp"`
	Password string `json:"password" binding:"required" validate:"min:6" example:"edp"`
}

// Register user
// @Summary     register
// @Description register
// @Tags        Account
// @Accept      json
// @Produce     json
// @Param       raw body     RegisterRequestJson true "body"
// @Success     200 {object} helper.PublicResponse
// @Failure     400 {object} helper.PublicResponse
// @Failure     404 {object} helper.PublicResponse
// @Failure     401 {object} helper.PublicResponse
// @Router      /register [POST]
func (usecase *userController) Register(c *gin.Context) {
	var request RegisterRequestJson
	if err := c.BindJSON(&request); err != nil {
		result := helper.PublicResponse{Status: false, Message: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	usercaseErr := usecase.userusecase.UserRegister(c, userusecase.UserRegisterProps(request))

	if usercaseErr != nil {
		result := helper.PublicResponse{Status: false, Message: usercaseErr.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result := helper.PublicResponse{Status: true, Message: "Data has been created"}
	c.AbortWithStatusJSON(http.StatusCreated, result)
}

type LoginRequestJson struct {
	Username string `json:"username" binding:"required" example:"edp"`
	Password string `json:"password" binding:"required" validate:"min:6" example:"edp"`
}

// login user
// @Summary     login
// @Description login service account
// @Tags        Account
// @Accept      json
// @Produce     json
// @Param       raw body     LoginRequestJson true "body"
// @Success     200 {object} helper.PublicResponse
// @Failure     400 {object} helper.PublicResponse
// @Failure     404 {object} helper.PublicResponse
// @Failure     401 {object} helper.PublicResponse
// @Router      /login [POST]
func (usecase *userController) Login(c *gin.Context) {
	var request LoginRequestJson
	if err := c.BindJSON(&request); err != nil {
		result := helper.PublicResponse{Status: false, Message: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	// account usecase login
	token, err := usecase.userusecase.UserLogin(c, userusecase.UserLoginProps(request))
	if err != nil {
		result := helper.PublicResponse{Status: false, Message: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result := helper.PublicResponse{Status: true, Message: "ok", Result: token}
	c.JSON(200, result)
}
