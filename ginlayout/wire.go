//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/ginlayout/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 基础第三方服务
		ioc.InitDB, ioc.InitRedis,
		// dao
		// ..
		// repository
		// ..
		// service
		// ..
		// handler
		// ..
		// web server
		ioc.InitMiddlewares, ioc.InitWebServer,
	)
	return gin.Default()
}
