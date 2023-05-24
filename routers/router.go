package routers

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	// 路由分组
	Debug(router)
	Family(router)
}
