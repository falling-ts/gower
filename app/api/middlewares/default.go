package middlewares

import (
	"gitee.com/falling-ts/gower/app/middlewares"
	"gitee.com/falling-ts/gower/services"
)

var _ = Default()

func Default() services.Handler {
	return middlewares.Default("api-token")
}
