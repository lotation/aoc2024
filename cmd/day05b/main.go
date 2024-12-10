package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"slices"
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

func (update *update) checkSafety() {
	(*update).safe = check((*update).pages)
}

func (update update) getMiddlePage() int {
	return update.pages[len(update.pages)/2]
}

func (update *update) order() {
	pages := (*update).pages
	ordered := make([]int, len(pages))
	before_count := make([]int, slices.Max(pages)+1)
	after_count := make([]int, slices.Max(pages)+1)

	// find first element
	for _, page := range pages {
		for _, rule := range rules {
			// count how many before occurrences
			if rule.before == page && slices.Contains(pages, rule.after) {
				before_count[page]++
			}
			// count how many after occurrences
			if rule.after == page && slices.Contains(pages, rule.before) {
				after_count[page]++
			}
		}
	}

	for indx := 0; indx < len(pages); indx++ {
		// increase to differentiate the 0 index from the 0 match
		before_count[pages[indx]] += 1
		after_count[pages[indx]] += 1
	}

	for j := 0; j < len(pages); {
		for i := 0; i < len(after_count); i++ {
			if after_count[i] == j+1 {
				vPrint("i=%d - after=%d - before=%d\n", i, after_count[i], before_count[i])
				ordered[j] = i
				j++
				break // end current loop as the right page was found
			}
		}
	}

	vPrint("  => %v\n", ordered)

	(*update).pages = slices.Clone(ordered)
	return
}

var verbose bool = true

func vPrint(format string, args ...interface{}) {
	if verbose {
		fmt.Printf(format, args...)
	}
}

// Rules
var rules []rule

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
	//var rules []rule = nil
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
				before: utils.ToInt(fields[0]),
				after:  utils.ToInt(fields[1]),
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
		update.checkSafety()
		if verbose {
			if update.safe {
				vPrint("Update %v is safe\n", update.pages)
			} else {
				vPrint("Update %v is NOT safe, fixing it...\n", update.pages)
			}
		}
		if update.safe {
			res1 += update.getMiddlePage()
		} else {
			vPrint("Ordering update %v...\n", update.pages)
			update.order()
			// vPrint("Update %v reordered to %v\n", old.pages, update.pages)
			res2 += update.getMiddlePage()
		}
		vPrint("\n\n")
	}

	// Print result
	fmt.Printf("Result 1: %d\n", res1)
	fmt.Printf("Result 2: %d\n", res2)
}

// Check if requested page is before the others
func check(pages []int) bool {
	vPrint("check() called on %v\n", pages)

	// Iterate through all pages except the last one
	for currentStart := 0; currentStart < len(pages)-1; currentStart++ {
		count := 0

		// Check all pages after the current start
		for i := currentStart + 1; i < len(pages); i++ {
			vPrint("%d) trying %d:\n", i, pages[i])
			for _, rule := range rules {
				vPrint("  pages[start=%d]=%d == rule.before=%d - pages[i]=%d - rule.after=%d",
					currentStart, pages[currentStart], rule.before, pages[i], rule.after)
				if pages[currentStart] == rule.before && pages[i] == rule.after {
					count++
					vPrint("  {count=%d}\n", count)
					break
				}
				vPrint("\n")
			}
		}

		// If count matches the number of pages after currentStart, continue to the next start
		if count == len(pages[currentStart:])-1 {
			continue
		}

		// If we reach here, it means the condition was not satisfied
		vPrint("\n")
		return false
	}

	// If we complete the loop without returning false, return true
	vPrint("\n")
	return true
}

func remove(slice []int, index int) []int {
	if index >= 0 && index < len(slice) {
		ret := make([]int, 0, len(slice)-1)
		ret = append(ret, slice[:index]...)
		return append(ret, slice[index+1:]...)
	} else {
		return slice
	}
}
