package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func consoleInfoLogger() zapcore.Core {
	encoderConfig := getEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	level := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		if mode != "debug" {
			return l == zapcore.InfoLevel
		}

		return l <= zapcore.InfoLevel
	})
	return zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level)
}

func consoleErrorLogger() zapcore.Core {
	encoderConfig := getEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	level := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.ErrorLevel
	})

	return zapcore.NewCore(encoder, zapcore.Lock(os.Stderr), level)
}

func fileInfoLogger() zapcore.Core {
	encoderConfig := getEncoderConfig()

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	level := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		if mode == "release" {
			return l == zapcore.InfoLevel
		}

		return l <= zapcore.InfoLevel
	})

	return zapcore.NewCore(encoder, zapcore.AddSync(getInfoFile()), level)
}

func getInfoFile() *os.File {
	var logFile string
	var logDir string

	channel, ok := config.Get("log.channel", "").(string)
	if !ok {
		panic("获取配置错误.")
	}
	dir := config.Get("log.dir").(string)
	createDir(dir)

	now := time.Now().Local()
	flatDay := now.Format("2006-01-02")
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	hour := now.Format("15")

	switch channel {
	case "flat-day":
		logFile = fmt.Sprintf("%s/gower.%s.log", dir, flatDay)
	case "day":
		logDir = filepath.Join(dir, year, month)
		createDir(logDir)
		logFile = fmt.Sprintf("%s/gower.%s-day.log", logDir, day)
	case "hour":
		logDir = filepath.Join(dir, year, month, day)
		createDir(logDir)
		logFile = fmt.Sprintf("%s/gower.%s-hour.log", logDir, hour)
	default:
		logFile = fmt.Sprintf("%s/gower.log", dir)
	}

	// 判断日志文件是否存在，如果存在则打开文件，否则创建文件
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(logFile)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	return f
}

func fileErrorLogger() zapcore.Core {
	encoderConfig := getEncoderConfig()

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	level := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.ErrorLevel
	})

	return zapcore.NewCore(encoder, zapcore.AddSync(getErrorFile()), level)
}

func getErrorFile() *os.File {
	dir := config.Get("log.dir").(string)
	flatDay := time.Now().Local().Format("2006-01-02")
	logFile := fmt.Sprintf("%s/error.%s.log", dir, flatDay)

	// 判断日志文件是否存在，如果存在则打开文件，否则创建文件
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(logFile)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	return f
}

func getEncoderConfig() zapcore.EncoderConfig {
	var encodeDuration zapcore.DurationEncoder
	switch config.Get("log.durationFormat", "caller").(string) {
	case "nanos":
		encodeDuration = zapcore.NanosDurationEncoder
	case "millis":
		encodeDuration = zapcore.MillisDurationEncoder
	case "string":
		encodeDuration = zapcore.StringDurationEncoder
	default:
		encodeDuration = zapcore.SecondsDurationEncoder
	}

	return zapcore.EncoderConfig{
		MessageKey:       config.Get("log.msgKey", "msg").(string),
		TimeKey:          config.Get("log.timeKey", "ts").(string),
		LevelKey:         config.Get("log.levelKey", "level").(string),
		NameKey:          config.Get("log.nameKey", "logger").(string),
		CallerKey:        config.Get("log.callerKey", "caller").(string),
		FunctionKey:      zapcore.OmitKey,
		StacktraceKey:    config.Get("log.stackKey", "stack").(string),
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout(config.Get("log.timeFormat", "2006-01-02 15:04:05").(string)),
		EncodeDuration:   encodeDuration,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: config.Get("log.consoleSep", "").(string),
	}
}

func createDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
	}
}
