package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	// Result of all mul instructions
	var res int = 0
	mulRE := `mul\((\d{1,3}),(\d{1,3})\)`
	doRE := `do\(\)`
	dontRE := `don't\(\)`

	// Parse inut file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile()
		results := re.FindAllStringSubmatch(line, -1)
		for _, elem := range results {
			//fmt.Printf("%v\n", elem[1:])
			res += mul(elem[1:])
		}
	}

	// Print result
	fmt.Printf("Result: %d\n", res)
}

func mul(str []string) int {
	return toInt(str[0]) * toInt(str[1])
}

func toInt(val string) int {
	num, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		log.Fatalf("Error converting value %s to uint32: %v.", val, err)
	}
	return int(num)
}
