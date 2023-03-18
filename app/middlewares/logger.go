package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gower/services"
	"gower/utils/slice"
)

func Logger() services.Handler {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		req := fmt.Sprintf("%s %s", method, path)
		skips := config.Log.SkipPaths
		paths := config.Log.Paths

		if !slice.Strings(skips).HasPrefix(req) || slice.Strings(paths).HasPrefix(req) {
			logger.Info("Request Info",
				zap.String("describe", req),
				zap.Int("http_status", c.Writer.Status()),
				zap.String("ip", c.ClientIP()))
			c.Next()

			ResBody, _ := c.Get("body-logger")

			errs := make([]error, 0)
			for _, err := range c.Errors {
				errs = append(errs, err)
			}
			logger.Info("Response Info",
				zap.Reflect("body", ResBody),
				zap.Errors("exception", errs))
		}

		c.Next()
	}
}
