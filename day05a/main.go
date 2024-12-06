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
			fields := toIntSlice(strings.Split(line, ","))
			updates = append(updates, update{
				pages: fields,
				safe:  false,
			})
		}
	}

	vPrint("Rules:\n%#v\n\n", rules)
	vPrint("Updates:\n%#v\n\n", updates)

	// Result of all middle values of valid updates
	var res int = 0

	// For each update check if is safe
	for _, update := range updates {
		// page is correctly first?
		// yes if there are rules that put it before each of the other pages
		vPrint("Running check against update %v:\n", update)
		update.checkSafety(rules)
		if verbose {
			str := ""
			if !update.safe {
				str = " NOT"
			}
			vPrint("Update %v is%s safe\n\n\n", update, str)
		}
		if update.safe {
			res += update.getMiddlePage()
		}
	}

	// Print result
	fmt.Printf("Result: %d\n", res)
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

func toIntSlice(slice []string) []int {
	ints := make([]int, len(slice))
	for i, s := range slice {
		ints[i] = toInt(s)
	}
	return ints
}

func toInt(val string) int {
	num, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		log.Fatalf("Error converting value %s to uint32: %v.", val, err)
	}
	return int(num)
}
