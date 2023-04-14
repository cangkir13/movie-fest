package rest

import (
	userrepository "movie-festival-app/src/repositories/user"
	usercontroller "movie-festival-app/src/rest/user"
	userusecase "movie-festival-app/src/usecase/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	var (
		// load repositories
		userrepository = userrepository.NewUserRepositories(db)

		// load usecase
		userusecase = userusecase.NewUserUsecase(userrepository)

		// load controller
		usercontroller = usercontroller.NewUserController(userusecase)
	)

	routes := r.Group("/api")

	routes.POST("/register", usercontroller.Register)
	routes.POST("/login", usercontroller.Login)

	return r
}
