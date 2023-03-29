package config

import (
	"github.com/gin-gonic/gin"

	"tdp-aiart/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/config")

	// 匿名接口

	{
		rg.POST("/ui", uiOption)
	}

	// 需授权接口

	rg.Use(midware.AuthGuard())

	{
		rg.POST("/list", list)
		rg.POST("/create", create)
		rg.POST("/detail", detail)
		rg.POST("/update", update)
		rg.POST("/delete", delete)
	}

}
