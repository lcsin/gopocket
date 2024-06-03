package excel

import (
	"fmt"
	"testing"
	"time"
)

const (
	sheet = "test"
	fp    = "./test.xlsx"
	count = 1000000
)

func TestStreamExportMultiSheet(t *testing.T) {
	headers := []interface{}{"班级", "姓名", "年龄"}
	rows := make([][]interface{}, 0, count+2000)

	for i := 0; i < count+2000; i++ {
		rows = append(rows, []interface{}{
			"计算机1班", "张三", 18,
		})
	}

	start := time.Now()
	if err := StreamExportMultiSheet(sheet, fp, headers, rows, count); err != nil {
		panic(err)
	}
	fmt.Println("StreamExportMultiSheet: ", time.Now().Sub(start))
}

func TestStreamExport(t *testing.T) {
	headers := []interface{}{"班级", "姓名", "年龄"}
	rows := make([][]interface{}, 0, count)

	for i := 0; i < count; i++ {
		rows = append(rows, []interface{}{
			"计算机1班", "张三", 18,
		})
	}

	start := time.Now()
	if err := StreamExport(sheet, fp, headers, rows); err != nil {
		panic(err)
	}
	fmt.Println("StreamExport: ", time.Now().Sub(start))
}

func TestExport(t *testing.T) {
	headers := []string{"班级", "姓名", "年龄"}
	cells := make(map[string][]interface{}, count)
	for i := 0; i < count; i++ {
		cells["班级"] = append(cells["班级"], "计算机1班")
		cells["姓名"] = append(cells["班级"], "张三")
		cells["年龄"] = append(cells["班级"], 18)
	}

	start := time.Now()
	if err := Export(sheet, fp, headers, cells, nil); err != nil {
		panic(err)
	}
	fmt.Println("Export: ", time.Now().Sub(start))
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
