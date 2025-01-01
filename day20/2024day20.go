package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day20.txt")
	day20data := strings.Split(string(bytesread), "\n")
	part1Answer := 0
	part2Answer := 0
	var raceMap [][]string
	var startPoint [2]int
	var finishPoint [2]int
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	pathMap := make(map[[2]int]int)

	for i := range day20data {
		var row []string
		for j := range day20data[i] {
			row = append(row, string(day20data[i][j]))
			if day20data[i][j] == 'E' {
				finishPoint = [2]int{i, j}
			}
			if day20data[i][j] == 'S' {
				startPoint = [2]int{i, j}
			}
		}
		raceMap = append(raceMap, row)
	}
	distance := 0
	current := startPoint
	for {
		// fmt.Println(current)
		if tupleEquals(current, finishPoint) {
			pathMap[current] = distance
			break
		}
		for _, neighbour := range directions {
			candidate := tupleAddition(current, neighbour)
			if pathMap[candidate] != 0 {
				continue
			}
			if raceMap[candidate[0]][candidate[1]] == "." || tupleEquals(candidate, finishPoint) {
				pathMap[current] = distance
				current = candidate
				break
			}
		}
		distance++
	}

	cheatMap := make(map[int]int)

	for key := range pathMap {
		for _, direction := range directions {
			// neighbour := tupleAddition(direction, key)
			jump := tupleAddition(direction, direction)
			cheatSquare := tupleAddition(jump, key)

			if (pathMap[cheatSquare] - 2) > pathMap[key] {
				cheatMap[pathMap[cheatSquare]-pathMap[key]-2]++
			}
		}
	}
	for key, value := range cheatMap {
		if key >= 100 {
			part1Answer += value
		}
	}

	cheatMap2 := make(map[int]int)
	visited := make(map[[2]int]bool)

	for key := range pathMap {
		for key2 := range pathMap {
			manhattanDistance := manhattanDistance(key, key2)
			if manhattanDistance > 20 {
				continue
			}
			if pathMap[key]-manhattanDistance > pathMap[key2] {
				continue
			}
			cheatMap2[pathMap[key2]-pathMap[key]-manhattanDistance]++
		}
		visited[key] = true
	}

	for key, value := range cheatMap2 {
		if key >= 100 {
			part2Answer += value
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func tupleAddition(first [2]int, second [2]int) [2]int {
	return [2]int{first[0] + second[0], first[1] + second[1]}
}

func tupleEquals(first [2]int, second [2]int) bool {
	return first[0] == second[0] && first[1] == second[1]
}

func manhattanDistance(first [2]int, second [2]int) int {
	yDiff := first[0] - second[0]
	if yDiff < 1 {
		yDiff *= -1
	}
	xDiff := first[1] - second[1]
	if xDiff < 1 {
		xDiff *= -1
	}
	return xDiff + yDiff
}
