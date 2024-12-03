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
		} else {
			// try to remove one bad level at a time
			for e := 0; e < len(report); e++ {
				tmpreport := remove(report, e)
				// fmt.Printf("Retrying report %v without element %d\n", report, report[e])
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

	// fmt.Printf("Working with report %v\n", report)

	// Check if ascending
	for n := 0; n < len(report)-1; {
		found = false // Flag to check if a valid increment was found
		for inc := 1; inc <= maxIncr; {
			// fmt.Printf("%d) Checking asc %d + %d == %d ?", n, report[n], inc, report[n+1])
			if report[n]+inc == report[n+1] {
				n++
				found = true // valid increment
				// fmt.Println("  => yes")
				break
			} else {
				inc++
				// fmt.Println("  => no")
			}
		}
		if !found {
			// no valid increment was found
			break
		}
	}

	if found {
		// fmt.Printf("=> Report %v is safe\n\n", report)
		return true
	}

	// Check if descending
	for n := 0; n < len(report)-1; {
		found = false // flag to check if a valid increment was found
		for dec := 1; dec <= maxDecr; {
			// fmt.Printf("%d) Checking desc %d - %d == %d ?", n, report[n], dec, report[n+1])
			if report[n]-dec == report[n+1] {
				n++
				found = true // valid decrement
				// fmt.Println("  => yes")
				break
			} else {
				dec++
				// fmt.Println("  => no")
			}
		}
		if !found {
			// no valid increment was found
			break
		}
	}

	if found {
		// fmt.Printf("=> Report %v is safe\n\n", report)
		return true
	}

	// fmt.Printf("=> Report %v is NOT safe\n\n", report)
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

func toInt(val string) int {
	num, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		log.Fatalf("Error converting value %s to uint32: %v.", val, err)
	}
	return int(num)
}
