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

	// Part B
	var totalScore int = 0
	var hit int

	// Compute total similarity score
	for _, lval := range left {
		hit = 0
		for _, rval := range right {
			if lval == rval {
				hit++
			}
		}

		// Save result in total similarity score
		totalScore += lval * hit

		// // DEBUG
		// var i int = 0
		// log.Printf("%d) Got lval %d appearing in right %d times.", i, lval, hit)
		// i++
		// log.Println()
	}

	// Print result
	fmt.Printf("Total similarity score: %d\n", totalScore)
}

func toInt(val string) int {
	num, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		log.Fatalf("Error converting value %s to uint32: %v.", val, err)
	}
	return int(num)
}
