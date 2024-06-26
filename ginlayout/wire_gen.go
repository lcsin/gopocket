// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/ginlayout/internal/handler"
	"github.com/lcsin/ginlayout/internal/repository"
	"github.com/lcsin/ginlayout/internal/repository/cache"
	"github.com/lcsin/ginlayout/internal/repository/dao"
	"github.com/lcsin/ginlayout/internal/service"
	"github.com/lcsin/ginlayout/ioc"
)

// Injectors from wire.go:

func InitWebServer() *gin.Engine {
	cmdable := ioc.InitRedis()
	v := ioc.InitMiddlewares(cmdable)
	db := ioc.InitDB()
	iHelloDAO := dao.NewHelloDAO(db)
	iHelloCache := cache.NewHelloCache(cmdable)
	iHelloRepository := repository.NewHelloRepository(iHelloDAO, iHelloCache)
	iHelloService := service.NewHelloService(iHelloRepository)
	helloHandler := handler.NewHelloHandler(iHelloService)
	engine := ioc.InitWebServer(v, helloHandler)
	return engine
}
