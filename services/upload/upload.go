package upload

import "gitee.com/falling-ts/gower/services"

// Service 上传服务
type Service struct {
	services.Storage
}

var (
	config services.Config
	util   services.UtilService
)

// New 新建服务
func New() *Service {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(args ...services.Service) services.Service {
	config = args[0].(services.Config)
	util = args[1].(services.UtilService)

	s.Storage = s.storage(config.Get("upload.storage", "local").(string))
	return s
}

// Store 获取指定仓库
func (s *Service) Store(storage string) services.Storage {
	return s.storage(storage)
}

func (s *Service) storage(storage string) services.Storage {
	switch storage {
	default:
		return newLocal()
	}
}
