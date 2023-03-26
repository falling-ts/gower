package middlewares

import (
	"github.com/falling-ts/gower/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() services.Handler {
	return func(c *gin.Context) {
		fn := cors.New(cors.Config{
			AllowOrigins:     config.Cors.AllowOrigins,
			AllowMethods:     config.Cors.AllowMethods,
			AllowHeaders:     config.Cors.AllowHeaders,
			ExposeHeaders:    config.Cors.ExposeHeaders,
			AllowCredentials: true,
			MaxAge:           config.Cors.MaxAge,
		})

		fn(c)
	}
}
