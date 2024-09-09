package fileutil

import (
	"testing"
)

const dst = "./archive.zip"

func TestZipFiles(t *testing.T) {
	files := []string{"../ginlayout/main.go", "../ginlayout/wire.go"}
	if err := ZipFiles(files, dst); err != nil {
		panic(err)
	}
}

func TestZipDir(t *testing.T) {
	dir := "../ginlayout"
	if err := ZipDir(dir, dst); err != nil {
		panic(err)
	}
}
