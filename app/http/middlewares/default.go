package middlewares

import (
	"github.com/gin-gonic/gin"
	"gower/app/models"
	"gower/services"
)

func Default() services.Handler {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		if token == "" {
			token = c.GetHeader("Authorization")
		}

		if token != "" {
			userId, newToken, _ := auth.Check(token, c.RemoteIP())
			if newToken != "" {
				c.Set("token", newToken)
			}

			user := new(models.User)
			result := db.First(user, userId)
			if result.Error == nil {
				c.Set("Auth", &models.Auth{User: *user})
			}
		}

		c.Next()
	}
}
