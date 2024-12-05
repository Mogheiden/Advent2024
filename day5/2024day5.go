package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main(){
	start := time.Now()
	bytesread, _ := os.ReadFile("day5.txt")
	day5data := strings.Split(string(bytesread), "\n\n")
	pagePairs := strings.Split(day5data[0],"\n")
	instructions := strings.Split(day5data[1],"\n")
	orderManual := make(map[string][]string)

	for _, pair := range pagePairs{
		splitPair := strings.Split(pair, "|")
		orderManual[splitPair[0]] = append(orderManual[splitPair[0]], splitPair[1])
	}
	part1Answer := 0
	part2Answer := 0
	var invalidIndexes []int
	for i, instruction := range instructions{
		pageList := strings.Split(instruction, ",")
		bestIndex := len(pageList)
		midIndex := bestIndex/2
		valid := true

		for index, page := range pageList{
			predecessors, _ := orderManual[page]

			for _, predecessor := range predecessors{
				if slices.Contains(pageList, predecessor){
					// bestIndex --
					var predIndex int
					for othIndex, v := range pageList{
						if v == predecessor{
							predIndex = othIndex
							break
						}
					}
					if predIndex < index{
						// fmt.Println(predIndex, predecessor, index, page)
						valid = false
						break
					}
				}
			}
			if !valid{
				invalidIndexes = append(invalidIndexes,i)
				break
			}
		}
		if valid{
			val, _ := strconv.Atoi(pageList[midIndex])
			part1Answer += val
		}

	}
	for _, index := range invalidIndexes{
		pageblob := instructions[index]
		pageList := strings.Split(pageblob, ",")
		bestIndex := len(pageList)
		midIndex := bestIndex/2

		for _, page := range pageList{
			predecessors, _ := orderManual[page]

			// fmt.Println(instruction, predecessors)
			for _, predecessor := range predecessors{
				if slices.Contains(pageList, predecessor){
					bestIndex --
				}
			}
			// fmt.Println(instruction, bestIndex,orderManual[instruction])
			if bestIndex == midIndex + 1{
				val, _ := strconv.Atoi(page)
				// fmt.Println(val, pageList) 
				part2Answer += val
			}
			bestIndex = len(pageList)
		}

	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)	
	fmt.Println(time.Since(start))
}