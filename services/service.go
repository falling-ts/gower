package services

// Service 服务通用接口
type Service interface {
	Init(...Service) Service
}
