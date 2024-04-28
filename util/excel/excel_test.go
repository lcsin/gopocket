package excel

import (
	"fmt"
	"testing"
)

const (
	sheet = "test"
	fp    = "./test.xlsx"
)

func TestExport(t *testing.T) {
	headers := []string{"班级", "姓名", "年龄"}
	cells := make(map[string][]interface{})
	cells["班级"] = []interface{}{"计算机1班", "计算机1班", "计算机1班"}
	cells["姓名"] = []interface{}{"张三", "李四", "王五"}
	cells["年龄"] = []interface{}{18, 17, 19}

	if err := Export(sheet, fp, headers, cells); err != nil {
		panic(err)
	}
}

func TestImport(t *testing.T) {
	values, err := Import(fp, sheet)
	if err != nil {
		panic(err)
	}

	for i, v := range values {
		fmt.Println(i, v)
	}
}
