package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"

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
	vPrint("Starting position: %c %s\n\n", mmap.GetPosition().Char, mmap.Pos)

	//rect_buffer := make([]Position, 4)
	var loopCount int = 0

	//rect_buffer[0] = mmap.Pos

	var rect_buffer []Position
	rect_buffer = append(rect_buffer, mmap.Pos)
	fmt.Printf("treating start position as corner.  {%v  %d}\n", rect_buffer, len(rect_buffer))

	vPrint("Starting guard patrol:\n")
	for {
		switch mmap.Pos.direction {
		case Up:
			vPrint("Facing Up, ")
			// check if going out of the map
			if mmap.Pos.y-1 < 0 {
				vPrint("next step up is out of map, reached END!\n\n")
				goto end
			}
			if mmap.NextUp().IsObstacle() {
				// keep only the last 4 corners
				if len(rect_buffer) >= 4 {
					rect_buffer = rect_buffer[1:]
				}
				// this cell is a corner
				// add it to the rectangle buffer
				rect_buffer = append(rect_buffer, mmap.Pos)
				fmt.Printf("corner detected %s, saving it.  {%v  %d}\n", mmap.Pos, rect_buffer, len(rect_buffer))

				// facing an obstacle, turn right 90 degrees
				mmap.SetDirection(Right)
				vPrint("obstacle detected at [%d,%d], turning right 90 degrees, now going Rigth\n", mmap.Pos.y-1, mmap.Pos.x)
			} else {
				if len(rect_buffer) >= 4 {
					// check right as if you encountered an obstacle
					fmt.Printf("pos=%d {%s}- 2nd corner=%d {%s}\n", mmap.Pos.y, mmap.Pos, rect_buffer[1].y, rect_buffer[1])
					// prova a vedere se girando a destra incontri uno spigolo
					if mmap.Pos.y == rect_buffer[1].y {
						loopCount++
						fmt.Printf("  => LOOP detected at %s  {%d}\n", mmap.Pos, loopCount)
					}
				}

				// take a step forward up
				mmap.GetPosition().Char = Visited
				mmap.Pos.y -= 1
				if !mmap.GetPosition().Visited {
					// mmap.GetPosition().MarkVisited()
					mmap.GetPosition().Visited = true
					mmap.GetPosition().Char = Up
					count++
				}
				vPrint("taking a step forward => %s  {visited %d}\n", mmap.Pos, count)
			}
		case Right:
			vPrint("Facing Right, ")
			// check if going out of the map
			if mmap.Pos.x+1 >= len(mmap.Cells[0]) {
				vPrint("next step right is out of map, reached END!\n\n")
				goto end
			}
			if mmap.NextRight().IsObstacle() {
				// keep only the last 4 corners
				if len(rect_buffer) >= 4 {
					rect_buffer = rect_buffer[1:]
				}
				// this cell is a corner
				// add it to the rectangle buffer
				rect_buffer = append(rect_buffer, mmap.Pos)
				fmt.Printf("corner detected %s, saving it.  {%v  %d}\n", mmap.Pos, rect_buffer, len(rect_buffer))

				// facing an obstacle, turn right 90 degrees
				mmap.SetDirection(Down)
				vPrint("obstacle detected at [%d,%d], turning right 90 degrees, now going Down\n", mmap.Pos.y, mmap.Pos.x+1)
			} else {
				if len(rect_buffer) >= 4 {
					// check down as if you encountered an obstacle
					fmt.Printf("pos=%d {%s}- 2nd corner=%d {%s}\n", mmap.Pos.x, mmap.Pos, rect_buffer[1].x, rect_buffer[1])
					// prova a vedere se girando a destra incontri uno spigolo
					if mmap.Pos.x == rect_buffer[1].x {
						loopCount++
						fmt.Printf("  => LOOP detected at %s  {%d}\n", mmap.Pos, loopCount)
					}
				}

				// take a step forward right
				mmap.GetPosition().Char = Visited
				mmap.Pos.x += 1
				if !mmap.GetPosition().Visited {
					// mmap.GetPosition().MarkVisited()
					mmap.GetPosition().Visited = true
					mmap.GetPosition().Char = Right
					count++
				}
				vPrint("taking a step forward => %s  {visited %d}\n", mmap.Pos, count)
			}
		case Down:
			vPrint("Facing Down, ")
			// check if going out of the map
			if mmap.Pos.y+1 >= len(mmap.Cells) {
				vPrint("next step down is out of map, reached END!\n\n")
				goto end
			}
			if mmap.NextDown().IsObstacle() {
				// keep only the last 4 corners
				if len(rect_buffer) >= 4 {
					rect_buffer = rect_buffer[1:]
				}
				// this cell is a corner
				// add it to the rectangle buffer
				rect_buffer = append(rect_buffer, mmap.Pos)
				fmt.Printf("corner detected %s, saving it.  {%v  %d}\n", mmap.Pos, rect_buffer, len(rect_buffer))

				// facing an obstacle, turn right 90 degrees
				mmap.SetDirection(Left)
				vPrint("obstacle detected at [%d,%d], turning right 90 degrees, now going Left\n", mmap.Pos.y+1, mmap.Pos.x)
			} else {
				if len(rect_buffer) >= 4 {
					// check left as if you encountered an obstacle
					fmt.Printf("pos=%d {%s}- 2nd corner=%d {%s}\n", mmap.Pos.y, mmap.Pos, rect_buffer[1].y, rect_buffer[1])
					// prova a vedere se girando a destra incontri uno spigolo
					if mmap.Pos.y == rect_buffer[1].y {
						loopCount++
						fmt.Printf("  => LOOP detected at %s  {%d}\n", mmap.Pos, loopCount)
					}
				}

				// take a step forward down
				mmap.GetPosition().Char = Visited
				mmap.Pos.y += 1
				if !mmap.GetPosition().Visited {
					// mmap.GetPosition().MarkVisited()
					mmap.GetPosition().Visited = true
					mmap.GetPosition().Char = Down
					count++
				}
				vPrint("taking a step forward => %s  {visited %d}\n", mmap.Pos, count)
			}
		case Left:
			vPrint("Facing Left, ")
			// check if going out of the map
			if mmap.Pos.x-1 < 0 {
				vPrint("next step left is out of map, reached END!\n\n")
				goto end
			}
			if mmap.NextLeft().IsObstacle() {
				// keep only the last 4 corners
				if len(rect_buffer) >= 4 {
					rect_buffer = rect_buffer[1:]
				}
				// this cell is a corner
				// add it to the rectangle buffer
				rect_buffer = append(rect_buffer, mmap.Pos)
				fmt.Printf("corner detected %s, saving it.  {%v  %d}\n", mmap.Pos, rect_buffer, len(rect_buffer))

				// facing an obstacle, turn right 90 degrees
				mmap.SetDirection(Up)
				vPrint("obstacle detected at [%d,%d], turning right 90 degrees, now going Up\n", mmap.Pos.y, mmap.Pos.x-1)
			} else {
				if len(rect_buffer) >= 4 {
					// check up as if you encountered an obstacle
					fmt.Printf("pos=%d {%s}- 2nd corner=%d {%s}\n", mmap.Pos.x, mmap.Pos, rect_buffer[1].x, rect_buffer[1])
					// prova a vedere se girando a destra incontri uno spigolo
					if mmap.Pos.x == rect_buffer[1].x {
						loopCount++
						fmt.Printf("  => LOOP detected at %s  {%d}\n", mmap.Pos, loopCount)
					}
				}

				// take a step forward left
				mmap.GetPosition().Char = Visited
				mmap.Pos.x -= 1
				if !mmap.GetPosition().Visited {
					// mmap.GetPosition().MarkVisited()
					mmap.GetPosition().Visited = true
					mmap.GetPosition().Char = Left
					count++
				}
				vPrint("taking a step forward => %s  {visited %d}\n", mmap.Pos, count)
			}
		}
	}
end:
	// mark last position as visited
	mmap.GetPosition().MarkVisited()
	// Print result
	fmt.Printf("Distinct positions (%c) visited by the guard: %d\n", Visited, count)
}
