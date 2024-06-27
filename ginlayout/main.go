package main

import (
	"fmt"

	"github.com/lcsin/ginlayout/ioc"
	"github.com/lcsin/gopocket/util/httpx"
	"github.com/spf13/viper"
)

func main() {
	ioc.InitLocalConfig()
	ioc.InitLogger()
	r := InitWebServer()
	httpx.Graceful(r, fmt.Sprintf(":%s", viper.Get("server.port").(string)))
}
