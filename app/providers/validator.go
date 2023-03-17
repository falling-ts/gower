package providers

import (
	"gower/services"
	"gower/services/validator"
)

var _ services.ValidatorService = (*validator.Service)(nil)

func init() {
	ss.Validator = validator.New()
}
