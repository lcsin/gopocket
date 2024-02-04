package logx

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestDefaultConsoleLogger(t *testing.T) {
	Debug("hello, world!")
	Debugf("hello: %v", "张三")

	Info("hello, world!")
	Infof("hello: %v", "张三")

	Error("unknown error")
	Errorf("unknown error: %v", "internal server error")
}

func TestCustomLogger(t *testing.T) {
	SetLogger(Default(zap.IncreaseLevel(zapcore.InfoLevel)))

	Debug("hello, world!")
	Debugf("hello: %v", "张三")

	Info("hello, world!")
	Infof("hello: %v", "张三")

	Error("unknown error")
	Errorf("unknown error: %v", "internal server error")
}
