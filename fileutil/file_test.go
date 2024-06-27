package fileutil

import (
	"fmt"
	"os"
	"testing"
)

func TestPathExists(t *testing.T) {
	fmt.Println(PathExists("../fileutil/file.go"))
}

func TestCopyFile(t *testing.T) {
	_, err := os.Create("test1.txt")
	if err != nil {
		panic(err)
	}
	if err = CopyFile("test1.txt", "test2.txt"); err != nil {
		panic(err)
	}
}

func TestCopyDir(t *testing.T) {
	if err := CopyDir("dir1", "test/dir1", nil); err != nil {
		panic(err)
	}
}

func TestFileMD5(t *testing.T) {
	hash, err := FileMD5("file.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(hash)
}
