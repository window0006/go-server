package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/window0006/go-server/controllers"
)

func Debug(router *gin.Engine) {
	debugController := &controllers.Debug{}
	debugRoot := router.Group("/debug")
	debugRoot.GET("/hello", debugController.Hello)
}
