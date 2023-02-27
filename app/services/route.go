package services

import "github.com/gin-gonic/gin"

type RouteService interface {
	Run()
}

type Route struct {
	*gin.Engine
}

func (r *Route) Init() {
	r.Engine = gin.Default()
}

func init() {
	services.Register(new(Route))
}
