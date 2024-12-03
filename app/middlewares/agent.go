package middlewares

import (
	"gitee.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
	"strings"
)

func Agent() services.Handler {
	return func(c *gin.Context) {
		if isMobile(c.GetHeader("User-Agent")) {
			c.Set("isMobile", true)
		}
	}
}

func isMobile(userAgent string) bool {
	mobileKeywords := []string{
		"Android", "webOS", "iPhone", "iPad", "iPod", "BlackBerry", "Windows Phone",
	}

	for _, keyword := range mobileKeywords {
		if strings.Contains(userAgent, keyword) {
			return true
		}
	}
	return false
}
