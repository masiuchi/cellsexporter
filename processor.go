package excelpicker

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
		data[dataIndex] = pickCellValues(excel, args.Cells)
		dataIndex = dataIndex + 1
	}
	PrintAsCsv(args.GetHeaders(), data)
}

func pickCellValues(excel *Excel, cells []string) []string {
	cols := make([]string, len(cells))
	for index, cell := range cells {
		sheet, cellAxis := parseCellArgument(cell)
		value, err := excel.GetCellValue(sheet, cellAxis)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cols[index] = value
	}
	return cols
}

func parseCellArgument(cell string) (string, string) {
	cellSlice := strings.Split(cell, "/")
	var sheet, cellAxis string
	if len(cellSlice) >= 2 {
		sheet = cellSlice[0]
		cellAxis = cellSlice[1]
	} else {
		sheet = ""
		cellAxis = cellSlice[0]
	}
	return sheet, cellAxis
}
