package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const word string = "XMAS"

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

	var table []string

	// Parse inut file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, line)
	}
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file %s: %v", inputfile, err)
		return
	}

	// Start looking for word XMAS
	for i, row := range table {
		for j, char := range row {
			//fmt.Printf("%c ", table[i][j])

			// check if start of word
			if char == word[0] { // X
				checkHorizontalForward(table, i, j)
				checkHorizontalBackward(table, i, j)
				checkVericalForward(table, i, j)
				checkVerticalBackward(table, i, j)
				checkDiagonalForward(table, i, j)
				checkDiagonalForward(table, i, j)
			}
		}
		//fmt.Println()
	}
	fmt.Println()

	// Result
	var res int = 0

	// Print result
	fmt.Printf("Result: %d\n", res)
}

func checkHorizontalForward(table []string, i int, j int) bool {
	// if table[i][j:j+4] == []rune(word) {
	// 	return true
	// }
	return false
}

func checkHorizontalBackward(table []string, i int, j int) bool {
	return false
}

func checkVericalForward(table []string, i int, j int) bool {
	return false
}

func checkVerticalBackward(table []string, i int, j int) bool {
	return false
}

func checkDiagonalForward(table []string, i int, j int) bool {
	return false
}

func checkDiagonalBackward(table []string, i int, j int) bool {
	return false
}
