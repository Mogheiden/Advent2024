package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day9.txt")
	day9data := string(bytesread)
	filledMap := make(map[int][]int)
	part2FilledMap := make(map[int][]int)
	emptyMap := make(map[int][]int)
	var emptySpaces []int
	pc := 0
	fileID := -1
	emptyBlockId := 0

	for i := range day9data {
		length, _ := strconv.Atoi(string(day9data[i]))
		if i%2 == 0 {
			fileID++
			for range length {
				filledMap[fileID] = append(filledMap[fileID], pc)
				part2FilledMap[fileID] = append(part2FilledMap[fileID], pc)
				pc++
			}
		} else {
			for range length {
				emptySpaces = append(emptySpaces, pc)
				emptyMap[emptyBlockId] = append(emptyMap[emptyBlockId], pc)
				pc++
			}
			emptyBlockId++
		}
	}

	maxLen := pc
	pc = 0
	for {
		for j := len(filledMap[fileID]) - 1; j >= 0; j-- {
			maxLen = filledMap[fileID][j]
			if emptySpaces[pc] >= maxLen {
				break
			}
			filledMap[fileID][j] = emptySpaces[pc]
			pc++
			if pc >= len(emptySpaces) {
				break
			}
		}
		if pc >= len(emptySpaces) || emptySpaces[pc] >= maxLen {
			break
		}
		fileID--
	}

	emptyMapSize := len(emptyMap)

	for i := len(part2FilledMap) - 1; i >= 0; i-- {
		currentBlock := part2FilledMap[i]
		currentBlockLength := len(currentBlock)
		for j := 0; j < emptyMapSize; j++ {
			currentSpace := emptyMap[j]
			if currentBlockLength > len(currentSpace) {
				continue
			}
			if currentBlock[0] < currentSpace[0] {
				break
			}
			for k := 0; k < currentBlockLength; k++ {
				currentBlock[k] = currentSpace[k]
			}
			emptyMap[j] = currentSpace[len(currentBlock):]
		}
	}

	part1Answer := 0
	part2Answer := 0

	for key, value := range filledMap {
		for _, space := range value {
			part1Answer += key * space
		}
	}

	for key, value := range part2FilledMap {
		for _, space := range value {
			part2Answer += key * space
		}
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}
