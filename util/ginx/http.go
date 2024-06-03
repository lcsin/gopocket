package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OK(c *gin.Context, data interface{}) {
	resp := Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func Error(c *gin.Context, code int, errmsg string) {
	resp := Response{
		Code:    code,
		Message: errmsg,
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}
