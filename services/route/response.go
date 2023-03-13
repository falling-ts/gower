package route

import (
	"net/http"

	"gower/services"

	"github.com/gin-gonic/gin"
)

// Response 响应结构体
type Response struct {
	HttpStatus int
	config     gin.Negotiate
}

// Accepts 接受的响应数据类型
var Accepts = []string{
	gin.MIMEJSON,
	gin.MIMEHTML,
	gin.MIMEXML,
	gin.MIMEYAML,
	gin.MIMETOML,
	gin.MIMEPlain,
}

// Response 设置响应结构体内容
func (r *Route) Response(data services.ResponseData, args ...any) services.Response {
	res := new(Response)
	res.HttpStatus = http.StatusOK
	res.config.Offered = Accepts

	res.DecideType(data, "SUCCESS")
	if len(args) > 0 {
		res.DecideType(data, args[0])
	}
	if len(args) > 1 {
		res.DecideType(data, args[1])
	}
	if len(args) > 2 {
		res.DecideType(data, args[2])
	}
	if len(args) > 3 {
		res.DecideType(data, args[3])
	}
	if len(args) > 4 {
		res.DecideType(data, args[4])
	}
	if len(args) > 5 {
		res.DecideType(data, args[5])
	}

	return res
}

// DecideType 判定类型
func (r *Response) DecideType(data services.ResponseData, arg any) {
	switch arg.(type) {
	case int:
		code := arg.(int)
		if code >= http.StatusOK && code < http.StatusMultipleChoices {
			r.HttpStatus = code
		} else {
			data.Set(arg)
		}
	case string:
		data.Set(arg)
		r.config.HTMLName = arg.(string)
	default:
		data.Set(arg)
		r.config.HTMLData = arg
	}

	r.config.Data = data
}

func handleResponse(r services.Response, c *gin.Context) {
	if response, ok := r.(*Response); ok {
		c.Negotiate(response.HttpStatus, response.config)
	}

	handleException(exceptions.New(http.StatusBadRequest, "响应体错误."), c)
}

func handleException(e services.Exceptions, c *gin.Context) {
	err := e.Get("RawErr")
	if err != nil {
		_ = c.Error(err.(error))
	} else {
		_ = c.Error(e)
	}

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
			c.Redirect(http.StatusFound, c.Request.URL.String())
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
