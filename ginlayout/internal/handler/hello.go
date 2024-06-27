package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/ginlayout/internal/service"
)

type HelloHandler struct {
	srv service.IHelloService
}

func NewHelloHandler(srv service.IHelloService) *HelloHandler {
	return &HelloHandler{srv: srv}
}

func (h *HelloHandler) RegisterRoutes(v1 *gin.RouterGroup) {
	v1.GET("/sayHello", func(c *gin.Context) {

		// 调用service层逻辑
		h.srv.SayHello(c, 001)

		c.String(http.StatusOK, "hello, world!")
	})
}
