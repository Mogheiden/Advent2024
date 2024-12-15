package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	bytesread, _ := os.ReadFile("day15.txt")
// 	day15data := strings.Split(string(bytesread), "\n\n")
// 	boxMap := strings.Split(day15data[0], "\n")
// 	var boxMapArray [][]string

// 	part1Answer := 0
// 	part2Answer := 0

// 	for _, row := range boxMap {
// 		arrayRow := strings.Split(row, "")
// 		boxMapArray = append(boxMapArray, arrayRow)
// 	}

// 	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 	inputString := day15data[1]
// 	currentPos := [2]int{0, 0}

// 	for i := range boxMapArray {
// 		for j := range boxMapArray[0] {
// 			if boxMapArray[i][j] == "@" {
// 				currentPos[0] = i
// 				currentPos[1] = j
// 				boxMapArray[i][j] = "."
// 				break
// 			}
// 		}
// 		if currentPos[0] != 0 {
// 			break
// 		}
// 	}

// 	for _, command := range inputString {
// 		switch command {
// 		case '^':
// 			if moveable(boxMapArray, currentPos, directions[UP]) {
// 				boxMapArray = push(boxMapArray, currentPos, directions[UP])
// 				currentPos[0] += directions[UP][0]
// 				currentPos[1] += directions[UP][1]
// 			}
// 		case 'v':
// 			if moveable(boxMapArray, currentPos, directions[DOWN]) {
// 				boxMapArray = push(boxMapArray, currentPos, directions[DOWN])
// 				currentPos[0] += directions[DOWN][0]
// 				currentPos[1] += directions[DOWN][1]
// 			}
// 		case '<':
// 			if moveable(boxMapArray, currentPos, directions[LEFT]) {
// 				boxMapArray = push(boxMapArray, currentPos, directions[LEFT])
// 				currentPos[0] += directions[LEFT][0]
// 				currentPos[1] += directions[LEFT][1]
// 			}
// 		case '>':
// 			if moveable(boxMapArray, currentPos, directions[RIGHT]) {
// 				boxMapArray = push(boxMapArray, currentPos, directions[RIGHT])
// 				currentPos[0] += directions[RIGHT][0]
// 				currentPos[1] += directions[RIGHT][1]
// 			}
// 		}
// 	}

// 	for i := range boxMapArray {
// 		for j := range boxMapArray {
// 			if boxMapArray[i][j] == "O" {
// 				part1Answer += (100*i + j)
// 			}
// 		}
// 	}

// 	fmt.Println(part1Answer)
// 	fmt.Println(part2Answer)
// 	fmt.Println(time.Since(start))
// }

// const (
// 	UP    int = 0
// 	DOWN  int = 1
// 	LEFT  int = 2
// 	RIGHT int = 3
// )

// func moveable(boxMap [][]string, current [2]int, direction [2]int) bool {
// 	currentY := current[0]
// 	currentX := current[1]
// 	for {
// 		currentY += direction[0]
// 		currentX += direction[1]
// 		switch boxMap[currentY][currentX] {
// 		case "O":
// 			continue
// 		case "#":
// 			return false
// 		case ".":
// 			return true
// 		default:
// 			panic("unrecognized escape character")
// 		}
// 	}
// }

// func push(boxMap [][]string, current [2]int, direction [2]int) [][]string {
// 	firstBox := [2]int{current[0] + direction[0], current[1] + direction[1]}
// 	if boxMap[firstBox[0]][firstBox[1]] == "." {
// 		return boxMap
// 	}
// 	currentY := current[0]
// 	currentX := current[1]

// 	for {
// 		currentY += direction[0]
// 		currentX += direction[1]
// 		if boxMap[currentY][currentX] == "." {
// 			boxMap[currentY][currentX] = "O"

// 			boxMap[firstBox[0]][firstBox[1]] = "."
// 			return boxMap
// 		} else if boxMap[currentY][currentX] == "O" {
// 			// fmt.Println("boing")
// 			continue
// 		} else {
// 			panic("unrecognized character")
// 		}
// 	}
// 	panic("should never exit loop")
// }
