package route

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func setRecovery(engine *gin.Engine) {
	engine.Use(gin.RecoveryWithWriter(errOutput()))
}

func errOutput() io.Writer {
	dir := cfg.Get("log.dir").(string)
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

	return io.MultiWriter(f, os.Stderr)
}
