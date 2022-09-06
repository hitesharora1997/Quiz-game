package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "This document contains all the problem statements in format of `question,answer``")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Println("file doesn't exist", *csvFilename)
		os.Exit(1)
	}
	_ = file
}
