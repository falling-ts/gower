package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gower/services"
)

var Accepts = []string{
	gin.MIMEJSON,
	gin.MIMEHTML,
	gin.MIMEXML,
	gin.MIMEYAML,
	gin.MIMETOML,
	gin.MIMEPlain,
}

// HandleBy 处理异常
func handleException(e services.Exceptions, c *gin.Context) {
	_ = c.Error(e.Get("RawErr").(error))

	switch c.NegotiateFormat(Accepts...) {
	case gin.MIMEJSON:
		c.JSON(http.StatusOK, e)
	case gin.MIMEHTML:
		key := getKey()
		cache.SetDefault(key, e)
		c.SetCookie(
			"err-key",
			key,
			300,
			"/",
			configs.Get("app.domain", "localhost").(string),
			false,
			true)

		referer := c.Request.Referer()
		if referer == "" {
			c.Redirect(http.StatusMovedPermanently, c.Request.URL.String())
		} else {
			c.Redirect(http.StatusFound, referer)
		}
	case gin.MIMEXML:
		c.XML(http.StatusOK, e)
	case gin.MIMEYAML:
		c.YAML(http.StatusOK, e)
	case gin.MIMETOML:
		c.TOML(http.StatusOK, e)
	default:
		c.String(http.StatusOK, e.Error())
	}

}
