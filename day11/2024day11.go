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
	bytesread, _ := os.ReadFile("day11.txt")
	day11data := strings.Split(string(bytesread), " ")
	part1Answer := 0
	part2Answer := 0
	cachedPossibilites := make(map[string][]string)
	oldMap := make(map[string]int)
	newMap := make(map[string]int)
	for _, i := range day11data {
		part1Answer += recursiveSolution(0, 25, i, cachedPossibilites)
		oldMap[i] += 1
	}
	for range 75 {
		newMap = make(map[string]int)
		for key, value := range oldMap {
			// fmt.Println(key, value)
			if len(key)%2 == 0 {
				left, _ := strconv.Atoi(key[len(key)/2:])
				leftstr := strconv.Itoa(left)
				right, _ := strconv.Atoi(key[:len(key)/2])
				rightstr := strconv.Itoa(right)
				newMap[leftstr] += value
				newMap[rightstr] += value
				continue
			}
			if key == "0" {
				newMap["1"] += value
			} else {
				number, _ := strconv.Atoi(key)
				number *= 2024
				newMap[strconv.Itoa(number)] += value
			}
		}
		oldMap = newMap
	}

	for _, value := range newMap {
		part2Answer += value
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func recursiveSolution(depth int, maxDepth int, number string, cachedPossibilites map[string][]string) int {
	if depth == maxDepth {
		return 1
	}
	value, in := cachedPossibilites[number]

	if in {
		if len(value) == 2 {
			return recursiveSolution(depth+1, maxDepth, value[0], cachedPossibilites) + recursiveSolution(depth+1, maxDepth, value[1], cachedPossibilites)
		} else {
			return recursiveSolution(depth+1, maxDepth, value[0], cachedPossibilites)
		}
	}
	if len(number)%2 == 0 {
		left, _ := strconv.Atoi(number[len(number)/2:])
		leftstr := strconv.Itoa(left)
		right, _ := strconv.Atoi(number[:len(number)/2])
		rightstr := strconv.Itoa(right)
		cachedPossibilites[number] = append(cachedPossibilites[number], leftstr)
		cachedPossibilites[number] = append(cachedPossibilites[number], rightstr)
		return recursiveSolution(depth+1, maxDepth, leftstr, cachedPossibilites) + recursiveSolution(depth+1, maxDepth, rightstr, cachedPossibilites)
	} else {
		number, _ := strconv.Atoi(number)
		if number == 0 {
			return recursiveSolution(depth+1, maxDepth, "1", cachedPossibilites)
		} else {
			cachedPossibilites[strconv.Itoa(number)] = append(cachedPossibilites[strconv.Itoa(number)], strconv.Itoa(number*2024))
			return recursiveSolution(depth+1, maxDepth, strconv.Itoa(number*2024), cachedPossibilites)
		}
	}
}
