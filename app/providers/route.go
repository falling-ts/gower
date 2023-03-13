package providers

import (
	"gower/services"
	"gower/services/route"
)

var _ services.Route = (*route.Route)(nil)

func init() {
	s.Route = route.New()
}
