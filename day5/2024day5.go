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
	for _, instruction := range instructions{
		pageList := strings.Split(instruction, ",")
		bestIndex := len(pageList)
		midIndex := bestIndex/2
		valid := true

		for index, page := range pageList{
			predecessors, _ := orderManual[page]

			// fmt.Println(instruction, predecessors)
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
				break
			}
			// // fmt.Println(instruction, bestIndex,orderManual[instruction])
			// if bestIndex == midIndex{
			// 	val, _ := strconv.Atoi(instruction)
			// 	// fmt.Println(val)
			// 	part1Answer += val
			// }
			// bestIndex = len(instructionList)
		}
		if valid{
			val, _ := strconv.Atoi(pageList[midIndex])
			part1Answer += val
		}

	}
	fmt.Println(part1Answer)
	fmt.Println(time.Since(start))
}