package excelexporter

import (
	"fmt"
	"os"
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
		sheet := cellSlice[0]
		cellAxis := cellSlice[1]
		value, err := excel.GetCellValue(sheet, cellAxis)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cols[index] = value
	}
	return cols
}
