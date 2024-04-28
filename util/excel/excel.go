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
func Export(sheet, fp string, headers []string, cells map[string][]interface{}) error {
	f := excelize.NewFile()
	defer f.Close()

	if _, err := f.NewSheet(sheet); err != nil {
		return err
	}
	if err := f.SetColWidth(sheet, "A", string(byte('A'+len(headers))), 25); err != nil {
		return err
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
