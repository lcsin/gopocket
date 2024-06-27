package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CORSBuilder struct {
	allowMethods  []string
	allowHeaders  []string
	allowOrigins  []string
	exposeHeaders []string
	maxAge        time.Duration
}

func (c *CORSBuilder) AllowMethod(method ...string) *CORSBuilder {
	c.allowHeaders = append(c.allowHeaders, method...)
	return c
}

func (c *CORSBuilder) AllowHeader(header ...string) *CORSBuilder {
	c.allowHeaders = append(c.allowHeaders, header...)
	return c
}

func (c *CORSBuilder) ExposeHeader(header ...string) *CORSBuilder {
	c.exposeHeaders = append(c.exposeHeaders, header...)
	return c
}

func (c *CORSBuilder) AllowOrigin(origin ...string) *CORSBuilder {
	c.allowOrigins = append(c.allowOrigins, origin...)
	return c
}

func (c *CORSBuilder) MaxAge(maxAge time.Duration) *CORSBuilder {
	c.maxAge = maxAge
	return c
}

func (c *CORSBuilder) Build() gin.HandlerFunc {
	if len(c.allowHeaders) == 0 {
		c.allowHeaders = append(c.allowHeaders, "Content-Type")
	}
	if len(c.allowOrigins) == 0 {
		c.allowOrigins = append(c.allowOrigins, "*")
	}

	return cors.New(cors.Config{
		// 是否允许带上用户认证比如cookie
		AllowCredentials: true,
		// 允许哪些请求方式
		AllowMethods: c.allowMethods,
		// 允许哪些头
		AllowHeaders: c.allowHeaders,
		// 允许暴露的头
		ExposeHeaders: c.exposeHeaders,
		// 允许哪些源
		AllowOrigins: c.allowOrigins,
		// 最大时间
		MaxAge: c.maxAge,
	})
}
