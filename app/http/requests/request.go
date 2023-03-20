package requests

import (
	"gower/app"
	"gower/app/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	excp  = app.Exception()
	valid = app.Validator()
)

type request struct {
	*gin.Context
}

// Validate 执行验证
func (r *request) Validate(ctx *gin.Context, req http.Request) error {
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
func (r *request) SetContext(c *gin.Context) {
	r.Context = c
}
