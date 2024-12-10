package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"sort"
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
	flag.StringVar(&inputfile, "input", "input-01a.txt", "path to file containing current day input")
	flag.BoolVar(&verbose, "verbose", true, "enable verbose logging")
	flag.Parse()

	// Open input file
	fp := utils.Fopen(inputfile)
	defer func() {
		if err := fp.Close(); err != nil {
			log.Panic(err)
		}
	}()

	// Parse inut file
	var left, right []int

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Convert values to uint32 and save
		left = append(left, utils.ToInt(fields[0]))
		right = append(right, utils.ToInt(fields[1]))
	}

	if len(left) != len(right) {
		log.Fatalf("Malformed input: Location IDs list are not the same length: %d != %d.", len(left), len(right))
	}

	// Start with sorting the two lists
	sort.Ints(left)
	sort.Ints(right)

	// Compute and sum the  difference
	var diff int
	for indx := range len(left) {
		diff += max(left[indx]-right[indx], right[indx]-left[indx])
	}

	// Print result
	fmt.Printf("Difference: %d\n", diff)
}
