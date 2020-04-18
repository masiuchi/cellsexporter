package excelpicker

import (
	"fmt"
	"os"
	"strings"
)

// ProcessFiles ...
func ProcessFiles(args *Args) {
	lines := make([][]string, len(args.Files))
	for i, file := range args.Files {
		excel, err := LoadExcel(file)
		if err != nil {
			panic(err)
		}
		line := pickCellValues(excel, args.Cells)
		lines[i] = append([]string{file}, line...)
	}
	headers := append([]string{"_filename"}, args.GetHeaders()...)
	// if err := PrintAsCsv(headers, lines); err != nil {
	// 	panic(err)
	// }
	if err := PrintAsJSON(headers, lines); err != nil {
		panic(err)
	}
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
