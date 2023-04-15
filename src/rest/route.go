package rest

import (
	"movie-festival-app/src/middlewares"
	movierepository "movie-festival-app/src/repositories/movie"
	movievoterepository "movie-festival-app/src/repositories/movievote"
	userrepository "movie-festival-app/src/repositories/user"
	moviecontroller "movie-festival-app/src/rest/movie"
	movievotecontroller "movie-festival-app/src/rest/movievote"
	usercontroller "movie-festival-app/src/rest/user"
	movieusecase "movie-festival-app/src/usecase/movie"
	movievoteusecase "movie-festival-app/src/usecase/movievote"
	userusecase "movie-festival-app/src/usecase/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	var (
		// load repositories
		userrepository      = userrepository.NewUserRepositories(db)
		movierepository     = movierepository.NewMovieRepositories(db)
		movievoterepository = movievoterepository.NewMovieVoteRepositories(db)

		// load usecase
		userusecase      = userusecase.NewUserUsecase(userrepository)
		movieusecase     = movieusecase.NewMovieUsecase(movierepository)
		movievoteusecase = movievoteusecase.NewMovieVoteUsecase(movierepository, movievoterepository)

		// load controller
		usercontroller      = usercontroller.NewUserController(userusecase)
		moviecontroller     = moviecontroller.NewMovieController(movieusecase)
		movievotecontroller = movievotecontroller.NewMovieVoteController(movieusecase, movievoteusecase)
	)

	routes := r.Group("/api")

	routes.POST("/register", usercontroller.Register)
	routes.POST("/login", usercontroller.Login)

	movieRoutes := routes.Group("movies")
	{
		movieRoutes.GET("", moviecontroller.GetListMovie)
		movieRoutes.POST("", middlewares.JWTAdministrator(), moviecontroller.Store)
		movieRoutes.PUT(":id", middlewares.JWTAdministrator(), moviecontroller.Update)

		movieRoutes.POST("/vote", middlewares.JWTClient(), movievotecontroller.Vote)
		movieRoutes.POST("/unvote", middlewares.JWTClient(), movievotecontroller.Unvote)
	}

	return r
}
