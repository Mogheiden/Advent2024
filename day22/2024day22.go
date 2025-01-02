package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day22.txt")
	day22data := strings.Split(string(bytesread), "\n")
	part1Answer := 0
	part2Answer := 0
	ranges := make(map[int][]int)

	for _, secret := range day22data {
		visited := make(map[int]struct{})
		changes := []int{}
		seed, _ := strconv.Atoi(secret)
		for range 2000 {
			nextSeed := process(seed)
			changes = append(changes, (nextSeed%10)-(seed%10))
			if len(changes) == 4 {
				key := hashMaker(changes)
				if _, found := visited[key]; !found {
					if _, exists := ranges[key]; !exists {
						ranges[key] = []int{}
					}
					ranges[key] = append(ranges[key], nextSeed%10)
					visited[key] = struct{}{}
				}
				changes = changes[1:]
			}
			seed = nextSeed
		}
		part1Answer += int(seed)
	}
	maxSum := 0
	for _, rangeValues := range ranges {
		sum := 0
		for _, val := range rangeValues {
			sum += val
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	part2Answer = maxSum
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func process(seed int) int {
	seed = ((seed << 6) ^ seed) % 16777216
	seed = ((seed >> 5) ^ seed)
	seed = ((seed << 11) ^ seed) % 16777216
	return seed
}

func hashMaker(slice []int) int {
	value := float64(0)
	for i, val := range slice {
		value += float64(val) * math.Pow10(i)
	}
	return int(value)
}
