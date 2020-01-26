package main

import "excelexporter"

func main() {
	args := excelexporter.ParseArgs()
	excelexporter.ProcessFiles(args)
}
