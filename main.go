package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "This document contains all the problem statements in format of `question,answer``")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("file doesn't exist %v", *csvFilename))
	}
	read := csv.NewReader(file)
	lines, err := read.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	fmt.Println(lines)
}

type problem struct {
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
