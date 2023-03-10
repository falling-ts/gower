package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"gower/services"
	"reflect"
	"strings"
)

type Validator struct {
	*validator.Validate
	ut.Translator
}

func New() services.Validator {
	return &Validator{
		Validate: binding.Validator.Engine().(*validator.Validate),
	}
}

// Init 初始化
func (v *Validator) Init(args ...any) {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	v.Translator = trans

	if err := zhTrans.RegisterDefaultTranslations(v.Validate, trans); err != nil {
		panic(err)
	}

	v.Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		if name := field.Tag.Get("zh"); name != "" {
			return name
		}
		return strings.ToLower(field.Name)
	})

}

// Translate 翻译错误群
func (v *Validator) Translate(errs validator.ValidationErrors) validator.ValidationErrorsTranslations {
	trans := make(validator.ValidationErrorsTranslations)

	var f validator.FieldError

	for i := 0; i < len(errs); i++ {
		f = errs[i].(validator.FieldError)
		rawKey := f.StructField()
		j := strings.LastIndex(rawKey, ".")
		key := strings.ToLower(rawKey[j+1:])

		trans[key] = f.Translate(v.Translator)
	}

	return trans
}

// GetTrans 获取翻译器
func (v *Validator) GetTrans() ut.Translator {
	return v.Translator
}
