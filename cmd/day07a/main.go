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

	var expressions [][]int
	vPrint("Parsing input:\n")

	// Parse inut file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		fields[0] = fields[0][:len(fields[0])-1]
		var intExpr []int
		for _, expr := range fields {
			intExpr = append(intExpr, utils.ToInt(expr))
		}

		vPrint("result=%d - operands='%v'\n", intExpr[0], intExpr[1:])

		expressions = append(expressions, intExpr)
	}
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input file %s: %v", inputfile, err)
	}
	vPrint("\n")

	var count int = 0

	for _, expr := range expressions {
		if checkExpr(expr) {
			count += expr[0]
		}
	}

	// Print result
	fmt.Printf("Result: %d\n", count)
}

func checkExpr(expr []int) bool {
	res := expr[0]
	ops := expr[1:]
	for i := 0; i < len(expr)-1; i++ {
		expr[i] = expr[i+1]
	}
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}
