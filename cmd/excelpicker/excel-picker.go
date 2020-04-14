package main

import "excelpicker"

func main() {
	args := excelpicker.ParseArgs()
	excelpicker.ProcessFiles(args)
}
