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
	s := table[i][j : j+len(word)]
	fmt.Printf("%s [%d,%d] %s  {%t}\n", "Horizontal-Forward", i, j, s, s == word)
	return s == word
}

func checkLeft(word string, table []string, i int, j int) bool {
	if i < 0 || i >= len(table) || j < len(word)-1 || j >= len(table[i]) {
		return false
	}
	s := table[i][j-len(word)+1 : j+1]
	fmt.Printf("%s [%d,%d] %s  {%t}\n", "Horizontal-Backward", i, j, s, s == rev(word))
	return s == rev(word)
}

func checkDown(word string, table []string, i int, j int) bool {
	if j < 0 || j >= len(table[0]) || i < 0 || i+len(word) > len(table) {
		return false
	}
	s := downSubstring(table, i, j) // table[i : i+len(word)][j]
	fmt.Printf("%s [%d,%d] %s  {%t}\n", "Vertical-Forward", i, j, s, s == word)
	return s == word
}

func checkUp(word string, table []string, i int, j int) bool {
	if j < 0 || j >= len(table[0]) || i < len(word)-1 || i >= len(table) {
		return false
	}
	s := upSubstring(table, i, j) // table[i-len(word)+1 : i][j]
	fmt.Printf("%s [%d,%d] %s  {%t}\n", "Verical-Backward", i, j, s, s == word)
	return s == word
}

func checkRightDown(word string, table []string, i int, j int) bool {
	if i < 0 || i >= len(table) || j < 0 || j >= len(table[0]) {
		return false
	}
	s := rightDownSubstring(table, i, j) // table[i : i+len(word)][j : j+len(word)]
	fmt.Printf("%s [%d,%d] %s  {%t}\n", "Diagonal-Forward", i, j, s, s == word)
	return s == word
}

// TODO
// write similar to checkRightDown
func checkRightUp(word string, table []string, i int, j int) bool {
	return false
}

// TODO
// write similar to checkRightDown
func checkLeftDown(word string, table []string, i int, j int) bool {
	return false
}

func checkLeftUp(word string, table []string, i int, j int) bool {
	if i < 0 || i >= len(table) || j < 0 || j >= len(table[0]) {
		return false
	}
	s := leftUpSubstring(table, i, j) // table[i-len(word)+1 : i][j-len(word)+1 : j]
	fmt.Printf("%s [%d,%d] %s  {%t}\n", "Diagonal-Backward", i, j, s, s == word)
	return s == word
}

func rev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func downSubstring(table []string, i, j int) string {
	var result []rune
	for ; i < len(word); i++ {
		result = append(result, rune(table[i][j]))
	}
	return string(result)
}

func upSubstring(table []string, i, j int) string {
	var result []rune
	for ; i >= 0; i-- {
		result = append(result, rune(table[i][j]))
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

// TODO
// write similar to rightDownSubstring
func rightUpSubstring(table []string, i, j int) string {
	return ""
}

// TODO
// write similar to rightDownSubstring
func leftDownSubstring(table []string, i, j int) string {
	return ""
}

// TODO
// write similar to rightDownSubstring
func leftUpSubstring(table []string, i, j int) string {
	// var result []rune
	// for i >= 0 && j >= 0 {
	// 	result = append(result, rune(table[i][j]))
	// 	i--
	// 	j--
	// }
	// return string(result)
	return ""
}
