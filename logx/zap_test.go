package logx

import (
	"testing"

	"go.uber.org/zap"
)

func TestExample(t *testing.T) {
	log := zap.NewExample()

	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")

	log.Info("info field log", zap.String("k1", "v1"), zap.Int("k2", 9))
}

func TestDevelopment(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")

	log.Info("info field log", zap.String("k1", "v1"), zap.Int("k2", 9))
}

func TestProduction(t *testing.T) {
	log, err := zap.NewProduction(zap.Development())
	if err != nil {
		panic(err)
	}

	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")
}

func TestDefaultConsoleLogger(t *testing.T) {
	SetLogger(GetZapLogger())

	Debug("hello, world!")
	Debugf("hello: %v", "张三")

	Info("hello, world!")
	Infof("hello: %v", "张三")

	Error("unknown error")
	Errorf("unknown error: %v", "internal server error")
}
