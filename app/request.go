package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gower/iface"
)

var (
	excp  = Exception()
	valid = Validator()
)

type Request struct {
	*gin.Context
}

// Validate 执行验证
func (r *Request) Validate(ctx *gin.Context, req iface.Request) error {
	r.Context = ctx

	if err := ctx.ShouldBind(req); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return excp.BadRequest("验证器错误")
		}

		errs := err.(validator.ValidationErrors)
		return excp.UnprocessableEntity(errs, errs[0].Translate(valid.GetTrans()), valid.Translate(errs))
	}

	return nil
}

// SetContext 设置 gin 的请求体
func (r *Request) SetContext(c *gin.Context) {
	r.Context = c
}
