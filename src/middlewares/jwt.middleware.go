package middlewares

import (
	"movie-festival-app/src/helper"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Disposition, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func JWTAdministrator() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			result := helper.PublicResponse{Status: false, Message: "request does not contain an access token"}
			c.AbortWithStatusJSON(401, result)
			return
		}

		// decode jwt
		tokenString := authHeader[len(BEARER_SCHEMA):]
		data, err := helper.DecodeJWTAdmin(tokenString)
		if err != nil {
			result := helper.PublicResponse{Status: false, Message: err.Error()}
			c.AbortWithStatusJSON(401, result)
			return
		}

		errCompare := helper.ValidateToken(tokenString)
		if errCompare != nil {
			result := helper.PublicResponse{Status: false, Message: errCompare.Error()}
			c.AbortWithStatusJSON(401, result)
			return
		}

		if data.RoleID != 1 {
			result := helper.PublicResponse{Status: false, Message: "Acess denied id role"}
			c.AbortWithStatusJSON(401, result)
			return
		}

		c.Set("id", data.UserID)
		c.Set("username", data.Username)
		c.Set("role_id", data.RoleID)
		c.Next()
	}
}

func JWTClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			result := helper.PublicResponse{Status: false, Message: "request does not contain an access token"}
			c.AbortWithStatusJSON(401, result)
			return
		}

		// decode jwt
		tokenString := authHeader[len(BEARER_SCHEMA):]
		data, err := helper.DecodeJWTAdmin(tokenString)
		if err != nil {
			result := helper.PublicResponse{Status: false, Message: err.Error()}
			c.AbortWithStatusJSON(401, result)
			return
		}

		errCompare := helper.ValidateToken(tokenString)
		if errCompare != nil {
			result := helper.PublicResponse{Status: false, Message: errCompare.Error()}
			c.AbortWithStatusJSON(401, result)
			return
		}

		if data.RoleID != 2 {
			result := helper.PublicResponse{Status: false, Message: "Acess denied id role"}
			c.AbortWithStatusJSON(401, result)
			return
		}

		c.Set("id", data.UserID)
		c.Set("username", data.Username)
		c.Set("role_id", data.RoleID)
		c.Next()
	}
}

func UserGetId(c *gin.Context) int {
	return c.GetInt("id")
}

func UserGetRole(c *gin.Context) int {
	return c.GetInt("role_id")
}
