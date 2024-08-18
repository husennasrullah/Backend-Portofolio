package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// color
const (
	Red    = "41"
	Yellow = "43"
	Green  = "42"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
)

func HTTPLogger(param gin.LogFormatterParams) string {
	var color string
	switch {
	case param.StatusCode >= 500:
		color = Red
	case param.StatusCode >= 400:
		color = Yellow
	default:
		color = Green
	}

	requestBody, _ := param.Keys["RequestBody"].(string)
	responseBody, _ := param.Keys["ResponseBody"].(string)
	serviceError, _ := param.Keys["ServiceError"].(string)

	title := fmt.Sprintf("[\033[%smLOGGING HTTP\033[0m]", color)
	return fmt.Sprintf("%s [Time: %s] [Status: \033[%sm%d\033[0m] [Method: %s] [Path: %s] [Latency: %d] [Error: %s] [UserAgent: %s] [RequestBody: %s] [ResponseBody: %s]\n",
		title,
		param.TimeStamp.Format("2006-01-02 15:04:05"),
		color,
		param.StatusCode,
		param.Method,
		param.Path,
		param.Latency,
		serviceError,
		param.Request.UserAgent(),
		requestBody,
		responseBody,
	)
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor string
	switch entry.Level {
	case logrus.InfoLevel:
		levelColor = colorBlue
	case logrus.WarnLevel:
		levelColor = colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = colorRed
	default:
		levelColor = colorReset
	}

	logs := fmt.Sprintf(
		"%s[LOGGING HTTP]%s [Time: %s] [Status: %d] [Method: %s] [Path: %s] [Latency: %d] [Error: %v] [UserAgent: %s] [RequestBody: %s] [ResponseBody: %s]%s\n",
		levelColor, colorReset,
		entry.Time.Format("2006-01-02 15:04:05"),
		entry.Data["status"],
		entry.Data["method"],
		entry.Data["path"],
		entry.Data["latency"],
		entry.Data["error"],
		entry.Data["user_agent"],
		entry.Data["request_body"],
		entry.Data["response_body"],
		colorReset,
	)
	return []byte(logs), nil
}
