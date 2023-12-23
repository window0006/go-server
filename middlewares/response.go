package middlewares

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/window0006/go-server/apis"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	responseBody *apis.ResponseBody // 声明为指针类型 responseBody 可以被修改
}

// context 的 writeJSON 最终会调用 ResponseWriter 的 Write 方法，可以重写
func (cw *CustomResponseWriter) Write(data []byte) (int, error) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err == nil {
		// 要先做类型转换
		cw.responseBody.Retcode = int(jsonData["retcode"].(float64))
		cw.responseBody.Message = string(jsonData["message"].(string))
	}
	return cw.ResponseWriter.Write(data)
}

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 存到 context 中
		c.Writer = &CustomResponseWriter{c.Writer, &apis.ResponseBody{}}
		c.Next()
	}
}
