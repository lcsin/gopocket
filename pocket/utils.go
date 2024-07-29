package pocket

import (
	"errors"
	"strconv"
	"strings"
)

// Paginate 逻辑分页
func Paginate[T any](list []T, pageNo, pageSize int64) []T {
	if pageNo <= 0 || pageSize <= 0 {
		return []T{}
	}

	total := int64(len(list))
	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage++
	}

	start := (pageNo - 1) * pageSize
	if start >= total {
		return []T{}
	}

	end := start + pageSize
	if end > total {
		end = total
	}

	return list[start:end]
}

func SizeStr2Bytes(str string) (int64, error) {
	const (
		B  int64 = 1
		KB       = 1024 * B
		MB       = 1024 * KB
		GB       = 1024 * MB
		TB       = 1024 * GB
	)

	var (
		size float64
		unit string
	)

	str = strings.TrimSpace(str)
	if strings.HasSuffix(str, "KB") {
		unit = "KB"
		str = strings.TrimSuffix(str, "KB")
	} else if strings.HasSuffix(str, "MB") {
		unit = "MB"
		str = strings.TrimSuffix(str, "MB")
	} else if strings.HasSuffix(str, "GB") {
		unit = "GB"
		str = strings.TrimSuffix(str, "GB")
	} else if strings.HasSuffix(str, "TB") {
		unit = "TB"
		str = strings.TrimSuffix(str, "TB")
	} else {
		return 0, errors.New("invalid size format")
	}

	var err error
	size, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}

	var res int64
	switch unit {
	case "KB":
		res = int64(size) * KB
	case "MB":
		res = int64(size) * MB
	case "GB":
		res = int64(size) * GB
	case "TB":
		res = int64(size) * TB
	}

	return res, nil
}
