package services

// Service 服务的通用接口
type Service interface {
	Register(Services)
}

// Services 服务集合的接口
type Services interface {
	Mount()
	SetService(Service)
}
