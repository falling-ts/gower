package providers

import (
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/services/validator"
)

var _ services.ValidatorService = (*validator.Service)(nil)

func init() {
	P.Register("validator", func(...services.Service) services.Service {
		return validator.New().Init()
	})

}
