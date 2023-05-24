package middlewares

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/window0006/go-server/apis"
	"github.com/window0006/go-server/utils"
)

type customFormatter struct{}

func (formatter *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// buffer 似乎有更好的内存操作过程，因为它是可变大小的，大型字符串拼接时，性能更好，但是这里场景很简单，所以直接用 []byte("xxx") 了
	// var b bytes.Buffer
	// b.WriteString(fmt.Sprintln(entry.Message))
	// return b.Bytes(), nil

	// 类型转换，将 entry.Message 由 string 转换为 []byte 字节切片
	return []byte(fmt.Sprintln(entry.Message)), nil
}

var filelog = logrus.New()
var consolelog = logrus.New()

func Logs() gin.HandlerFunc {
	filelog.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	consolelog.SetFormatter(new(customFormatter))

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		cw := c.Writer.(*CustomResponseWriter) // 断言为 CustomResponseWriter 类型
		businessResult := cw.responseBody
		if businessResult == nil {
			businessResult = &apis.ResponseBody{}
		}

		logCont := logrus.Fields{
			"startTime":   startTime,
			"latencyTime": latencyTime,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
			"clientIP":    clientIP,
			"statusCode":  statusCode,
		}
		bytes, err := json.Marshal(logCont)
		if err != nil {
			consolelog.Errorf("json.Marshal error: %v", err)
			return
		}
		logTextCont := string(bytes)

		temp := filelog.WithFields(logCont)
		logFn := temp.Infof
		consolelogFn := os.Stdout
		if statusCode >= 500 {
			logFn = temp.Errorf
			consolelogFn = os.Stderr
		} else if statusCode >= 400 {
			logFn = temp.Warnf
			consolelogFn = os.Stderr
		} else if businessResult.Retcode != 0 {
			logFn = temp.Warnf
			consolelogFn = os.Stderr
		}
		filelog.SetOutput(utils.NewLogFileWriter())
		logFn("request handled") // 此时才会写入日志文件

		consolelog.SetOutput(consolelogFn)
		consolelog.Printf("[%s] - %v", startTime.Format("2006-01-02T15:04:05.000"), logTextCont)
	}
}
