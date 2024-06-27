//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lcsin/ginlayout/internal/handler"
	"github.com/lcsin/ginlayout/internal/repository"
	"github.com/lcsin/ginlayout/internal/repository/cache"
	"github.com/lcsin/ginlayout/internal/repository/dao"
	"github.com/lcsin/ginlayout/internal/service"
	"github.com/lcsin/ginlayout/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 基础第三方服务
		ioc.InitDB, ioc.InitRedis,
		// dao
		dao.NewHelloDAO, cache.NewHelloCache,
		// repository
		repository.NewHelloRepository,
		// service
		service.NewHelloService,
		// handler
		handler.NewHelloHandler,
		// web server
		ioc.InitMiddlewares, ioc.InitWebServer,
	)
	return gin.Default()
}
