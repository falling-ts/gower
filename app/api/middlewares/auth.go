package middlewares

import (
	"github.com/falling-ts/gower/app/middlewares"
	"github.com/falling-ts/gower/services"
)

var _ = Auth()

func Auth() services.Handler {
	return middlewares.Auth("api-auth", "Authorization")
}
