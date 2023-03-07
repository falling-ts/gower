package requests

import (
	"github.com/gin-gonic/gin"
	"gower/app"
)

var excp = app.Exception()

type Request interface {
	Validate(...any) bool
}

type request struct {
	*gin.Context
}

func (r *request) Validate(ctx *gin.Context, req Request) error {
	r.Context = ctx

	if err := ctx.ShouldBind(req); err != nil {
		return excp.BadRequest(err)
	}

	return nil
}
