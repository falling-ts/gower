package middlewares

import (
	"github.com/falling-ts/gower/app/middlewares"
	"github.com/falling-ts/gower/services"
)

var _ = Default()

func Default() services.Handler {
	return middlewares.Default("admin-auth", "Authorization")
}
