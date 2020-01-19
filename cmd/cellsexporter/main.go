package main

import "cellsexporter"

func main() {
	args := cellsexporter.ParseArgs()
	cellsexporter.ProcessFiles(args)
}
