package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputfile string
	flag.StringVar(&inputfile, "input", "input.txt", "path to file containing current day input")
	flag.Parse()

	// Open input file
	fp, err := os.Open(inputfile)
	if err != nil {
		log.Fatalf("Error opening input file %s: %v", inputfile, err)
		return
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Panic(err)
		}
	}()

	// Report represent a file line
	var report []int
	// Safe reports
	var safeReports int = 0

	// Parse inut file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		report = nil
		line := scanner.Text()
		fields := strings.Fields(line)

		for _, f := range fields {
			report = append(report, toInt(f))
		}

		// Check for report safety
		if isSafe(report) {
			safeReports++
		}
	}

	// Print result
	fmt.Printf("Total safe reports: %d\n", safeReports)
}

func isSafe(report []int) bool {
	maxIncr := 3
	maxDecr := 3
	found := false

	// Check if ascending
	for n := 0; n < len(report)-1; {
		found = false // Flag to check if a valid increment was found
		for inc := 1; inc <= maxIncr; {
			if report[n]+inc == report[n+1] {
				n++
				found = true // Set flag to indicate a valid increment was found
				break
			} else {
				inc++
			}
		}
		if !found {
			// If no valid increment was found, break out of the outer loop
			break
		}
	}

	if found {
		return true
	}

	// Check if descending
	for n := 0; n < len(report)-1; {
		found = false // Flag to check if a valid increment was found
		for dec := 1; dec <= maxDecr; {
			if report[n]-dec == report[n+1] {
				n++
				found = true // Set flag to indicate a valid decrement was found
				break
			} else {
				dec++
			}
		}
		if !found {
			// If no valid increment was found, break out of the outer loop
			break
		}
	}

	if found {
		return true
	}

	return false
}

func toInt(val string) int {
	num, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		log.Fatalf("Error converting value %s to uint32: %v.", val, err)
	}
	return int(num)
}
