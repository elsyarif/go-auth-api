package middleware

import (
	"github.com/elsyarif/go-auth-api/pkg/helper/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		log.Info("origin", logrus.Fields{"origin": origin})

		if origin != "" && isAllowedOrigin(origin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
		}

		c.Next()
	}
}

func isAllowedOrigin(origin string) bool {
	allowedOrigin := []string{"http://localhost:3000", "http://127.0.0.1:3000/"}

	for _, o := range allowedOrigin {
		if o == origin {
			return true
		}
	}

	return false
}