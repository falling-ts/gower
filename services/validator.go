package services

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Service
	Translate(errs validator.ValidationErrors) validator.ValidationErrorsTranslations
	GetTrans() ut.Translator
}
