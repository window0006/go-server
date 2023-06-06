package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/window0006/go-server/dao/entity"
	"github.com/window0006/go-server/middlewares"
	"github.com/window0006/go-server/routers"
)

func main() {
	router := gin.New()

	env := os.Getenv("ENV")
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	entity.DB.Init()

	router.Use(middlewares.Request())
	router.Use(middlewares.Response())
	router.Use(middlewares.Logs())
	router.Use(middlewares.Recovery())
	// 路由
	routers.SetupRouter(router)

	router.Run(":8080")
}
