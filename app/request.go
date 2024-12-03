package app

import (
	"crypto/subtle"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var (
	exc   = Exception()
	valid = Validator()
)

// RequestIFace 通用请求接口
type RequestIFace interface {
	Validate(c *gin.Context, req RequestIFace) error
}

type Request struct {
	*gin.Context `json:"-" xml:"-" form:"-" query:"-" protobuf:"-" msgpack:"-" yaml:"-" uri:"-" header:"-" toml:"-"`
	CsrfToken    string `form:"csrfToken" json:"csrfToken" xml:"csrfToken" uri:"csrfToken"`
}

type IndexRequest struct {
	Request
	Page    uint `form:"page,default=1" json:"page,default=1" binding:"numeric"`
	PageNum uint `form:"pageNum,default=10" json:"pageNum,default=10" binding:"numeric"`
}

type ModalRequest struct {
	Request
	IsModal bool `form:"isModal,default=false" json:"isModal,default=false" binding:"boolean"`
}

// Validate 执行验证
func (r *Request) Validate(c *gin.Context, req RequestIFace) error {
	r.Context = c

	var ok bool
	if err := c.ShouldBind(req); err != nil {
		if _, ok = err.(*validator.InvalidValidationError); ok {
			return exc.BadRequest("验证器错误")
		}

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return exc.BadRequest(err)
		}
		return exc.UnprocessableEntity(errs, errs[0].Translate(valid.GetTrans()), valid.Translate(errs))
	}

	realToken := c.GetString("csrfToken")
	if realToken == "" {
		return nil
	}

	csrfField := reflect.Indirect(reflect.ValueOf(req)).FieldByName("CsrfToken")
	if !csrfField.IsValid() {
		return exc.BadRequest("CSRF 校验失败")
	}

	csrfToken, ok := csrfField.Interface().(string)
	if !ok {
		return exc.BadRequest("CSRF 校验失败")
	}

	if subtle.ConstantTimeCompare([]byte(realToken), []byte(csrfToken)) == 0 {
		return exc.NotAcceptable("CSRF 校验失败")
	}

	return nil
}
