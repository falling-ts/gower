package services

type Configs interface {
	Set(arg any)
	Get(fieldStr string, args ...any) any
}

// Config 适配接口
type Config interface {
	Service
	Get(fieldStr string, args ...any) any
}
