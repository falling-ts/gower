package middlewares

import (
	"crypto/subtle"
	"github.com/gin-gonic/gin"
	"gower/services"
	"net/http"
)

func CsrfToken() services.Handler {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			c.Next()
			return
		}

		realToken, _ := cookie.Get(c, "csrf_token")
		if realToken == "" {
			c.Next()
			return
		}
		sendToken := c.PostForm("csrf_token")
		if sendToken == "" {
			sendToken = c.Query("csrf_token")
		}
		if sendToken == "" {
			excp.NotAcceptable("CSRF 校验失败").Handle(c)
			return
		}

		if subtle.ConstantTimeCompare([]byte(realToken), []byte(sendToken)) == 0 {
			excp.NotAcceptable("CSRF 校验失败").Handle(c)
			return
		}

		c.Next()
	}
}
