package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day1.txt")
	day1data := strings.Split(string(bytesread), "\n")

	var leftCol []int
	var rightCol []int

	for _, pair := range day1data {
		leftRight := strings.Split(pair, "   ")
		leftNum, _ := strconv.Atoi(leftRight[0])
		rightNum, _ := strconv.Atoi(leftRight[1])
		leftCol = append(leftCol, leftNum)
		rightCol = append(rightCol, rightNum)
	}

	sort.Slice(leftCol, func(i, j int) bool { return leftCol[i] < leftCol[j] })
	sort.Slice(rightCol, func(i, j int) bool { return rightCol[i] < rightCol[j] })

	part1Answer := 0
	for index, _ := range leftCol {
		diff := (leftCol[index] - rightCol[index])
		if diff < 0 {
			part1Answer += diff * -1
		} else {
			part1Answer += diff
		}
	}
	fmt.Println(part1Answer)

	frequencyMap := make(map[int]int)
	part2Answer := 0

	for _, value := range leftCol {
		if frequencyMap[value] == 0 {
			occurences := 0
			for _, rightVals := range rightCol {
				if rightVals == value {
					occurences++
				}
			}
			if occurences == 0 {
				frequencyMap[value] = -1
				continue
			}
			frequencyMap[value] = value * occurences
		}
		part2Answer += frequencyMap[value]
	}

	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}
