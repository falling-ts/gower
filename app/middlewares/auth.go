package middlewares

import (
	"github.com/falling-ts/gower/app/models"
	"github.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

func Auth(args ...string) services.Handler {
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
			excp.Unauthorized("未登录").Handle(c)
			return
		}

		userId, newToken, err := auth.Check(token, c.RemoteIP())
		if err != nil {
			excp.Unauthorized(err).Handle(c)
			return
		}
		if newToken != "" {
			c.Set("auth", newToken)
		}

		user := new(models.User)
		result := db.First(user, userId)
		if result.Error != nil {
			excp.Unauthorized(trans.DBError(result.Error)).Handle(c)
			return
		}

		c.Set("Auth", &models.Auth{User: *user})
		c.Next()
	}
}
