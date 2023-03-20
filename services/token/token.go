package token

import (
	"github.com/golang-jwt/jwt/v5"
	"gower/services"
)

type Service struct {
	jwt.Token
}

func New() *Service {
	return new(Service)
}

func (s *Service) Init(args ...services.Service) services.Service {
	return s
}
