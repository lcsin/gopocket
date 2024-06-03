package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/gopocket/util/httpx"
)

func main() {
	r := gin.Default()

	httpx.Graceful(r, ":8080")
}
