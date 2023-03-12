package exception

import (
	"gower/services/cache"
	"gower/services/config"
	"net/http"

	"gower/services"

	"github.com/gin-gonic/gin"
)

// Content 异常内容
type Content interface {
	error
	SetException(exception *Struct)
	SetMsg(msg string)
	SetData(data any)
	Throw(code uint, args ...any) Content
	HandleBy(any)
}

// Struct 异常主结构体
type Struct struct {
	Content
	RawErr  error
	handled bool
}

var (
	Entity  = new(Struct)
	Accepts = []string{
		gin.MIMEJSON,
		gin.MIMEHTML,
		gin.MIMEXML,
		gin.MIMEYAML,
		gin.MIMETOML,
		gin.MIMEPlain,
	}
)

// New 创建新异常服务
func New() *Struct {
	return new(Struct)
}

// Init 服务初始化
func (e *Struct) Init(args ...any) services.Service {
	if len(args) == 0 {
		panic("初始化参数不存在")
	}

	content, ok := args[0].(Content)
	if !ok {
		panic("异常服务初始化失败")
	}
	e.Content = content

	e.Content.SetException(e)
	return e
}

// Build 构建每个请求的异常
func (e *Struct) Build(code uint, args ...any) Content {
	e.Content.SetMsg("未知异常")
	argsNum := len(args)

	if argsNum > 0 {
		decideType(args[0], e)
	}
	if argsNum > 1 {
		decideType(args[1], e)
	}
	if argsNum > 2 {
		decideType(args[2], e)
	}
	if argsNum > 3 {
		decideType(args[3], e)
	}
	if argsNum > 4 {
		decideType(args[4], e)
	}
	if argsNum > 5 {
		decideType(args[5], e)
	}

	return e.Content
}

// Exception 获取异常实体
func (e *Struct) Exception() Content {
	return e.Content
}

// HandleBy 处理异常
func (e *Struct) HandleBy(arg any) {
	if e.handled {
		return
	}

	c, ok := arg.(*gin.Context)
	if !ok {
		panic("处理异常参数错误")
	}

	_ = c.Error(e.RawErr)
	switch c.NegotiateFormat(Accepts...) {
	case gin.MIMEJSON:
		c.JSON(http.StatusOK, e.Content)
	case gin.MIMEHTML:
		key := getKey()
		cache.Entity.SetDefault(key, e)
		c.SetCookie(
			"err-key",
			key,
			300,
			"/",
			config.Entity.Get("app.domain", "localhost").(string),
			false,
			true)

		referer := c.Request.Referer()
		if referer == "" {
			c.Redirect(http.StatusMovedPermanently, c.Request.URL.String())
		} else {
			c.Redirect(http.StatusFound, referer)
		}
	case gin.MIMEXML:
		c.XML(http.StatusOK, e.Content)
	case gin.MIMEYAML:
		c.YAML(http.StatusOK, e.Content)
	case gin.MIMETOML:
		c.TOML(http.StatusOK, e.Content)
	default:
		c.String(http.StatusOK, e.Content.Error())
	}

	e.handled = true
}
