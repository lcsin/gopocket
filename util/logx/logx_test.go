package logx

import (
	"testing"
)

func TestLogx(t *testing.T) {
	Info("zhangsan", "hello, world!")
	Infof("你好我叫%s, 我的年龄是%d", "zhangsan", 18)

	Error("zhangsan", "hello, world!")
	Errorf("你好我叫%s, 我的年龄是%d", "zhangsan", 18)
}
