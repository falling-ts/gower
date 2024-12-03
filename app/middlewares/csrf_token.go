package middlewares

import (
	"gitee.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CsrfToken() services.Handler {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			c.Next()
			return
		}

		csrfToken, _ := cookie.Get(c, "csrfToken")
		if csrfToken == "" {
			c.Next()
			return
		}

		c.Set("csrfToken", csrfToken)
		c.Next()
	}
}
