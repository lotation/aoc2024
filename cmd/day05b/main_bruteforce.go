package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/lotation/aoc2024/internal/utils"
)

type rule struct {
	before int
	after  int
}

type update struct {
	pages []int
	safe  bool
}

func (update *update) checkSafety(rules []rule) {
	(*update).safe = check((*update).pages, 0, rules)
}

func (update update) getMiddlePage() int {
	return update.pages[len(update.pages)/2]
}

func (update *update) Order(rules []rule) {
	// bruteforce all possible permutations of the array
	// to find the correct one
	for _, u := range permutations((*update).pages) {
		if check(u, 0, rules) {
			// found the right permutation
			(*update).pages = u
			return
		}
	}
}

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

	// Rules
	var rules []rule = nil
	// Updates
	var updates []update = nil
	// flag to know if still parsing rules
	var stillRules bool = true

	// Parse inut file
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if separator was hit
		if line == "" {
			stillRules = false
			continue
		}
		if stillRules {
			// Still parsing rules
			fields := strings.Split(line, "|")
			rules = append(rules, rule{
				before: toInt(fields[0]),
				after:  toInt(fields[1]),
			})
		} else {
			fields := utils.ToIntSlice(strings.Split(line, ","))
			updates = append(updates, update{
				pages: fields,
				safe:  false,
			})
		}
	}

	vPrint("Rules:\n%#v\n\n", rules)
	vPrint("Updates:\n%#v\n\n", updates)

	// Sum of all the middle values of valid updates
	var res1 int = 0
	// Sum of all the middle values of invalid updates
	var res2 int = 0

	// For each update check if is safe
	for _, update := range updates {
		// page is correctly first?
		// yes if there are rules that put it before each of the other pages
		vPrint("Running check against update %v:\n", update.pages)
		update.checkSafety(rules)
		if verbose {
			str := ""
			if !update.safe {
				str = " NOT"
			}
			vPrint("Update %v is%s safe\n", update.pages, str)
		}
		if update.safe {
			res1 += update.getMiddlePage()
		} else {
			vPrint("Ordering update %v ...\n", update.pages)
			update.Order(rules)
			vPrint("Update %v reordered to %v\n", old.pages, update.pages)
			res2 += update.getMiddlePage()
		}
		// vPrint("\n\n")
		vPrint("Done update %v\n\n", update)
	}

	// Print result
	fmt.Printf("Result 1: %d\n", res1)
	fmt.Printf("Result 2: %d\n", res2)
}

// Check if requested page is before the others
func check(pages []int, start int, rules []rule) bool {
	vPrint("check() called on %v with start %d [%d]\n", pages, pages[start], start)

	// caso base
	if start == len(pages)-1 {
		vPrint("\n")
		return true
	}

	count := 0

	for i := start + 1; i < len(pages); i++ {
		vPrint("%d) trying %d:\n", i, pages[i])
		// check if there is a rule with start in before and all the other pages in after
		for _, rule := range rules {
			vPrint("  pages[start=%d]=%d == rule.before=%d - pages[i]=%d - rule.after=%d",
				start, pages[start], rule.before, pages[i], rule.after)
			if pages[start] == rule.before && pages[i] == rule.after {
				count++
				vPrint("  {count=%d}\n", count)
				break
			}
			vPrint("\n")
		}

		if count == len(pages[start:])-1 {
			vPrint("\n")
			return check(pages, start+1, rules)
		}
		vPrint("\n")
	}

	vPrint("\n")
	return false
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
