package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const word string = "XMAS"

var verbose = true

func main() {
	var inputfile string
	flag.StringVar(&inputfile, "input", "input.txt", "path to file containing current day input")
	flag.BoolVar(&verbose, "verbose", true, "enable verbose logging")
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

	// Slice of functions with all the different checks
	checks := []func(string, []string, int, int) bool{
		checkRight,
		checkLeft,
		checkDown,
		checkUp,
		checkRightDown,
		checkRightUp,
		checkLeftDown,
		checkLeftUp,
	}

	// Result
	var count int = 0

	// Start looking for word XMAS
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			//fmt.Printf("%c ", table[i][j])

			// Check if start of the word
			if table[i][j] == 'X' { // word[0]
				// Try all checks
				for _, check := range checks {
					if check(word, table, i, j) {
						count++ // word found in table
					}
				}
				fmt.Println()
			}
		}
		//fmt.Println()
	}
	// fmt.Println()

	// Print result
	fmt.Printf("Result: %d\n", count)
}

func checkRight(word string, table []string, i int, j int) bool {
	if i < 0 || i >= len(table) || j < 0 || j+len(word) > len(table[i]) {
		return false
	}
	s := rightSubstring(table, i, j) // table[i][j : j+len(word)]
	logAction("Right", i, j, s, s == word)
	return s == word
}

func checkLeft(word string, table []string, i int, j int) bool {
	if i < 0 || i >= len(table) || j < len(word)-1 || j >= len(table[i]) {
		return false
	}
	s := leftSubstring(table, i, j) // table[i][j-len(word)+1 : j+1]
	logAction("Left", i, j, s, s == word)
	return s == word //rev(word)
}

func checkDown(word string, table []string, i int, j int) bool {
	if i < 0 || i+len(word) > len(table) || j < 0 || j >= len(table[i]) {
		return false
	}
	s := downSubstring(table, i, j) // table[i : i+len(word)][j]
	logAction("Down", i, j, s, s == word)
	return s == word
}

func checkUp(word string, table []string, i int, j int) bool {
	if i < len(word)-1 || i >= len(table) || j < 0 || j >= len(table[i]) {
		return false
	}
	s := upSubstring(table, i, j) // table[i-len(word)+1 : i][j]
	logAction("Up", i, j, s, s == word)
	return s == word
}

func checkRightDown(word string, table []string, i int, j int) bool {
	if i < 0 || i+len(word) >= len(table) || j < 0 || j+len(word) >= len(table[i]) {
		return false
	}
	s := rightDownSubstring(table, i, j) // table[i : i+len(word)][j : j+len(word)]
	logAction("Right-Down", i, j, s, s == word)
	return s == word
}

func checkRightUp(word string, table []string, i int, j int) bool {
	if i < len(word)-1 || i >= len(table) || j < 0 || j+len(word) >= len(table[i]) {
		return false
	}
	s := rightUpSubstring(table, i, j) // table[i-len(word)+1 : i][j : j+len(word)]
	logAction("Right-Up", i, j, s, s == word)
	return s == word
}

func checkLeftDown(word string, table []string, i int, j int) bool {
	if i < 0 || i+len(word) >= len(table) || j < len(word)-1 || j >= len(table[i]) {
		return false
	}
	s := leftDownSubstring(table, i, j) // table[i : i+len(word)][j-len(word)+1 : j]
	logAction("Left-Down", i, j, s, s == word)
	return s == word
}

func checkLeftUp(word string, table []string, i int, j int) bool {
	if i < len(word)-1 || i >= len(table) || j < len(word)-1 || j >= len(table[i]) {
		return false
	}
	s := leftUpSubstring(table, i, j) // table[i-len(word)+1 : i][j-len(word)+1 : j]
	logAction("Left-Up", i, j, s, s == word)
	return s == word
}

// func rev(str string) string {
// 	runes := []rune(str)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

func rightSubstring(table []string, i, j int) string {
	var result []rune
	for k := 0; k < len(word); k++ {
		if j < len(table[i]) {
			result = append(result, rune(table[i][j]))
			j++
		}
	}
	return string(result)
}

func leftSubstring(table []string, i, j int) string {
	var result []rune
	for k := 0; k < len(word); k++ {
		if j >= 0 {
			result = append(result, rune(table[i][j]))
			j--
		}
	}
	return string(result)
}

func downSubstring(table []string, i, j int) string {
	var result []rune
	for k := 0; k < len(word); k++ {
		if i < len(table) {
			result = append(result, rune(table[i][j]))
			i++
		}
	}
	return string(result)
}

func upSubstring(table []string, i, j int) string {
	var result []rune
	for k := 0; k < len(word); k++ {
		if i >= 0 {
			result = append(result, rune(table[i][j]))
			i--
		}
	}
	return string(result)
}

func rightDownSubstring(table []string, i, j int) string {
	// var result []rune
	// for i < len(table) && j < len(table[0]) {
	// 	result = append(result, rune(table[i][j]))
	// 	i++
	// 	j++
	// }
	// return string(result)
	var result []rune
	for k := 0; k < len(word); k++ {
		if i < len(table) && j < len(table[0]) {
			result = append(result, rune(table[i][j]))
			i++
			j++
		}
	}
	return string(result)
}

func rightUpSubstring(table []string, i, j int) string {
	var result []rune
	for k := 0; k < len(word); k++ {
		if i >= 0 && j < len(table[0]) {
			result = append(result, rune(table[i][j]))
			i--
			j++
		}
	}
	return string(result)
}

func leftDownSubstring(table []string, i, j int) string {
	var result []rune
	for k := 0; k < len(word); k++ {
		if i < len(table) && j >= 0 {
			result = append(result, rune(table[i][j]))
			i++
			j--
		}
	}
	return string(result)
}

func leftUpSubstring(table []string, i, j int) string {
	// var result []rune
	// for i >= 0 && j >= 0 {
	// 	result = append(result, rune(table[i][j]))
	// 	i--
	// 	j--
	// }
	// return string(result)
	var result []rune
	for k := 0; k < len(word); k++ {
		if i >= 0 && j >= 0 {
			result = append(result, rune(table[i][j]))
			i--
			j--
		}
	}
	return string(result)
}

func logAction(action string, i, j int, substr string, cond bool) {
	if verbose {
		fmt.Printf("[%d,%d] %-10s  %s  {%t}\n", i, j, action, substr, cond)
	}
}
