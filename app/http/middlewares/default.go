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

		if token == "" {
			c.Next()
			return
		}

		userId, newToken, err := auth.Check(token, c.RemoteIP())
		if err != nil {
			c.Next()
			return
		}

		user := new(models.User)
		result := db.First(user, userId)
		if result.Error != nil {
			c.Next()
			return
		}

		if newToken != "" {
			c.Set("token", newToken)
		}
		c.Set("Auth", &models.Auth{User: *user})
		
		c.Next()
	}
}
