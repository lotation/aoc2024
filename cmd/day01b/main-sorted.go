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

	// Part B
	// Start with sorting the right list so
	// we can avoid scanning the whole slice
	sort.Ints(left)
	sort.Ints(right)

	var totalScore int = 0
	var hit int
	var rindx int

	// Compute total similarity score
	for _, lval := range left {
		rindx = 0
		hit = 0

		// Look for first occurence of lval in right list
		for right[rindx] < lval {
			rindx++
		}

		// Now count how many equal occurences
		for right[rindx] == lval {
			hit++
			rindx++
		}

		// Save result in total similarity score
		totalScore += lval * hit

		if verbose {
			i := 0
			vPrint("%d) Got lval %d appearing in right %d times.\n", i, lval, hit)
			i++
		}
	}

	// Print result
	fmt.Printf("Total similarity score: %d\n", totalScore)
}
