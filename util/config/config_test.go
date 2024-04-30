package config

import (
	"fmt"
	"reflect"
	"testing"
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

func TestConfig(t *testing.T) {
	var cfg Cfg

	ret := reflect.TypeOf(&cfg).Elem()
	if ret.Kind() == reflect.Struct {
		for i := 0; i < ret.NumField(); i++ {
			field := ret.Field(i)
			fieldRet := reflect.TypeOf(field)
			if fieldRet.Kind() == reflect.Struct {
				for i := 0; i < fieldRet.NumField(); i++ {
					fmt.Println(fieldRet.Field(i).Type)
				}
			}
		}
	}

	//if ret.Elem().Kind() == reflect.Struct {
	//	for i := 0; i < ret.Elem().NumField(); i++ {
	//		fmt.Println(ret.Elem().Field(i).Name, ret.Elem().Field(i).Type)
	//	}
	//}

}
