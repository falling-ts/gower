package requests

import (
	"gower/app"

	"github.com/gin-gonic/gin"
)

var excp = app.Exceptions()

type Request interface {
	Validate(ctx *gin.Context, req Request) error
	SetContext(c *gin.Context)
}

type request struct {
	*gin.Context
}

// Validate 执行验证
func (r *request) Validate(ctx *gin.Context, req Request) error {
	r.Context = ctx

	return excp.BadRequest("test error")
	if err := ctx.ShouldBind(req); err != nil {
		return excp.BadRequest(err)
	}

	return nil
}

// SetContext 设置 gin 的请求体
func (r *request) SetContext(c *gin.Context) {
	r.Context = c
}
