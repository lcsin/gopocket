package api

import (
	"sort"

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
