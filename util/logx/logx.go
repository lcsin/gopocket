package logx

import (
	"fmt"
	"log"
	"strings"

	"github.com/gookit/color"
)

type logx struct {
}

const (
	infoFormat  = "  [INFO]  ->  "
	errorFormat = "  [ERROR]  ->  "
)

func (l *logx) Info(a ...any) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(infoFormat))
	for _, v := range a {
		sb.WriteString(fmt.Sprintf("%v ", v))
	}

	log.Println(color.FgGreen.Render(sb.String()))
}

func (l *logx) Infof(format string, a ...any) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(infoFormat))
	sb.WriteString(fmt.Sprintf(format, a...))

	log.Println(color.FgGreen.Render(sb.String()))
}

func (l *logx) Error(a ...any) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(errorFormat))
	for _, v := range a {
		sb.WriteString(fmt.Sprintf("%v ", v))
	}

	log.Println(color.FgRed.Render(sb.String()))
}

func (l *logx) Errorf(format string, a ...any) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(errorFormat))
	sb.WriteString(fmt.Sprintf(format, a...))

	log.Println(color.FgRed.Render(sb.String()))
}
