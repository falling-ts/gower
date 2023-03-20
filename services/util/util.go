package util

import "gower/services"

type Service struct{}

func New() *Service {
	return new(Service)
}

func (s *Service) Init(...services.Service) services.Service {
	return s
}
