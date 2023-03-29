package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const divTime = time.Hour

func consoleInfoLogger() zapcore.Core {
	encoderConfig := getEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	level := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		if mode != "development" {
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

// 信息日志分割器
type infoLogDivider struct {
	*os.File
	updateAt time.Time
}

// Write 执行日志分割, 并写入日志
func (i *infoLogDivider) Write(p []byte) (n int, err error) {
	if i.updateAt.Before(time.Now()) {
		i.File = getInfoFile()
		i.updateAt = time.Now().Add(divTime)
	}

	return i.File.Write(p)
}

func fileInfoLogger() zapcore.Core {
	encoderConfig := getEncoderConfig()

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	level := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		if mode == "production" {
			return l == zapcore.InfoLevel
		}

		return l <= zapcore.InfoLevel
	})

	return zapcore.NewCore(encoder, zapcore.AddSync(newInfoLogDivider()), level)
}

func newInfoLogDivider() *infoLogDivider {
	il := new(infoLogDivider)

	il.File = getInfoFile()
	il.updateAt = time.Now().Add(divTime)

	return il
}

func getInfoFile() *os.File {
	var logFile string
	var logDir string

	channel, ok := config.Get("log.channel", "").(string)
	if !ok {
		panic("获取配置错误.")
	}
	dir := config.Get("log.dir").(string)
	util.CreateDir(dir)

	now := time.Now().Local()
	flatDay := now.Format("2006-01-02")
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	hour := now.Format("15")

	logFile = map[string]func() string{
		"stack": func() string {
			return fmt.Sprintf("%s/gower.log", dir)
		},
		"flat-day": func() string {
			return fmt.Sprintf("%s/gower.%s.log", dir, flatDay)
		},
		"day": func() string {
			logDir = filepath.Join(dir, year, month)
			util.CreateDir(dir)
			return fmt.Sprintf("%s/%s.log", logDir, day)
		},
		"hour": func() string {
			logDir = filepath.Join(dir, year, month, day)
			util.CreateDir(dir)
			return fmt.Sprintf("%s/%s.log", logDir, hour)
		},
	}[channel]()
	if logFile == "" {
		logFile = fmt.Sprintf("%s/gower.log", dir)
	}

	return getFile(logFile)
}

// 错误日志分割器
type errorLogDivider struct {
	*os.File
	updateAt time.Time
}

// Write 执行日志分割, 并写入日志
func (e *errorLogDivider) Write(p []byte) (n int, err error) {
	if e.updateAt.Before(time.Now()) {
		e.File = getErrorFile()
		e.updateAt = time.Now().Add(divTime)
	}

	return e.File.Write(p)
}

func fileErrorLogger() zapcore.Core {
	encoderConfig := getEncoderConfig()

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	level := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.ErrorLevel
	})

	return zapcore.NewCore(encoder, zapcore.AddSync(newErrorLogDivider()), level)
}

func newErrorLogDivider() *errorLogDivider {
	el := new(errorLogDivider)

	el.File = getErrorFile()
	el.updateAt = time.Now().Add(divTime)

	return el
}

func getErrorFile() *os.File {
	dir := config.Get("log.dir").(string)
	flatDay := time.Now().Local().Format("2006-01-02")
	logFile := fmt.Sprintf("%s/error.%s.log", dir, flatDay)

	return getFile(logFile)
}

func getEncoderConfig() zapcore.EncoderConfig {
	encodeDuration, ok := map[string]zapcore.DurationEncoder{
		"seconds": zapcore.SecondsDurationEncoder,
		"nanos":   zapcore.NanosDurationEncoder,
		"millis":  zapcore.MillisDurationEncoder,
		"string":  zapcore.StringDurationEncoder,
	}[config.Get("log.durationFormat", "caller").(string)]
	if !ok {
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

func getFile(file string) *os.File {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(file)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	return f
}
