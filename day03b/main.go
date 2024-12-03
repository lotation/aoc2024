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

	var input string

	// Parse inut file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		input += line // idk why I need to join everything in one string
	}

	// Result of all mul instructions
	var res int = 0

	re := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)|don't\(\)|do\(\)`)
	results := re.FindAllStringSubmatch(input, -1)

	var dont bool = false // reset dont flag before each iteration

	for _, elem := range results {
		if elem[0] == `don't()` {
			dont = true
		} else if elem[0] == `do()` {
			dont = false
			continue
		}

		if !dont {
			re = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
			numbers := re.FindStringSubmatch(elem[0])
			res += mul(numbers[1:])
			//fmt.Printf("%#v [%d]\n", numbers, res)
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
