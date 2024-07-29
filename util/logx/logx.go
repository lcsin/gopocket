package logx

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/lcsin/gopocket/pocket/fileutil"
)

const (
	infoFormat  = "  INFO  "
	errorFormat = "  ERROR  "
)

func Info(a ...any) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(infoFormat))
	for _, v := range a {
		sb.WriteString(fmt.Sprintf("%v ", v))
	}

	log.Println(color.FgGreen.Render(sb.String()))
	saveLog(sb.String(), "info")
}

func Infof(format string, a ...any) {
	log.Println(color.FgGreen.Render(fmt.Sprintf(infoFormat+format, a...)))
	saveLog(fmt.Sprintf(infoFormat+format, a...), "info")
}

func Error(a ...any) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(errorFormat))
	for _, v := range a {
		sb.WriteString(fmt.Sprintf("%v ", v))
	}

	log.Println(color.FgRed.Render(sb.String()))
	saveLog(sb.String(), "error")
}

func Errorf(format string, a ...any) {
	log.Println(color.FgRed.Render(fmt.Sprintf(errorFormat+format, a...)))
	saveLog(fmt.Sprintf(errorFormat+format, a...), "error")
}

func saveLog(msg, level string) {
	date := time.Now().Format(time.DateOnly)
	dateTime := time.Now().Format(time.DateTime)
	dir := "./log"
	fname := fmt.Sprintf("%s/%s-%s.log", dir, date, level)
	if fileutil.PathExists(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}

	f, _ := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer f.Close()

	switch level {
	case "info":
		f.WriteString(dateTime + msg + "\r\n")
	case "error":
		f.WriteString("========================================================================================\r\n")
		f.WriteString(dateTime + msg + "\r\n")
		f.WriteString(string(debug.Stack()) + "\r\n")
	}
}
