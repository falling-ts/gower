package providers

import "gower/services/config"

var _ ConfigService = (*config.Config)(nil)

type ConfigService interface {
	Service
}
