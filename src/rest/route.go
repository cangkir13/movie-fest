package rest

import (
	movierepository "movie-festival-app/src/repositories/movie"
	userrepository "movie-festival-app/src/repositories/user"
	moviecontroller "movie-festival-app/src/rest/movie"
	usercontroller "movie-festival-app/src/rest/user"
	movieusecase "movie-festival-app/src/usecase/movie"
	userusecase "movie-festival-app/src/usecase/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	var (
		// load repositories
		userrepository  = userrepository.NewUserRepositories(db)
		movierepository = movierepository.NewMovieRepositories(db)

		// load usecase
		userusecase  = userusecase.NewUserUsecase(userrepository)
		movieusecase = movieusecase.NewMovieUsecase(movierepository)

		// load controller
		usercontroller  = usercontroller.NewUserController(userusecase)
		moviecontroller = moviecontroller.NewMovieController(movieusecase)
	)

	routes := r.Group("/api")

	routes.POST("/register", usercontroller.Register)
	routes.POST("/login", usercontroller.Login)

	movieRoutes := routes.Group("movies")
	{
		movieRoutes.GET("", moviecontroller.GetListMovie)
		movieRoutes.POST("", moviecontroller.Store)
		movieRoutes.PUT(":id", moviecontroller.Update)
	}

	return r
}
