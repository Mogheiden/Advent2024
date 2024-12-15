package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day15pt2.txt")
	day15data := strings.Split(string(bytesread), "\n\n")
	boxMap := strings.Split(day15data[0], "\n")
	var boxMapArray [][]string

	part2Answer := 0

	for _, row := range boxMap {
		arrayRow := strings.Split(row, "")
		boxMapArray = append(boxMapArray, arrayRow)
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	inputString := day15data[1]
	currentPos := [2]int{0, 0}

	for i := range boxMapArray {
		for j := range boxMapArray[0] {
			if boxMapArray[i][j] == "@" {
				currentPos[0] = i
				currentPos[1] = j
				break
			}
		}
		if currentPos[0] != 0 {
			break
		}
	}

	for _, command := range inputString {
		switch command {
		case '^':
			if recursiveMoveable(boxMapArray, currentPos, directions[UP]) {
				boxMapArray = pushTiles(boxMapArray, currentPos, directions[UP])
				boxMapArray[currentPos[0]][currentPos[1]] = "."
				currentPos[0] += directions[UP][0]
				currentPos[1] += directions[UP][1]
				boxMapArray[currentPos[0]][currentPos[1]] = "@"
			}
		case 'v':
			if recursiveMoveable(boxMapArray, currentPos, directions[DOWN]) {
				boxMapArray = pushTiles(boxMapArray, currentPos, directions[DOWN])
				boxMapArray[currentPos[0]][currentPos[1]] = "."
				currentPos[0] += directions[DOWN][0]
				currentPos[1] += directions[DOWN][1]
				boxMapArray[currentPos[0]][currentPos[1]] = "@"
			}
		case '<':
			if recursiveMoveable(boxMapArray, currentPos, directions[LEFT]) {
				boxMapArray = pushTiles(boxMapArray, currentPos, directions[LEFT])
				boxMapArray[currentPos[0]][currentPos[1]] = "."
				currentPos[0] += directions[LEFT][0]
				currentPos[1] += directions[LEFT][1]
				boxMapArray[currentPos[0]][currentPos[1]] = "@"
			}
		case '>':
			if recursiveMoveable(boxMapArray, currentPos, directions[RIGHT]) {
				boxMapArray = pushTiles(boxMapArray, currentPos, directions[RIGHT])
				boxMapArray[currentPos[0]][currentPos[1]] = "."
				currentPos[0] += directions[RIGHT][0]
				currentPos[1] += directions[RIGHT][1]
				boxMapArray[currentPos[0]][currentPos[1]] = "@"
			}
		}

	}

	for i := range boxMapArray {
		for j := range boxMapArray[0] {
			if boxMapArray[i][j] == "[" {
				part2Answer += (100*i + j)
			}
		}
	}

	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

const (
	UP    int = 0
	DOWN  int = 1
	LEFT  int = 2
	RIGHT int = 3
)

func recursiveMoveable(boxMap [][]string, current [2]int, direction [2]int) bool {
	current[0] += direction[0]
	current[1] += direction[1]
	switch boxMap[current[0]][current[1]] {
	case "[":
		left := [2]int{current[0], current[1]}
		right := [2]int{current[0], current[1] + 1}
		if direction[0] == 0 && direction[1] == 1 {
			return recursiveMoveable(boxMap, right, direction)
		} else {
			return recursiveMoveable(boxMap, left, direction) && recursiveMoveable(boxMap, right, direction)
		}
	case "]":
		left := [2]int{current[0], current[1] - 1}
		right := [2]int{current[0], current[1]}
		if direction[0] == 0 && direction[1] == -1 {
			return recursiveMoveable(boxMap, left, direction)
		} else {
			return recursiveMoveable(boxMap, left, direction) && recursiveMoveable(boxMap, right, direction)
		}
	case "#":
		return false
	case ".":
		return true
	default:
		panic("unrecognized escape character")
	}
}

func pushTiles(boxMap [][]string, current [2]int, direction [2]int) [][]string {
	firstTile := tileRecord{tileOriginal: current, tileValue: boxMap[current[0]][current[1]]}
	tileQueue := []tileRecord{firstTile}
	affectedCoords := make(map[[2]int]bool)
	tileSet := make(map[tileRecord]bool)

	for len(tileQueue) > 0 {
		current := tileQueue[0]
		currentCoord := current.tileOriginal
		tileQueue = tileQueue[1:]
		nextY := currentCoord[0] + direction[0]
		nextX := currentCoord[1] + direction[1]
		nextVal := boxMap[nextY][nextX]

		if nextVal == "." {
			continue
		} else if direction[0] == 0 {
			newTileCoord := [2]int{nextY, nextX}
			newTileoordNext := [2]int{newTileCoord[0] + direction[0], newTileCoord[1] + direction[1]}
			newTile := tileRecord{tileOriginal: newTileCoord, tileCoord: newTileoordNext, tileValue: nextVal}
			tileQueue = append(tileQueue, newTile)
			tileSet[newTile] = true
			affectedCoords[newTileCoord] = true
		} else if nextVal == "[" {
			leftTileCoord := [2]int{nextY, nextX}
			leftTileCoordNext := [2]int{leftTileCoord[0] + direction[0], leftTileCoord[1] + direction[1]}
			rightTileCoord := [2]int{nextY, nextX + 1}
			rightTileCoordNext := [2]int{rightTileCoord[0] + direction[0], rightTileCoord[1] + direction[1]}
			newLeftTile := tileRecord{tileOriginal: leftTileCoord, tileCoord: leftTileCoordNext, tileValue: nextVal}
			newRightTile := tileRecord{tileOriginal: rightTileCoord, tileCoord: rightTileCoordNext, tileValue: boxMap[nextY][nextX+1]}
			tileQueue = append(tileQueue, newLeftTile)
			tileQueue = append(tileQueue, newRightTile)
			tileSet[newLeftTile] = true
			tileSet[newRightTile] = true
			affectedCoords[leftTileCoord] = true
			affectedCoords[rightTileCoord] = true
		} else if nextVal == "]" {
			leftTileCoord := [2]int{nextY, nextX - 1}
			leftTileCoordNext := [2]int{leftTileCoord[0] + direction[0], leftTileCoord[1] + direction[1]}
			rightTileCoord := [2]int{nextY, nextX}
			rightTileCoordNext := [2]int{rightTileCoord[0] + direction[0], rightTileCoord[1] + direction[1]}
			newLeftTile := tileRecord{tileOriginal: leftTileCoord, tileCoord: leftTileCoordNext, tileValue: boxMap[nextY][nextX-1]}
			newRightTile := tileRecord{tileOriginal: rightTileCoord, tileCoord: rightTileCoordNext, tileValue: boxMap[nextY][nextX]}
			tileQueue = append(tileQueue, newLeftTile)
			affectedCoords[leftTileCoord] = true
			affectedCoords[rightTileCoord] = true
			tileQueue = append(tileQueue, newRightTile)
			tileSet[newLeftTile] = true
			tileSet[newRightTile] = true
		}
	}
	for key, _ := range tileSet {
		tileBehind := [2]int{key.tileOriginal[0] - direction[0], key.tileOriginal[1] - direction[1]}
		if !affectedCoords[tileBehind] {
			boxMap[key.tileOriginal[0]][key.tileOriginal[1]] = "."
		}
		boxMap[key.tileCoord[0]][key.tileCoord[1]] = key.tileValue
	}
	return boxMap
}

type tileRecord struct {
	tileOriginal [2]int
	tileCoord    [2]int
	tileValue    string
}
