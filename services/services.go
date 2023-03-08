package services

// Services 服务集合的接口
type Services interface {
	Mount() Services
	BindContent() Services
	SetService(Service)
}

// Service 服务的通用接口
type Service interface {
	Register(Services)
}
