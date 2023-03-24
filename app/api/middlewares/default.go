package middlewares

import (
	"gower/app/middlewares"
	"gower/services"
)

func Default() services.Handler {
	return middlewares.Default("api-auth", "Api-Authorization")
}
