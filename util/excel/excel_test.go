package excel

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
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

	style := func(f *excelize.File) error {
		if err := f.SetColWidth(sheet, "A", string(byte('A'+len(headers))), 25); err != nil {
			return err
		}
		return nil
	}

	if err := Export(sheet, fp, headers, cells, style); err != nil {
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
