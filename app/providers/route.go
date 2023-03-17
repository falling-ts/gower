package providers

import (
	"gower/services"
	"gower/services/route"
)

var _ services.RouteService = (*route.Service)(nil)

func init() {
	ss.Route = route.New()
}
