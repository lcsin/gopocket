package util

import (
	"os"
	"testing"
)

func TestCopyFile(t *testing.T) {
	_, _ = os.Create("test1.txt")
	if err := CopyFile("test1.txt", "test2.txt"); err != nil {
		panic(err)
	}
}

func TestCopyDir(t *testing.T) {
	if err := CopyDir("dir1", "test/dir1", nil); err != nil {
		panic(err)
	}
}
