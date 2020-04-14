package excelpicker

import "github.com/360EntSecGroup-Skylar/excelize/v2"

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
func (e *Excel) GetCellValue(sheet, cellAxis string) (string, error) {
	if sheet == "" {
		sheet = e.getActiveSheetName()
	}
	return e.file.GetCellValue(sheet, cellAxis)
}

func (e *Excel) getActiveSheetName() string {
	index := e.file.GetActiveSheetIndex()
	return e.file.GetSheetName(index)
}
