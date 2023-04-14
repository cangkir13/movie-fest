package moviecontroller

import (
	"net/http"
	"strconv"

	"movie-festival-app/src/helper"
	"movie-festival-app/src/models"
	movieusecase "movie-festival-app/src/usecase/movie"

	"github.com/gin-gonic/gin"
)

type GetListMovieRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

// Movie
// @Summary     Movie
// @Description Movie service
// @Tags        Movies
// @Accept      json
// @Produce     json
// @Param       page  query    int false "page"
// @Param       limit query    int false "limit"
// @Success     200   {object} helper.PublicResponse
// @Failure     400   {object} helper.PublicResponse
// @Failure     404   {object} helper.PublicResponse
// @Failure     401   {object} helper.PublicResponse
// @Router      /movies [GET]
func (usecase *movieController) GetListMovie(c *gin.Context) {
	var request GetListMovieRequest
	c.ShouldBindQuery(&request)

	records, err := usecase.movieusecase.GetList(c, movieusecase.GetListProps(request))
	if err != nil {
		result := helper.PublicResponse{Status: false, Message: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result := helper.PublicResponse{Status: true, Message: "Data movies", Result: records}
	c.AbortWithStatusJSON(http.StatusOK, result)
}

type StoreJson struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Genres      string `json:"genres"`
	Url         string `json:"url"`
}

// Movie
// @Summary     Movie
// @Description Movie service
// @Tags        Movies
// @Accept      json
// @Produce     json
// @Param       raw body     StoreJson true "body"
// @Success     200 {object} helper.PublicResponse
// @Failure     400 {object} helper.PublicResponse
// @Failure     404 {object} helper.PublicResponse
// @Failure     401 {object} helper.PublicResponse
// @Router      /movies [POST]
func (usecase *movieController) Store(c *gin.Context) {
	var request StoreJson
	if err := c.BindJSON(&request); err != nil {
		result := helper.PublicResponse{Status: false, Message: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	errinsert := usecase.movieusecase.CreateNew(c, models.Movie{
		Title:       request.Title,
		Description: request.Description,
		Duration:    request.Duration,
		Genres:      request.Genres,
		Url:         request.Url,
	})

	if errinsert != nil {
		result := helper.PublicResponse{Status: false, Message: errinsert.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result := helper.PublicResponse{Status: true, Message: "Data has been created"}
	c.AbortWithStatusJSON(http.StatusCreated, result)
}

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Genres      string `json:"genres"`
	Url         string `json:"url"`
}

// Movie
// @Summary     Movie
// @Description Movie service
// @Tags        Movies
// @Accept      json
// @Produce     json
// @Param       id  path     int       true "ID"
// @Param       raw body     StoreJson true "body"
// @Success     200 {object} helper.PublicResponse
// @Failure     400 {object} helper.PublicResponse
// @Failure     404 {object} helper.PublicResponse
// @Failure     401 {object} helper.PublicResponse
// @Router      /movies/{id} [PUT]
func (usecase *movieController) Update(c *gin.Context) {
	var (
		request UpdateRequest
		id      = c.Param("id")
	)
	if err := c.ShouldBindJSON(&request); err != nil {
		result := helper.PublicResponse{Status: false, Message: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	// parsing to int
	intVar, _ := strconv.ParseInt(id, 0, 62)

	err := usecase.movieusecase.UpdateMovie(c, movieusecase.UpdateMovieProps{
		ID: int(intVar),
		Fields: []string{
			"title",
			"description",
			"duration",
			"genres",
			"url",
			"updated_by",
		},
		Data: models.Movie{
			Title:       request.Title,
			Description: request.Description,
			Duration:    request.Duration,
			Genres:      request.Genres,
			Url:         request.Url,
			UpdatedBy:   1,
		},
	})
	if err != nil {
		result := helper.PublicResponse{Status: false, Message: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, result)
		return
	}

	result := helper.PublicResponse{Status: true, Message: "Data has been updated"}
	c.AbortWithStatusJSON(http.StatusOK, result)
}
