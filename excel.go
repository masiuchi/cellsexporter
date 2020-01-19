package cellsexporter

import "github.com/360EntSecGroup-Skylar/excelize"

// Excel ...
type Excel struct {
	file *excelize.File
}

// LoadExcel ...
func LoadExcel(filename string) (*Excel, error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	return &Excel{file: file}, nil
}

// GetCellValue ...
func (e *Excel) GetCellValue(sheet, axis string) string {
	return e.file.GetCellValue(sheet, axis)
}
