package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day6.txt")
	day6data := strings.Split(string(bytesread), "\n")
	mapHeight := len(day6data)
	mapWidth := len(day6data[0])
	currentPosX := 0
	currentPosY := 0
	currentDir := 0
	part1Answer := 0
	part2Answer := 0

	visitedCoords := make(map[[2]int]bool)

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for i, _ := range day6data {
		for j, val := range day6data[i] {
			if val == '^' {
				currentPosY = i
				currentPosX = j
				break
			}
		}
		if currentPosX != 0 {
			break
		}
	}
	startingPosX := currentPosX
	startingPosY := currentPosY

	for {
		nextY := currentPosY + directions[currentDir][0]
		nextX := currentPosX + directions[currentDir][1]
		coord := [2]int{currentPosY, currentPosX}

		if nextX < 0 || nextY < 0 || nextX >= mapWidth || nextY >= mapHeight {
			visitedCoords[coord] = true
			break
		}
		if day6data[nextY][nextX] == '#' {
			currentDir = (currentDir + 1) % len(directions)
			continue
		}

		visitedCoords[coord] = true
		currentPosY = nextY
		currentPosX = nextX

	}
	part1Answer = len(visitedCoords)

	// for visitedCoord, _ := range visitedCoords {
	// 	if visitedCoord[0] == startingPosY && visitedCoord[1] == startingPosX {
	// 		continue
	// 	}
	// 	if loopDetector(startingPosX, startingPosY, 0, day6data, visitedCoord[0], visitedCoord[1]) {
	// 		part2Answer++
	// 	}
	// }

	var mu sync.Mutex
	var wg sync.WaitGroup

	for visitedCoord := range visitedCoords {
		if visitedCoord[0] == startingPosY && visitedCoord[1] == startingPosX {
			continue
		}
		wg.Add(1)
		go func(visitedCoord [2]int) {
			if loopDetector(startingPosX, startingPosY, 0, day6data, visitedCoord[0], visitedCoord[1]) {
				mu.Lock()
				part2Answer++
				mu.Unlock()
			}
			defer wg.Done()
		}(visitedCoord)

	}
	wg.Wait()
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func loopDetector(currentPosX int, currentPosY int, currentDir int, paraMap []string, blockY int, blockX int) bool {
	mapHeight := len(paraMap)
	mapWidth := len(paraMap[0])
	visitedCoords := make(map[[2]int][]int)
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for {
		nextY := currentPosY + directions[currentDir][0]
		nextX := currentPosX + directions[currentDir][1]
		coord := [2]int{currentPosY, currentPosX}
		if nextX < 0 || nextY < 0 || nextX >= mapWidth || nextY >= mapHeight {
			return false
		}
		if paraMap[nextY][nextX] == '#' || (nextX == blockX && nextY == blockY) {
			visitedCoords[coord] = append(visitedCoords[coord], currentDir)
			currentDir = (currentDir + 1) % len(directions)
			continue
		}

		val, ok := visitedCoords[coord]

		if slices.Contains(val, currentDir) && ok {
			return true
		}
		visitedCoords[coord] = append(visitedCoords[coord], currentDir)
		currentPosY = nextY
		currentPosX = nextX
	}
}
