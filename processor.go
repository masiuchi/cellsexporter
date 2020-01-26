package excelexporter

import (
	"strings"
)

// ProcessFiles ...
func ProcessFiles(args *Args) {
	data := make([][]string, len(args.Files))
	dataIndex := 0
	for _, file := range args.Files {
		excel, err := LoadExcel(file)
		if err != nil {
			panic(err)
		}
		data[dataIndex] = processExcel(excel, args.Cells)
		dataIndex = dataIndex + 1
	}
	PrintAsCsv(args.GetHeaders(), data)
}

func processExcel(excel *Excel, cells []string) []string {
	cols := make([]string, len(cells))
	for index, cell := range cells {
		cellSlice := strings.Split(cell, "!")
		value, err := excel.GetCellValue(cellSlice[0], cellSlice[1])
		if err != nil {
			panic(err)
		}
		cols[index] = value
	}
	return cols
}
