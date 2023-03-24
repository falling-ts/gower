package middlewares

import (
	"github.com/gin-contrib/cors"
	"gower/services"
)

func Cors() services.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowMethods:     config.Cors.AllowMethods,
		AllowHeaders:     config.Cors.AllowHeaders,
		ExposeHeaders:    config.Cors.ExposeHeaders,
		AllowCredentials: true,
		MaxAge:           config.Cors.MaxAge,
	})
}
