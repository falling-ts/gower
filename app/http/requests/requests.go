package requests

import (
	"reflect"
	"strings"

	"gower/app"
	"gower/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var excp = app.Exceptions()

// Request 通用请求接口
type Request interface {
	Validate(ctx *gin.Context, req Request) error
	SetContext(c *gin.Context)
}

type request struct {
	*gin.Context
}

var trans ut.Translator

func init() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := zhTrans.RegisterDefaultTranslations(v, trans); err != nil {
			panic(err)
		}
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			if name := field.Tag.Get("zh"); name != "" {
				return name
			}
			return strings.ToLower(field.Name)
		})
	}
}

// Validate 执行验证
func (r *request) Validate(ctx *gin.Context, req Request) error {
	r.Context = ctx

	if err := ctx.ShouldBind(req); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return excp.BadRequest("验证器错误.")
		}

		errs := err.(validator.ValidationErrors)
		return excp.BadRequest(errs, errs[0].Translate(trans), utils.MapStrStr(errs.Translate(trans)).MapKeys(func(key string) string {
			i := strings.LastIndex(key, ".")
			return strings.ToLower(key[i+1:])
		}))
	}

	return nil
}

// SetContext 设置 gin 的请求体
func (r *request) SetContext(c *gin.Context) {
	r.Context = c
}
