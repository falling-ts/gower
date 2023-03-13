package route

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func setLogger(engine *gin.Engine) {
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: logFormatter,
		Output:    output(),
		SkipPaths: configs.Get("log.skipPaths").([]string),
	}))
}

func output() io.Writer {
	var logFile string
	var logDir string

	channel, ok := configs.Get("log.channel", "").(string)
	if !ok {
		panic("获取配置错误.")
	}
	dir := configs.Get("log.dir").(string)
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

	return io.MultiWriter(f, os.Stdout)
}

func createDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
	}
}

var logFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor, ResBody string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	if val, ok := param.Keys["body-logger"]; ok && val != nil {
		ResBody, _ = val.(string)
		ResBody = fmt.Sprintf("%s\n", ResBody)
	}

	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		ResBody,
		param.ErrorMessage,
	)
}
