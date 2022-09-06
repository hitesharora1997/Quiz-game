package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "problems.csv", "This document contains all the problem statements in format of `question,answer``")
	flag.Parse()
	_ = csvFilename
}
