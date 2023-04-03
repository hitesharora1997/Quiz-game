package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// flag here for the new string
	csvFilename := flag.String("csv", "problems.csv", "This document contains all the problem statements in format of `question,answer`")
	// parsing all the flags
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("file doesn't exist %v", *csvFilename))
	}

	read := csv.NewReader(file)
	// Read all data from the file
	lines, err := read.ReadAll()
	// Check the error
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	// ParseLines is the function which take the Input data what we got from read.ReadAll()
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

// exit function
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
