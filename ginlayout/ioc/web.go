package ioc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/ginlayout/internal/handler"
	"github.com/lcsin/ginlayout/internal/handler/middleware"
	"github.com/redis/go-redis/v9"
)

func InitWebServer(middlewares []gin.HandlerFunc, helloHandler *handler.HelloHandler) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares...)
	v1 := r.Group("/api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	helloHandler.RegisterRoutes(v1)

	return r
}

func InitMiddlewares(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		// 跨域中间件
		middleware.CORS(),
		// JWT中间件
		middleware.NewJwtBuilder().
			IgnorePaths("/api/v1/ping").
			IgnorePaths("/api/v1/sayHello").
			Build(),
	}
}
