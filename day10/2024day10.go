package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day10.txt")
	day4data := strings.Split(string(bytesread), "\n")

	var convertedMap [][]int

	part1Answer := 0
	part2Answer := 0
	var trailheads [][2]int

	for i := 0; i < len(day4data); i++ {
		var row []int
		for j := 0; j < len(day4data[0]); j++ {
			value, _ := strconv.Atoi(string(day4data[i][j]))
			// fmt.Println(value)
			if value == 0 {
				trailheadCoord := [2]int{i, j}
				trailheads = append(trailheads, trailheadCoord)
			}
			row = append(row, value)
		}
		convertedMap = append(convertedMap, row)
	}
	for trail := range trailheads {
		retval := breadthFirstSearch(convertedMap, trailheads[trail])
		part1Answer += retval[0]
		part2Answer += retval[1]
	}

	// for trail := range trailheads {
	// 	part2Answer += breadthFirstSearchPaths(day4data, trailheads[trail])
	// }

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func breadthFirstSearch(trailMap [][]int, startingCoord [2]int) [2]int {
	visitedCoords := make(map[[2]int]bool)
	neighbours := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	stack := [][2]int{startingCoord}
	finalPeaks := make(map[[2]int]int)

	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		value := trailMap[current[0]][current[1]]
		for i := range neighbours {
			neighbour := [2]int{current[0] + neighbours[i][0], current[1] + neighbours[i][1]}
			if visitedCoords[neighbour] {
				continue
			}
			if neighbour[0] < 0 || neighbour[1] < 0 || neighbour[0] >= len(trailMap) || neighbour[1] >= len(trailMap[0]) {
				continue
			}
			neighbourValue := trailMap[neighbour[0]][neighbour[1]]

			// fmt.Println(neighbourValue)
			if neighbourValue == value+1 {
				stack = append(stack, neighbour)
				if neighbourValue == 9 {
					finalPeaks[neighbour] += 1
				}
			}
		}
		visitedCoords[current] = true
	}
	retVal := 0
	for _, val := range finalPeaks {
		retVal += val
	}
	peaks := len(finalPeaks)
	return [2]int{peaks, retVal}
}
