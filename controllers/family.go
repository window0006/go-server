package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/window0006/go-server/bindings"
	"github.com/window0006/go-server/dao/entity"
)

type Family struct{}

func (f *Family) List(c *gin.Context) {
	var query bindings.FamilyListQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	// 通过 db 获取 family 列表
	familyList, err := entity.GetFamilyList(&query)
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器错误")
		return
	}

	c.JSON(200, gin.H{
		"retcode": 0,
		"message": "success",
		"data": gin.H{
			"list": familyList,
		},
	})
}

func (f *Family) Create(c *gin.Context) {
	// TODO 创建一个新的 family
	var body bindings.CreateFamilyBody
	c.ShouldBind(&body)
	// 通过 db 存入数据库

	c.JSON(200, gin.H{
		"retcode": 0,
		"message": "success",
		"data": gin.H{
			"id":   1,
			"name": "张三家族",
		},
	})
}
