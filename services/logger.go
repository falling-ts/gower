package services

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

// LoggerService 日志服务
type LoggerService interface {
	Service

	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
	Core() zapcore.Core

	Debug(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Level() zapcore.Level
	Log(lvl zapcore.Level, msg string, fields ...zap.Field)
	Named(s string) *zap.Logger
	Panic(msg string, fields ...zap.Field)

	Sugar() *zap.SugaredLogger
	Sync() error

	WithOptions(opts ...zap.Option) *zap.Logger
	With(fields ...zap.Field) *zap.Logger
	Warn(msg string, fields ...zap.Field)

	Zap() *zap.Logger

	DB() DBLogger
}

type DBLogger interface {
	Set(arg any) DBLogger
	logger.Interface
}
