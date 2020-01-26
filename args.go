package excelexporter

import (
	"flag"
	"fmt"
)

// Args ...
type Args struct {
	Files []string
	Cells []string
}

type strslice []string

func (s *strslice) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *strslice) Set(v string) error {
	*s = append(*s, v)
	return nil
}

// ParseArgs ...
func ParseArgs() *Args {
	var cells, files strslice
	flag.Var(&cells, "cell", "sheetName!axis")
	flag.Var(&files, "file", "file")
	flag.Parse()
	return &Args{Files: files, Cells: cells}
}

// GetHeaders ...
func (a *Args) GetHeaders() []string {
	return a.Cells
}
