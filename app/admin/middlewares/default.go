package middlewares

import (
	"gower/app/middlewares"
	"gower/services"
)

func Default() services.Handler {
	return middlewares.Default("admin-auth", "Admin-Authorization")
}
