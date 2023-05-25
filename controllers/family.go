package controllers

import "github.com/gin-gonic/gin"

type Family struct{}

func (f *Family) List(c *gin.Context) {
	// TODO 从数据库中获取 famliy 列表
	c.JSON(200, gin.H{
		"retcode": 0,
		"message": "success",
		"data": []gin.H{
			{
				"id":   1,
				"name": "张三家族",
			},
			{
				"id":   2,
				"name": "李四家族",
			},
		},
	})
}

func (f *Family) Create(c *gin.Context) {
	// TODO 创建一个新的 family
	c.JSON(200, gin.H{
		"retcode": 0,
		"message": "success",
		"data": gin.H{
			"id":   1,
			"name": "张三家族",
		},
	})
}
