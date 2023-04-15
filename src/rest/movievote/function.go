package movievotecontroller

import (
	"net/http"

	"movie-festival-app/src/helper"
	"movie-festival-app/src/middlewares"
	movievoteusecase "movie-festival-app/src/usecase/movievote"

	"github.com/gin-gonic/gin"
)

type VoteRequest struct {
	IDMovie int `json:"id_movie"`
}

// Movie
// @Summary     Vote movie
// @Description Vote Movie service
// @Tags        Movies-Vote
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Param       raw body     VoteRequest true "body"
// @Success     200 {object} helper.PublicResponse
// @Failure     400 {object} helper.PublicResponse
// @Failure     404 {object} helper.PublicResponse
// @Failure     401 {object} helper.PublicResponse
// @Router      /movies/vote [POST]
func (usecase *movieVoteController) Vote(c *gin.Context) {
	var request VoteRequest
	if err := c.BindJSON(&request); err != nil {
		result := helper.PublicResponse{Status: false, Message: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	payload := movievoteusecase.UserVoteProps{
		IDMovie: request.IDMovie,
		IDUser:  middlewares.UserGetId(c),
		Vote:    1,
	}

	errinsert := usecase.movievoteusecase.UserVote(c, payload)

	if errinsert != nil {
		result := helper.PublicResponse{Status: false, Message: errinsert.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result := helper.PublicResponse{Status: true, Message: "voted"}
	c.AbortWithStatusJSON(http.StatusCreated, result)
}

// Movie
// @Summary     Unvote Movie
// @Description Unvote Movie service
// @Tags        Movies-Vote
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Param       raw body     VoteRequest true "body"
// @Success     200 {object} helper.PublicResponse
// @Failure     400 {object} helper.PublicResponse
// @Failure     404 {object} helper.PublicResponse
// @Failure     401 {object} helper.PublicResponse
// @Router      /movies/unvote [POST]
func (usecase *movieVoteController) Unvote(c *gin.Context) {
	var request VoteRequest
	if err := c.BindJSON(&request); err != nil {
		result := helper.PublicResponse{Status: false, Message: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	payload := movievoteusecase.UserUnVoteProps{
		IDMovie: request.IDMovie,
		IDUser:  middlewares.UserGetId(c),
		Vote:    0,
	}

	errinsert := usecase.movievoteusecase.UserUnVote(c, payload)

	if errinsert != nil {
		result := helper.PublicResponse{Status: false, Message: errinsert.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result := helper.PublicResponse{Status: true, Message: "voted"}
	c.AbortWithStatusJSON(http.StatusCreated, result)
}
