package middlewares

import (
	"gitee.com/falling-ts/gower/app/middlewares"
	"gitee.com/falling-ts/gower/services"
)

func Auth() services.Handler {
	return middlewares.Auth()
}
