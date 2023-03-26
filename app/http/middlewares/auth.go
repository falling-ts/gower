package middlewares

import (
	"github.com/falling-ts/gower/app/middlewares"
	"github.com/falling-ts/gower/services"
)

func Auth() services.Handler {
	return middlewares.Auth()
}
