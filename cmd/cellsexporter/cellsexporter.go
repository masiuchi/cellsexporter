package main

import "github.com/masiuchi/cellsexporter"

func main() {
	args := cellsexporter.ParseArgs()
	cellsexporter.ProcessFiles(args)
}
