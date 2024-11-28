package validator

import (
	"reflect"
	"strings"

	"gitee.com/falling-ts/gower/services"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

type Service struct {
	*validator.Validate
	ut.Translator
}

func New() services.ValidatorService {
	return &Service{
		Validate: binding.Validator.Engine().(*validator.Validate),
	}
}

// Init 初始化
func (s *Service) Init(...services.Service) services.Service {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	s.Translator = trans

	if err := zhTrans.RegisterDefaultTranslations(s.Validate, trans); err != nil {
		panic(err)
	}

	s.Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		if name := field.Tag.Get("zh"); name != "" {
			return name
		}
		return strings.ToLower(field.Name)
	})

	return s
}

// Translate 翻译错误
func (s *Service) Translate(errs validator.ValidationErrors) validator.ValidationErrorsTranslations {
	trans := make(validator.ValidationErrorsTranslations)

	var f validator.FieldError

	for i := 0; i < len(errs); i++ {
		f = errs[i].(validator.FieldError)
		rawKey := f.StructField()
		j := strings.LastIndex(rawKey, ".")
		key := strings.ToLower(rawKey[j+1:])

		trans[key] = f.Translate(s.Translator)
	}

	return trans
}

// GetTrans 获取翻译器
func (s *Service) GetTrans() ut.Translator {
	return s.Translator
}
