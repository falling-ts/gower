package middlewares

import (
	"github.com/falling-ts/gower/app/models"
	"github.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

func Default(args ...string) services.Handler {
	cookieKey := "auth"
	headerKey := "Authorization"
	if len(args) > 0 {
		cookieKey = args[0]
	}
	if len(args) > 1 {
		headerKey = args[1]
	}

	return func(c *gin.Context) {
		token, _ := cookie.Get(c, cookieKey)
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
			c.Set("auth", newToken)
		}
		c.Set("Auth", &models.Auth{User: *user})

		c.Next()
	}
}
