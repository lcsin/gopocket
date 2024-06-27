package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRUN(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	type helloReq struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	r.POST("/hello", func(c *gin.Context) {
		var req helloReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, "bad request")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name": req.Name,
			"age":  req.Age,
		})
	})

	RUN(r, ":8080")
}

func TestGet(t *testing.T) {
	resp, err := GET("http://localhost:8080/ping")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}

func TestJSONPost(t *testing.T) {
	type helloReq struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	req := helloReq{
		Name: "张三",
		Age:  18,
	}
	bytes, _ := json.Marshal(&req)
	resp, err := POSTJson("http://localhost:8080/hello", bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}

func TestSign(t *testing.T) {
	req := make(map[string]string)
	req["d"] = "v1"
	req["a"] = "v2"
	req["c"] = ""

	key := "sign key"
	fmt.Println(Sign(req, key))
}
