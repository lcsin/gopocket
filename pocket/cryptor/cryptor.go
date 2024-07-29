package cryptor

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MD5 生成32位MD5
func MD5(text string) string {
	md := md5.New()
	md.Write([]byte(text))
	return strings.ToUpper(hex.EncodeToString(md.Sum(nil)))
}
