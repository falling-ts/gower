package requests

import "github.com/gin-gonic/gin"

type Request interface {
	Validate(...any) bool
}

type request struct {
	*gin.Context
}

func (r *request) Validate(ctx *gin.Context, req Request) bool {
	r.Context = ctx

	if err := ctx.ShouldBind(req); err != nil {
		return false
	}

	return true
}
