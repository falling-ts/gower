package middlewares

import (
	"gower/app/middlewares"
	"gower/services"
)

func Auth() services.Handler {
	return middlewares.Auth("admin-token", "Admin-Authorization")
}
