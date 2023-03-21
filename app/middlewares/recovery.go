package middlewares

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"gower/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery 恐慌捕获
func Recovery() services.Handler {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}

				logger.Panic(err.(string),
					zap.String("headers", strings.Join(headers, "|")),
					zap.Stack("stack"))

				excp.New(http.StatusInternalServerError, err).Handle(c)
			}
		}()
		c.Next()
	}
}
