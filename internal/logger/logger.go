package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomGinLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func New(level slog.Level) *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%s] - ", level.String()), log.LUTC|log.LstdFlags|log.Lshortfile)
}
