package logx

import (
	"testing"
)

func TestLogx(t *testing.T) {
	l := new(logx)
	l.Info("zhangsan", "hello, world!")
	l.Infof("你好我叫%s, 我的年龄是%d", "zhangsan", 18)
	l.Errorf("你好我叫%s, 我的年龄是%d", "zhangsan", 18)
}
