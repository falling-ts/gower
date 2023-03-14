package providers

import (
	"gower/services"
	"gower/services/validator"
)

var _ services.Validator = (*validator.Validator)(nil)

func init() {
	s.Validator = validator.New()
}
