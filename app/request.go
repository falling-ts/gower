package app

import (
	"crypto/subtle"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
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
	CsrfToken    string `form:"csrf_token" json:"csrf_token" xml:"csrf_token" uri:"csrf_token"`
}

// Validate 执行验证
func (r *Request) Validate(c *gin.Context, req RequestIFace) error {
	r.Context = c

	var ok bool
	if err := c.ShouldBind(req); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			return excp.BadRequest("验证器错误")
		}

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return excp.BadRequest(err)
		}
		return excp.UnprocessableEntity(errs, errs[0].Translate(valid.GetTrans()), valid.Translate(errs))
	}

	realToken := c.GetString("csrf_token")
	if realToken == "" {
		return nil
	}

	csrfField := reflect.Indirect(reflect.ValueOf(req)).FieldByName("CsrfToken")
	if !csrfField.IsValid() {
		return excp.BadRequest("CSRF 校验失败")
	}

	csrfToken, ok := csrfField.Interface().(string)
	if !ok {
		return excp.BadRequest("CSRF 校验失败")
	}

	if subtle.ConstantTimeCompare([]byte(realToken), []byte(csrfToken)) == 0 {
		return excp.NotAcceptable("CSRF 校验失败")
	}

	return nil
}
