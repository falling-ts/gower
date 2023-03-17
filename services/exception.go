package services

// Exception 异常内容接口
type Exception interface {
	error
	Set(arg any)
	Get(arg string) any
	New(code int, args ...any) Exception
}

// ExceptionService 异常服务接口
type ExceptionService interface {
	Service
	Build(args ...any) Exception
}
