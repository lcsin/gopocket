package logx

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapx struct {
}

func GetZapLogger() Logger {
	zap.ReplaceGlobals(DefaultZapConsoleLogger())
	return new(zapx)
}

func DefaultZapConsoleLogger() *zap.Logger {
	cfg := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "time",
		CallerKey:        "caller",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseColorLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: "\t|\t",
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.AddSync(os.Stdout), zap.InfoLevel)
	return zap.New(core, zap.AddCaller(), zap.Development())
}

func (z *zapx) Debug(a ...any) {
	zap.L().Debug(fmt.Sprint(a...))
}
func (z *zapx) Debugf(msg string, a ...any) {
	zap.L().Debug(fmt.Sprintf(msg, a...))
}

func (z *zapx) Info(a ...any) {
	zap.L().Info(fmt.Sprint(a...))
}
func (z *zapx) Infof(msg string, a ...any) {
	zap.L().Info(fmt.Sprintf(msg, a...))
}

func (z *zapx) Error(a ...any) {
	zap.L().Error(fmt.Sprint(a...))
}
func (z *zapx) Errorf(msg string, a ...any) {
	zap.L().Error(fmt.Sprintf(msg, a...))
}
