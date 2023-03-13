package services

// Exceptions 异常内容接口
type Exceptions interface {
	error
	Set(arg any)
	Get(arg string) any
	New(code int, args ...any) Exceptions
}

// Exception 服务接口
type Exception interface {
	Service
	Build(args ...any) Exceptions
}
