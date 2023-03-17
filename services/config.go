package services

// Config 配置内容接口
type Config interface {
	Set(arg any)
	Get(fieldStr string, args ...any) any
}

// ConfigService 配置服务接口
type ConfigService interface {
	Service
	Get(fieldStr string, args ...any) any
}
