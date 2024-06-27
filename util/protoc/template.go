package main

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed template.go.tpl
var tmpl string

type ginSrv struct {
	Name string
}

func (g ginSrv) execute() string {
	buf := new(bytes.Buffer)

	t, err := template.New("http").Parse(strings.TrimSpace(tmpl))
	if err != nil {
		panic(err)
	}
	if err = t.Execute(buf, g); err != nil {
		panic(err)
	}

	return buf.String()
}
