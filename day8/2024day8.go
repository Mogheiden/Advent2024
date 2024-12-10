package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day8.txt")
	day8data := strings.Split(string(bytesread), "\n")

	mapHeight := len(day8data)
	mapWidth := len(day8data[0])

	antennaMap := make(map[byte][][2]int)
	antiNodes := make(map[[2]int]bool)
	part2AntiNodes := make(map[[2]int]bool)

	for i := range day8data {
		for j := range day8data[0] {
			if day8data[i][j] != '.' {
				key := day8data[i][j]
				value := [2]int{i, j}
				antennaMap[key] = append(antennaMap[key], value)
			}
		}
	}
	for _, list := range antennaMap {
		if len(list) > 1 {
			for i := range list {
				part2AntiNodes[list[i]] = true
			}
		}
	}

	for _, antennaList := range antennaMap {
		for i := range antennaList {
			for j := range antennaList {
				if antennaList[i] == antennaList[j] {
					continue
				}
				distance := distanceFinder(antennaList[i], antennaList[j])
				candidateNode := [2]int{antennaList[i][0] + distance[0], antennaList[i][1] + distance[1]}
				// fmt.Println(antiNodeCandidate)
				// fmt.Println(antennaList[i], antennaList[j])
				if isValid(candidateNode, mapHeight, mapWidth) {
					antiNodes[candidateNode] = true
					part2AntiNodes[candidateNode] = true
				}
				for isValid(candidateNode, mapHeight, mapWidth) {
					part2AntiNodes[candidateNode] = true
					candidateNode = [2]int{candidateNode[0] + distance[0], candidateNode[1] + distance[1]}
				}
			}
		}
	}

	// fmt.Println(antiNodes)

	part1Answer := len(antiNodes)
	part2Answer := len(part2AntiNodes)

	// fmt.Println(part2AntiNodes)

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func distanceFinder(antenna [2]int, antennaPair [2]int) [2]int {
	distance := [2]int{antenna[0] - antennaPair[0], antenna[1] - antennaPair[1]}
	return distance
}

func isValid(antinodeCandidate [2]int, height int, width int) bool {
	if antinodeCandidate[0] < 0 || antinodeCandidate[1] < 0 || antinodeCandidate[0] >= height || antinodeCandidate[1] >= width {
		// fmt.Println(antinodeCandidate)
		return false
	}
	return true
}
