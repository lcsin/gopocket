package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// Import excel导入
func Import(fp, sheet string) ([][]string, error) {
	f, err := excelize.OpenFile(fp)
	if err != nil {
		return nil, err
	}

	return f.GetRows(sheet)
}

// Export excel导出
// 字段数量只支持A~Z
func Export(sheet, fp string, headers []string, cells map[string][]interface{}, style func(f *excelize.File) error) error {
	f := excelize.NewFile()
	defer f.Close()

	if _, err := f.NewSheet(sheet); err != nil {
		return err
	}

	// 设置样式
	if style != nil {
		if err := style(f); err != nil {
			return err
		}
	}

	HMap := make(map[string]string, len(headers))
	for i, header := range headers {
		HMap[header] = string(byte('A' + i))
	}

	CMap := make(map[string]interface{}, len(cells)+len(headers))
	// 表头
	for _, header := range headers {
		CMap[fmt.Sprintf("%v%v", HMap[header], 1)] = header
	}
	// 内容
	for header, values := range cells {
		for i, v := range values {
			CMap[fmt.Sprintf("%v%v", HMap[header], i+2)] = v
		}
	}
	// 写入
	for k, v := range CMap {
		if err := f.SetCellValue(sheet, k, v); err != nil {
			return err
		}
	}
	_ = f.DeleteSheet("Sheet1")

	return f.SaveAs(fp)
}

// StreamExport 流式导出
func StreamExport(sheet, fp string, headers []interface{}, rows [][]interface{}) error {
	f := excelize.NewFile()
	defer f.Close()

	if _, err := f.NewSheet(sheet); err != nil {
		return err
	}
	sw, err := f.NewStreamWriter(sheet)
	if err != nil {
		return err
	}

	// 表头
	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		return err
	}
	if err = sw.SetRow(cell, headers); err != nil {
		return err
	}
	// 内容
	for i, row := range rows {
		cell, _ = excelize.CoordinatesToCellName(1, i+2)
		if err = sw.SetRow(cell, row); err != nil {
			return err
		}
	}
	if err = sw.Flush(); err != nil {
		return err
	}

	_ = f.DeleteSheet("Sheet1")
	return f.SaveAs(fp)
}

// StreamExportMultiSheet 多sheet流式导出
func StreamExportMultiSheet(sheet, fp string, headers []interface{}, rows [][]interface{}, size int) error {
	f := excelize.NewFile()
	defer f.Close()

	var num int
	if len(rows)%size == 0 {
		num = len(rows) / size
	} else {
		num = len(rows)/size + 1
	}

	for i := 1; i <= num; i++ {
		sheetName := fmt.Sprintf("%v%v", sheet, i)
		if _, err := f.NewSheet(sheetName); err != nil {
			return err
		}
		sw, err := f.NewStreamWriter(sheetName)
		if err != nil {
			return err
		}

		// 表头
		cell, err := excelize.CoordinatesToCellName(1, 1)
		if err != nil {
			return err
		}
		if err = sw.SetRow(cell, headers); err != nil {
			return err
		}
		// 内容
		var data [][]interface{}
		if i == num {
			data = rows[(i-1)*size:]
		} else {
			data = rows[(i-1)*size : i*size]
		}

		for j, row := range data {
			cell, _ = excelize.CoordinatesToCellName(1, j+2)
			if err = sw.SetRow(cell, row); err != nil {
				return err
			}
		}
		if err = sw.Flush(); err != nil {
			return err
		}
	}

	_ = f.DeleteSheet("Sheet1")
	return f.SaveAs(fp)
}
