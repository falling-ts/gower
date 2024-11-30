package middlewares

import (
	"gitee.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

func Breadcrumbs() services.Handler {
	return func(c *gin.Context) {

		c.Next()
	}
}
