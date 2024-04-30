package api

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/gopocket/util/cryptor"
)

// Sign 接口签名
func Sign(params map[string]string, key string) string {
	var fields []string
	for k, _ := range params {
		fields = append(fields, k)
	}
	sort.Strings(fields)

	var sign string
	for _, v := range fields {
		sign += v
		val, _ := params[v]
		sign += val
	}
	sign += key

	return cryptor.MD5(sign)
}

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
