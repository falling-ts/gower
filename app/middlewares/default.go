package middlewares

import (
	"github.com/gin-gonic/gin"
	"gower/app/models"
	"gower/services"
)

func Default(args ...string) services.Handler {
	cookieKey := "token"
	headerKey := "Authorization"
	if len(args) > 0 {
		cookieKey = args[0]
	}
	if len(args) > 1 {
		headerKey = args[1]
	}

	return func(c *gin.Context) {
		token, _ := c.Cookie(cookieKey)
		if token == "" {
			token = c.GetHeader(headerKey)
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
