package middleware

import (
	"fmt"
	"git.lifewood.dev/services/skeleton/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
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

func LoggerCustome() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}
