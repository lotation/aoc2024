package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"

	"github.com/lotation/aoc2024/internal/utils"
)

var verbose bool = true

func vPrint(format string, args ...interface{}) {
	if verbose {
		fmt.Printf(format, args...)
	}
}

func main() {
	var inputfile string
	flag.StringVar(&inputfile, "input", "input.txt", "path to file containing current day input")
	flag.BoolVar(&verbose, "verbose", true, "enable verbose logging")
	flag.Parse()

	// Open input file
	fp := utils.Fopen(inputfile)
	defer func() {
		if err := fp.Close(); err != nil {
			log.Panic(err)
		}
	}()

	var slice []string

	// Parse inut file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		slice = append(slice, line)
	}
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input file %s: %v", inputfile, err)
	}

	var count int = 0

	// Print result
	fmt.Printf("Result: %d\n", count)
}
