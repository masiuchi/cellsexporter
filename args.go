package excelpicker

import (
	"flag"
	"fmt"
	"os"
)

// Args ...
type Args struct {
	Files      []string
	Cells      []string
	ExportType string
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
	var exportType string
	flag.Var(&cells, "cell", "sheetName/axis")
	flag.Var(&files, "file", "filename")
	flag.StringVar(&exportType, "type", "json", "json or csv")
	flag.Parse()

	if len(files) == 0 || len(cells) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	return &Args{Files: files, Cells: cells, ExportType: exportType}
}

// GetHeaders ...
func (a *Args) GetHeaders() []string {
	return a.Cells
}
