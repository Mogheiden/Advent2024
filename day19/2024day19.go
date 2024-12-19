package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day19.txt")
	day19data := strings.Split(string(bytesread), "\n\n")
	stripeList := strings.Split(day19data[0], ", ")
	towels := strings.Split(day19data[1], "\n")
	stripeDict := make(map[string]bool)
	for _, stripe := range stripeList {
		stripeDict[stripe] = true
	}

	part1Answer := 0
	part2Answer := 0

	// for _, towel := range towels {
	// 	if recursiveStringBuilder(towel, 0, 0, stripeDict) {
	// 		part1Answer++
	// 	}
	// 	// fmt.Println(i)
	// }

	maxLen := 0

	for key := range stripeDict {
		if len(key) > maxLen {
			maxLen = len(key)
		}
	}

	// lanternFishDict := makeSubstringDict(towels[0], maxLen, stripeDict)
	// fmt.Println(lanternFishDict)
	for _, towel := range towels {
		lanternFishDict := makeSubstringDict(towel, maxLen, stripeDict)
		// fmt.Println(lanternFishDict)
		if len(lanternFishDict[0]) == 0 {
			continue
		}
		// fmt.Println(lanternFishDict)
		// increment := recursiveLanternFish(0, lanternFishDict, len(towel))
		increment := 0

		pathsMap := make(map[int][]int)
		var queue [][2]int
		for key, value := range lanternFishDict {
			if sort.SearchInts(value, len(towel)) != len(value) {
				next := [2]int{key, 1}
				queue = append(queue, next)
			}
		}
		if increment > 0 {
			part1Answer++
		}
		// fmt.Println(increment, len(towel))
		part2Answer += increment
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func recursiveStringBuilder(towel string, depth int, possibleSolutions int, stripes map[string]bool) bool {
	stringBuilder := string(towel[depth])
	for depth < len(towel)-1 {
		if stripes[stringBuilder] {
			if recursiveStringBuilder(towel, depth+1, possibleSolutions, stripes) {
				return true
			}
		}
		depth++
		stringBuilder += string(towel[depth])
	}
	if stripes[stringBuilder] {
		return true
	} else {
		return false
	}
}

func makeSubstringDict(towel string, maxLen int, stripes map[string]bool) map[int][]int {
	subStringDict := make(map[int][]int)
	for i := 0; i <= maxLen; i++ {
		for j := 0; j < len(towel)-i+1; j++ {
			substring := towel[j : j+i]
			// fmt.Println(substring)
			if stripes[substring] {
				subStringDict[j] = append(subStringDict[j], j+i)
			}
		}
	}
	return subStringDict
}

// func recursiveLanternFish(depth int, fishDict map[int][]int, maxLength int) int {
// 	currentSlice := fishDict[depth]
// 	returnVal := 0
// 	for i, val := range currentSlice {
// 		if val >= maxLength {
// 			returnVal++
// 		} else if len(fishDict[val]) < 1 {
// 			continue
// 		} else {
// 			returnVal += recursiveLanternFish(currentSlice[i], fishDict, maxLength)
// 		}
// 	}
// 	return returnVal
// }
