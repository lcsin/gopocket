package fileutil

import (
	"os"
	"path/filepath"
)

// IsExists 判断文件、目录是否存在
func IsExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// CopyFile 拷贝文件
func CopyFile(src, dst string) error {
	srcinfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	buf, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, buf, srcinfo.Mode())
}

// CopyDir 拷贝目录
func CopyDir(src, dst string, ignores []string) error {
	srcinfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, srcinfo.Mode())
	if err != nil {
		return err
	}

	fds, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	hasIgnores := func(name string, ignores []string) bool {
		for _, v := range ignores {
			if v == name {
				return true
			}
		}
		return false
	}

	for _, fd := range fds {
		if hasIgnores(fd.Name(), ignores) {
			continue
		}
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())
		var e error
		if fd.IsDir() {
			e = CopyDir(srcfp, dstfp, ignores)
		} else {
			e = CopyFile(srcfp, dstfp)
		}
		if e != nil {
			return e
		}
	}
	return nil
}
