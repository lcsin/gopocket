package ioc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/ginlayout/internal/handler/middleware"
	"github.com/redis/go-redis/v9"
)

func InitWebServer(middlewares []gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares...)
	v1 := r.Group("/api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func InitMiddlewares(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		// 跨域中间件
		middleware.CORS(),
		// JWT中间件
		middleware.Jwt(),
		// 限流中间件
		//ratelimit.NewBuilder(redisClient, time.Minute, 100).Build(),
	}
}
