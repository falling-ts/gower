package util

type Service struct{}

func New() *Service {
	return new(Service)
}

func (s *Service) Init(...any) {}
