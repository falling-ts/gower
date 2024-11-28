package middlewares

import (
	"gitee.com/falling-ts/gower/app/middlewares"
	"gitee.com/falling-ts/gower/services"
)

var _ = Auth()

func Auth() services.Handler {
	return middlewares.Auth("api-token")
}
