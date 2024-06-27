package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"testing"
	"text/template"
)

//go:embed template.go.tpl
var tpl string

func TestTemplate(t *testing.T) {
	type Hello struct {
		Name string
	}
	h := &Hello{Name: "Hello"}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}

	if err = tmpl.Execute(buf, h); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
