package logger

import (
	"github.com/falling-ts/gower/services"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Service struct {
	*zap.Logger
	DBLogger services.DBLogger
}

var (
	config services.Config
	util   services.UtilService
	mode   string
)

// New 新建日志服务
func New() *Service {
	return new(Service)
}

// Init 初始化 logger
func (s *Service) Init(args ...services.Service) services.Service {
	config = args[0].(services.Config)
	util = args[1].(services.UtilService)
	mode = config.Get("app.mode", "test").(string)

	core := zapcore.NewTee(
		consoleInfoLogger(),
		consoleErrorLogger(),
		fileInfoLogger(),
		fileErrorLogger())
	s.Logger = zap.New(core).Named(config.Get("app.name", "Gower").(string))
	s.DBLogger = new(DB).Set(s)

	return s
}

// Zap 获取 zap logger
func (s *Service) Zap() *zap.Logger {
	return s.Logger
}

// DB 获取 DB 的日志输出器
func (s *Service) DB() services.DBLogger {
	return s.DBLogger
}
