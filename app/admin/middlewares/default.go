package middlewares

import (
	"github.com/falling-ts/gower/app/middlewares"
	"github.com/falling-ts/gower/services"
)

func Default() services.Handler {
	return middlewares.Default("admin-auth", "Admin-Authorization")
}
