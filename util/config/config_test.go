package config

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Cfg struct {
	App struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"app"`
	Nacos struct {
		IP        string `yaml:"ip"`
		Port      int    `yaml:"port"`
		Namespace string `yaml:"namespace"`
		Group     string `yaml:"group"`
		DataId    string `yaml:"dataid"`
	}
}

var cfg Cfg

func TestConfig(t *testing.T) {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		_ = viper.Unmarshal(&cfg)
	})

	time.Sleep(time.Second)
	go func() {
		for {
			fmt.Println(cfg.App, cfg.Nacos)
			time.Sleep(time.Second)
		}
	}()

	select {}
}
