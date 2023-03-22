package middlewares

import (
	"github.com/gin-gonic/gin"
	"gower/app/models"
	"gower/services"
)

func Auth() services.Handler {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		if token == "" {
			token = c.GetHeader("Authorization")
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
			c.Set("token", newToken)
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
