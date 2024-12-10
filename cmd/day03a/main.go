package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"regexp"

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

	// Result of all mul instructions
	var res int = 0

	// Parse inut file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		results := re.FindAllStringSubmatch(line, -1)
		for _, elem := range results {
			vPrint("%v\n", elem[1:])
			res += mul(elem[1:])
		}
	}

	// Print result
	fmt.Printf("Result: %d\n", res)
}

func mul(str []string) int {
	return utils.ToInt(str[0]) * utils.ToInt(str[1])
}
