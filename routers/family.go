package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/window0006/go-server/controllers"
)

func Family(router *gin.Engine) {
	familyController := &controllers.Family{}
	familyRoot := router.Group("/family")
	familyRoot.GET("/list", familyController.List)
	familyRoot.POST("/create", familyController.Create)
}
