package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var inputfile string
	flag.StringVar(&inputfile, "input", "input-01a.txt", "path to file containing current day input")
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

	// Parse inut file
	var left, right []int

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Convert values to uint32 and save
		left = append(left, toInt(fields[0]))
		right = append(right, toInt(fields[1]))
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

func toInt(val string) int {
	num, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		log.Fatalf("Error converting value %s to uint32: %v.", val, err)
	}
	return int(num)
}
