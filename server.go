package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/window0006/go-server/apis"
	"github.com/window0006/go-server/middlewares"
)

func main() {
	router := gin.New()

	env := os.Getenv("ENV")
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 中间件
	router.Use(middlewares.Request())
	router.Use(middlewares.Response())
	router.Use(middlewares.Logs())
	router.Use(middlewares.Recovery())
	// 路由
	router.GET("/", func(context *gin.Context) {
		data := apis.PrintHelloworld()
		context.JSON(200, data)
	})

	router.Run(":8080")
}
