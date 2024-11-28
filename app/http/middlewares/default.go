package middlewares

import (
	"gitee.com/falling-ts/gower/app/middlewares"
	"gitee.com/falling-ts/gower/services"
)

func Default() services.Handler {
	return middlewares.Default()
}
