package main

import (
	"os"

	"movie-festival-app/config"
	"movie-festival-app/docs"
	"movie-festival-app/src/rest"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var db *gorm.DB = config.SetupDatabaseConnection()

// @title   Service message docs
// @version 1.0.0

// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @host                       localhost
// @BasePath                   /api
func main() {
	defer config.CloseDatabaseConnection(db)

	docs.SwaggerInfo.Host = os.Getenv("BASE_URL")
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := rest.Routes(db)

	swaggerGroup := r.Group("/api/swagger", gin.BasicAuth(gin.Accounts{
		"root": "root",
	}))

	swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	r.Run()
}
