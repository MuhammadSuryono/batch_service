package middleware

import (
	"git.lifewood.dev/services/skeleton/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AuthorizeHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		token := os.Getenv("APP_TOKEN")
		TOKEN := BEARER_SCHEMA + " " + token
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.CommonResponse{
				IsSuccess: "1",
				Message:   "Token not found!",
				Data:      authHeader,
			})
			return
		}

		if authHeader != TOKEN {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.CommonResponse{
				IsSuccess: "1",
				Message:   "Token is not valid",
				Data:      authHeader,
			})
			return
		}

		c.Next()
	}
}
