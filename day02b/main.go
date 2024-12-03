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
			for _, e := range report {
				report = remove(report, e)
				if isSafe(report) {
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

	fmt.Print("Working with report ")
	fmt.Println(report)

	// Check if ascending
	for n := 0; n < len(report)-1; {
		found = false // Flag to check if a valid increment was found
		for inc := 1; inc <= maxIncr; {
			fmt.Printf("%d) Checking asc %d + %d == %d ?", n, report[n], inc, report[n+1])
			if report[n]+inc == report[n+1] {
				n++
				found = true // valid increment
				fmt.Println("  => yes")
				break
			} else {
				inc++
				fmt.Println("  => no")
			}
		}
		if !found {
			// no valid increment was found
			// and more than 1 bad level,
			// break out of the outer loop
			break
		}
	}

	if found {
		fmt.Print("=> Report ")
		fmt.Print(report)
		fmt.Printf(" is safe\n\n")
		return true
	}

	// Check if descending
	for n := 0; n < len(report)-1; {
		found = false // flag to check if a valid increment was found
		for dec := 1; dec <= maxDecr; {
			fmt.Printf("%d) Checking desc %d - %d == %d ?", n, report[n], dec, report[n+1])
			if report[n]-dec == report[n+1] {
				n++
				found = true // valid decrement
				fmt.Println("  => yes")
				break
			} else {
				dec++
				fmt.Println("  => no")
			}
		}
		if !found {
			// no valid increment was found
			// and more than 1 bad level,
			// break out of the outer loop
			break
		}
	}

	if found {
		fmt.Print("=> Report ")
		fmt.Print(report)
		fmt.Printf(" is safe\n\n")
		return true
	}

	fmt.Print("=> Report ")
	fmt.Print(report)
	fmt.Printf(" is NOT safe\n\n")
	return false
}

// controllo out of bounds index
func removeIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func toInt(val string) int {
	num, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		log.Fatalf("Error converting value %s to uint32: %v.", val, err)
	}
	return int(num)
}
