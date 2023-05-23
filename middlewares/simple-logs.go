package middlewares

import (
	"encoding/json"
	"io"
	internalLog "log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/window0006/go-server/utils"
)

func InternalLogs() gin.HandlerFunc {
	logFileWriter := utils.NewLogFileWriter()
	internalLog.SetOutput(io.MultiWriter(logFileWriter, os.Stdout))

	return func(c *gin.Context) {
		// 存到 context 中
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		logCont := map[string]interface{}{
			"startTime":   startTime,
			"latencyTime": latencyTime,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
			"clientIP":    clientIP,
			"statusCode":  statusCode,
		}
		bytes, err := json.Marshal(logCont)
		if err != nil {
			internalLog.Printf("json.Marshal error: %v", err)
			return
		}
		internalLog.Println(string(bytes))
	}
}
