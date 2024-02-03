package logx

var logx Logger

type Logger interface {
	Debug(i ...any)
	Debugf(msg string, i ...any)

	Info(i ...interface{})
	Infof(msg string, i ...any)

	Error(i ...interface{})
	Errorf(msg string, i ...any)
}

func SetLogger(log Logger) {
	logx = log
}

func Debug(a ...any) {
	logx.Debug(a...)
}
func Debugf(msg string, a ...any) {
	logx.Debugf(msg, a...)
}

func Info(a ...any) {
	logx.Info(a...)
}
func Infof(msg string, a ...any) {
	logx.Infof(msg, a...)
}

func Error(a ...any) {
	logx.Error(a...)
}
func Errorf(msg string, a ...any) {
	logx.Errorf(msg, a...)
}
