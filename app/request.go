package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	excp  = Exception()
	valid = Validator()
)

// RequestIFace 通用请求接口
type RequestIFace interface {
	Validate(c *gin.Context, req RequestIFace) error
}

type Request struct {
	*gin.Context `json:"-" xml:"-" form:"-" query:"-" protobuf:"-" msgpack:"-" yaml:"-" uri:"-" header:"-" toml:"-"`
}

// Validate 执行验证
func (r *Request) Validate(c *gin.Context, req RequestIFace) error {
	r.Context = c

	if err := c.ShouldBind(req); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return excp.BadRequest("验证器错误")
		}

		errs := err.(validator.ValidationErrors)
		return excp.UnprocessableEntity(errs, errs[0].Translate(valid.GetTrans()), valid.Translate(errs))
	}

	return nil
}
