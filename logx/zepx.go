package logx

import (
	"fmt"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapx struct {
}

func Default(opts ...zap.Option) *zapx {
	zap.ReplaceGlobals(defaultZapConsoleLogger(opts...))
	return new(zapx)
}

func defaultZapConsoleLogger(options ...zap.Option) *zap.Logger {
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

	atomicLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(defaultZapHook())),
		atomicLevel,
	)
	return zap.New(core, zap.AddCaller(), zap.Development()).WithOptions(options...)
}

func defaultZapHook() *lumberjack.Logger {
	dir := fmt.Sprintf("./log/%v.log", time.Now().Format(time.DateOnly))
	_ = os.MkdirAll(path.Dir(dir), os.ModePerm)
	return &lumberjack.Logger{
		Filename:   dir,   // 日志文件路径
		MaxSize:    100,   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 10,    // 日志文件最多保存多少个备份
		MaxAge:     30,    // 文件最多保存多少天
		Compress:   false, // 是否压缩
	}
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
