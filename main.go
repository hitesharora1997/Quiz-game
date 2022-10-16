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
	// Read all data from the file
	lines, err := read.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)
	fmt.Println(problems)

	for i, p := range problems {
		fmt.Printf("Problem #%d:%s= \n", i+1, p.q)
	}
}

func parseLines(lines [][]string) []problem {
	// Making a slice of problem which lines length
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
