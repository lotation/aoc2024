package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/lotation/aoc2024/internal/utils"
)

var verbose bool = true

func vPrint(format string, args ...interface{}) {
	if verbose {
		fmt.Printf(format, args...)
	}
}

var graphic bool = false

func gPrint(format string, args ...interface{}) {
	if graphic {
		fmt.Printf(format, args...)
	}
}

func gClearScreen() {
	gPrint("\033[2J\033[H")
}

func gPrintMap(m Map) {
	if graphic {
		fmt.Println(m)
		//time.Sleep(500 * time.Millisecond)
		time.Sleep(100 * time.Millisecond)
		// return the cursor to the line after the fixed lines
		gPrint("\033[2;1H")
		// clear remaining lines
		for i := 0; i < len(m.Cells)-1; i++ {
			gPrint("\033[K")
		}
	}
}

func main() {
	var inputfile string
	flag.StringVar(&inputfile, "input", "input.txt", "path to file containing current day input")
	flag.BoolVar(&verbose, "verbose", true, "enable verbose logging")
	flag.BoolVar(&graphic, "graphic", true, "enable graphic representation of the guard patrol")
	flag.Parse()

	if verbose && graphic {
		log.Panic("graphics can't be enabled with verbose logging")
	}

	// Open input file
	fp := utils.Fopen(inputfile)
	defer func() {
		if err := fp.Close(); err != nil {
			log.Panic(err)
		}
	}()

	var count int = 0
	var mmap Map
	var linecells []Cell

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		linecells = nil
		for _, c := range line { //[]rune(line)
			linecells = append(linecells, NewCell(c))
		}
		mmap.Cells = append(mmap.Cells, linecells)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input file %s: %v", inputfile, err)
		return
	}

	// clear the screen
	gClearScreen()
	gPrint("Starting guard patrol:\n")
	// move the cursor to the line after the fixed lines
	gPrint("\033[2;1H")
	// print initial map
	gPrintMap(mmap)
	// mmap.PrintMap()
	// time.Sleep(250 * time.Millisecond)
	// // return the cursor to the line after the fixed lines
	// gPrint("\033[2;1H")
	// // clear remaining lines
	// for i := 0; i < len(mmap.Cells)-1; i++ {
	// 	gPrint("\033[K")
	//}

	// look for the start position
	for row := 0; row < len(mmap.Cells); row++ {
		if len(mmap.Cells[row]) != len(mmap.Cells[0]) {
			log.Panicf("Error: length of map rows not equal: %d\n", len(mmap.Cells[row]))
		}
		for col := 0; col < len(mmap.Cells[row]); col++ {
			switch mmap.Cells[row][col].Char {
			case Up, Right, Down, Left:
				mmap.Pos = NewPosition(row, col, mmap.Cells[row][col].Char)
				mmap.Cells[row][col].Visited = true
				count++
			}
			vPrint("%c", mmap.Cells[row][col].Char)
		}
		vPrint("\n")
	}
	vPrint("\n")
	vPrint("Starting position: %c [%d,%d]\n\n", mmap.GetPosition().Char, mmap.Pos.y, mmap.Pos.x)

	vPrint("Starting guard patrol:\n")

	for {
		// update map
		gPrintMap(mmap)
		// if graphic {
		// 	mmap.PrintMap()
		// 	time.Sleep(250 * time.Millisecond)
		// 	// return the cursor to the line after the fixed lines
		// 	gPrint("\033[2;1H")
		// 	// clear remaining lines
		// 	for i := 0; i < len(mmap.Cells)-1; i++ {
		// 		gPrint("\033[K")
		// 	}
		// }
		switch mmap.Pos.direction {
		case Up:
			vPrint("Facing Up, ")
			// check if going out of the map
			if mmap.Pos.y-1 < 0 {
				vPrint("next step up is out of map, reached END!\n\n")
				goto end
			}
			if mmap.NextUp().IsObstacle() {
				// facing an obstacle, turn right 90 degrees
				mmap.SetDirection(Right)
				vPrint("obstacle detected at [%d,%d], turning right 90 degrees, now going Rigth\n", mmap.Pos.y-1, mmap.Pos.x)
			} else { // take a step forward up
				mmap.GetPosition().Char = Visited
				mmap.Pos.y -= 1
				if !mmap.GetPosition().Visited {
					// mmap.GetPosition().MarkVisited()
					mmap.GetPosition().Visited = true
					mmap.GetPosition().Char = Up
					count++
				}
				vPrint("taking a step forward => [%d,%d]  {visited %d}\n", mmap.Pos.y, mmap.Pos.x, count)
			}
		case Right:
			vPrint("Facing Right, ")
			// check if going out of the map
			if mmap.Pos.x+1 >= len(mmap.Cells[0]) {
				vPrint("next step right is out of map, reached END!\n\n")
				goto end
			}
			if mmap.NextRight().IsObstacle() {
				// facing an obstacle, turn right 90 degrees
				mmap.SetDirection(Down)
				vPrint("obstacle detected at [%d,%d], turning right 90 degrees, now going Down\n", mmap.Pos.y, mmap.Pos.x+1)
			} else { // take a step forward right
				mmap.GetPosition().Char = Visited
				mmap.Pos.x += 1
				if !mmap.GetPosition().Visited {
					// mmap.GetPosition().MarkVisited()
					mmap.GetPosition().Visited = true
					mmap.GetPosition().Char = Right
					count++
				}
				vPrint("taking a step forward => [%d,%d]  {visited %d}\n", mmap.Pos.y, mmap.Pos.x, count)
			}
		case Down:
			vPrint("Facing Down, ")
			// check if going out of the map
			if mmap.Pos.y+1 >= len(mmap.Cells) {
				vPrint("next step down is out of map, reached END!\n\n")
				goto end
			}
			if mmap.NextDown().IsObstacle() {
				// facing an obstacle, turn right 90 degrees
				mmap.SetDirection(Left)
				vPrint("obstacle detected at [%d,%d], turning right 90 degrees, now going Left\n", mmap.Pos.y+1, mmap.Pos.x)
			} else { // take a step forward down
				mmap.GetPosition().Char = Visited
				mmap.Pos.y += 1
				if !mmap.GetPosition().Visited {
					// mmap.GetPosition().MarkVisited()
					mmap.GetPosition().Visited = true
					mmap.GetPosition().Char = Down
					count++
				}
				vPrint("taking a step forward => [%d,%d]  {visited %d}\n", mmap.Pos.y, mmap.Pos.x, count)
			}
		case Left:
			vPrint("Facing Left, ")
			// check if going out of the map
			if mmap.Pos.x-1 < 0 {
				vPrint("next step left is out of map, reached END!\n\n")
				goto end
			}
			if mmap.NextLeft().IsObstacle() {
				// facing an obstacle, turn right 90 degrees
				mmap.SetDirection(Up)
				vPrint("obstacle detected at [%d,%d], turning right 90 degrees, now going Up\n", mmap.Pos.y, mmap.Pos.x-1)
			} else { // take a step forward left
				mmap.GetPosition().Char = Visited
				mmap.Pos.x -= 1
				if !mmap.GetPosition().Visited {
					// mmap.GetPosition().MarkVisited()
					mmap.GetPosition().Visited = true
					mmap.GetPosition().Char = Left
					count++
				}
				vPrint("taking a step forward => [%d,%d]  {visited %d}\n", mmap.Pos.y, mmap.Pos.x, count)
			}
		}
	}
end:
	// mark last position as visited
	mmap.GetPosition().MarkVisited()
	// print last map update
	gPrintMap(mmap)
	// move the cursor to the end
	gPrint("\033[%d;1H", 1+len(mmap.Cells)+2)
	// Print result
	fmt.Printf("Distinct positions (%c) visited by the guard: %d\n", Visited, count)
}
