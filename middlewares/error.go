package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// 每个请求处理过程中使用 recover 函数来捕获 panic
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 处理 panic
				filelog.Infof("Recovered from panic: %v\n%s", r, debug.Stack())
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
