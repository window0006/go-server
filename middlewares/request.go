package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用 crypto/md5 生成 requestID, 存到 context 中
		requestId := uuid.New().String()
		c.Set("requestId", requestId)
		c.Writer.Header().Set("X-Request-Id", string(requestId))
		c.Next()
	}
}
