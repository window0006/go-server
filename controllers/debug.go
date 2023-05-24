package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/window0006/go-server/apis"
)

type Debug struct{}

func (d *Debug) Hello(c *gin.Context) {
	c.JSON(200, apis.PrintHelloworld())
}
