package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/falling-ts/gower/services"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type DB struct {
	*Service
}

// Set 设置 DBLogger
func (d *DB) Set(arg any) services.DBLogger {
	switch arg.(type) {
	case *Service:
		d.Service = arg.(*Service)
	case services.DBLogger:
		d.Service.DBLogger = arg.(services.DBLogger)
	}

	return d
}

func (d *DB) LogMode(logger.LogLevel) logger.Interface {
	return d
}

func (d *DB) Info(ctx context.Context, msg string, data ...interface{}) {
	_ = ctx
	d.Service.Logger.Debug("DB Debug", zap.String("db", fmt.Sprintf(msg, data...)))
}

func (d *DB) Warn(ctx context.Context, msg string, data ...interface{}) {
	_ = ctx
	d.Service.Logger.Warn("DB Warn", zap.String("db", fmt.Sprintf(msg, data...)))
}

func (d *DB) Error(ctx context.Context, msg string, data ...interface{}) {
	_ = ctx
	d.Service.Logger.Error("DB Error",
		zap.String("db", fmt.Sprintf(msg, data...)),
		zap.Stack("stack"))
}

func (d *DB) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	_ = ctx
	elapsed := time.Since(begin)
	switch {
	case err != nil && d.Service.Core().Enabled(zap.ErrorLevel):
		sql, rows := fc()
		d.Service.Logger.Error("Error Trace", zap.Error(err), zap.String("sql", sql), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows))
	case elapsed > 200*time.Millisecond && d.Service.Core().Enabled(zap.WarnLevel):
		sql, rows := fc()
		d.Service.Logger.Warn("Slow Trace", zap.String("sql", sql), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows))
	case d.Service.Core().Enabled(zap.DebugLevel):
		sql, rows := fc()
		d.Service.Logger.Debug("Debug Trace", zap.String("sql", sql), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows))
	}
}
