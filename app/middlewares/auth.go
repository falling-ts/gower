package middlewares

import (
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

func Auth(args ...any) services.Handler {
	tokenKey := "token"
	authKey := "Authorization"
	fn := func(id string) (*models.Auth, error) {
		user := new(models.User)
		result := db.First(user, id)
		if result.Error != nil {
			return nil, trans.DBError(result.Error)
		}

		return &models.Auth{User: *user}, nil
	}

	if len(args) > 0 {
		if key, ok := args[0].(string); ok {
			tokenKey = key
		}
	}
	if len(args) > 1 {
		if key, ok := args[1].(string); ok {
			authKey = key
		}
	}

	return func(c *gin.Context) {
		c.Set("token-key", tokenKey)

		token, _ := cookie.Get(c, tokenKey)
		if token == "" {
			token = c.GetHeader(authKey)
		}
		if token == "" {
			exc.Unauthorized("未登录").Handle(c)
			return
		}

		userId, newToken, err := auth.Check(token, c.RemoteIP())
		if err != nil {
			exc.Unauthorized(err).Handle(c)
			return
		}
		if newToken != "" {
			c.Set(tokenKey, newToken)
		}

		if len(args) > 2 {
			if f, ok := args[2].(func(id string) (*models.Auth, error)); ok {
				fn = f
			}
		}

		model, err := fn(userId)
		if err != nil {
			exc.Unauthorized(err).Handle(c)
			return
		}

		c.Set("Auth", model)
		c.Next()
	}
}
