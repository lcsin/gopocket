package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    ErrCode `json:"code"`
	State   bool    `json:"state"`
	Message string  `json:"message"`
	Data    any     `json:"data"`
}

func ResponseOK(c *gin.Context, data interface{}) {
	resp := Response{
		Code:    ErrSuccess,
		State:   true,
		Message: ErrSuccess.String(),
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func ResponseError(c *gin.Context, code ErrCode) {
	resp := Response{
		Code:    code,
		State:   false,
		Message: code.String(),
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}

func ResponseErrorMessage(c *gin.Context, code ErrCode, message string) {
	if message == "" {
		message = code.String()
	}
	resp := Response{
		Code:    code,
		State:   false,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}
