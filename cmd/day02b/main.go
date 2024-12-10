package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"strings"

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
			report = append(report, utils.ToInt(f))
		}

		// Check for report safety
		if isSafe(report) {
			safeReports++
		} else {
			// try to remove one bad level at a time
			for e := 0; e < len(report); e++ {
				tmpreport := remove(report, e)
				vPrint("Retrying report %v without element %d\n", report, report[e])
				if isSafe(tmpreport) {
					safeReports++
					break
				}
			}
		}
	}

	// Print result
	fmt.Printf("Total safe reports: %d\n", safeReports)
}

func isSafe(report []int) bool {
	maxIncr := 3
	maxDecr := 3
	found := false

	vPrint("Working with report %v\n", report)

	// Check if ascending
	for n := 0; n < len(report)-1; {
		found = false // Flag to check if a valid increment was found
		for inc := 1; inc <= maxIncr; {
			vPrint("%d) Checking asc %d + %d == %d ?", n, report[n], inc, report[n+1])
			if report[n]+inc == report[n+1] {
				n++
				found = true // valid increment
				vPrint("  => yes")
				break
			} else {
				inc++
				vPrint("  => no")
			}
		}
		if !found {
			// no valid increment was found
			break
		}
	}

	if found {
		vPrint("=> Report %v is safe\n\n", report)
		return true
	}

	// Check if descending
	for n := 0; n < len(report)-1; {
		found = false // flag to check if a valid increment was found
		for dec := 1; dec <= maxDecr; {
			vPrint("%d) Checking desc %d - %d == %d ?", n, report[n], dec, report[n+1])
			if report[n]-dec == report[n+1] {
				n++
				found = true // valid decrement
				vPrint("  => yes")
				break
			} else {
				dec++
				vPrint("  => no")
			}
		}
		if !found {
			// no valid increment was found
			break
		}
	}

	if found {
		vPrint("=> Report %v is safe\n\n", report)
		return true
	}

	vPrint("=> Report %v is NOT safe\n\n", report)
	return false
}

// remove element in slice with index performing bound checking
func remove(slice []int, index int) []int {
	if index >= 0 && index < len(slice) {
		ret := make([]int, 0, len(slice)-1)
		ret = append(ret, slice[:index]...)
		return append(ret, slice[index+1:]...)
	} else {
		return slice
	}
}
