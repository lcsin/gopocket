package fileutil

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
)

// ZipFiles 压缩文件
// 支持不同路径下的多个文件
func ZipFiles(src []string, dst string) error {
	if !PathExists(dst) {
		if err := os.MkdirAll(path.Dir(dst), os.ModePerm); err != nil {
			return err
		}
	}

	zipFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range src {
		err = addFile2ZipWithDir(zipWriter, file)
		if err != nil {
			return err
		}
	}

	return nil
}

// ZipDir 压缩目录
func ZipDir(dir, dst string) error {
	// Create a new zip archive.
	zipFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the source directory.
	if err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a header for the file or directory.
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Use the relative path as the header name.
		header.Name, err = filepath.Rel(filepath.Dir(dir), path)
		if err != nil {
			return err
		}

		// Adjust header attributes.
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		// Create a writer for the file in the zip archive.
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If it's a file, copy the file content to the zip writer.
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// 添加文件到压缩包不包含原始目录结构
func addFile2ZipWithoutDir(zipWriter *zip.Writer, fp string) error {
	file2Zip, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file2Zip.Close()

	fileInfo, err := file2Zip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(fp)

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file2Zip)
	if err != nil {
		return err
	}

	return nil
}

// 添加文件到压缩包且包含原始结构目录
func addFile2ZipWithDir(zipWriter *zip.Writer, fp string) error {
	file2Zip, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file2Zip.Close()

	writer, err := zipWriter.Create(fp)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file2Zip)
	if err != nil {
		return err
	}

	return nil
}
