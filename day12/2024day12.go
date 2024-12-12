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
	bytesread, _ := os.ReadFile("day12.txt")
	day12data := strings.Split(string(bytesread), "\n")
	answer := 0
	// part2Answer := 0
	visitedNodes := make(map[byte][][2]int)
	part2 := true

	for i := range day12data {
		for j := range day12data {
			nodeVal := day12data[i][j]
			nodesOfVal := visitedNodes[nodeVal]
			if contains(nodesOfVal, i, j) {
				continue
			} else {
				answer += floodfill(visitedNodes, day12data, nodeVal, i, j, part2)
			}
		}
	}
	fmt.Println(answer)
	fmt.Println(time.Since(start))
}

func contains(list [][2]int, y int, x int) bool {
	for _, val := range list {
		if val[0] == y && val[1] == x {
			return true
		}
	}
	return false
}

func floodfill(visitedNodes map[byte][][2]int, farmMap []string, nodeVal byte, y int, x int, part2 bool) int {
	queue := [][2]int{{y, x}}
	neighbours := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	plotSize := 1
	visitedNodes[nodeVal] = append(visitedNodes[nodeVal], queue[0])
	inside := make(map[[2]int]bool)
	outside := make(map[[2]int]bool)
	inside[queue[0]] = true
	edges := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for i := range neighbours {
			neighbour := [2]int{current[0] + neighbours[i][0], current[1] + neighbours[i][1]}
			if inside[neighbour] {
				continue
			}
			if neighbour[0] < 0 || neighbour[1] < 0 || neighbour[0] >= len(farmMap) || neighbour[1] >= len(farmMap[0]) {
				edges++
				outside[neighbour] = true
				continue
			}
			if farmMap[neighbour[0]][neighbour[1]] != nodeVal {
				edges++
				outside[neighbour] = true
				continue
			} else {
				plotSize++
				inside[neighbour] = true
				queue = append(queue, neighbour)
				visitedNodes[nodeVal] = append(visitedNodes[nodeVal], neighbour)
			}
		}
	}
	// fmt.Println(string(nodeVal))
	trace := edgeCounter(inside, outside)
	// fmt.Println(plotSize, trace)
	if part2 {
		return plotSize * trace
	}
	return plotSize * edges
}

func edgeCounter(inside map[[2]int]bool, outside map[[2]int]bool) int {

	var right [][2]int
	var left [][2]int
	var up [][2]int
	var down [][2]int
	edges := 0

	for key, _ := range inside {
		coord := [2]int{key[0] - 1, key[1]}
		if outside[coord] {
			up = append(up, coord)
		}
		coord = [2]int{key[0], key[1] - 1}
		if outside[coord] {
			left = append(left, coord)
		}
		coord = [2]int{key[0] + 1, key[1]}
		if outside[coord] {
			down = append(down, coord)
		}
		coord = [2]int{key[0], key[1] + 1}
		if outside[coord] {
			right = append(right, coord)
		}
	}
	sort.Slice(right, func(i, j int) bool {
		if right[i][1] == right[j][1] {
			return right[i][0] < right[j][0]
		} else {
			return right[i][1] < right[j][1]
		}
	})
	sort.Slice(left, func(i, j int) bool {
		if left[i][1] == left[j][1] {
			return left[i][0] < left[j][0]
		} else {
			return left[i][1] < left[j][1]
		}
	})
	sort.Slice(up, func(i, j int) bool {
		if up[i][0] == up[j][0] {
			return up[i][1] < up[j][1]
		} else {
			return up[i][0] < up[j][0]
		}
	})
	sort.Slice(down, func(i, j int) bool {
		if down[i][0] == down[j][0] {
			return down[i][1] < down[j][1]
		} else {
			return down[i][0] < down[j][0]
		}
	})
	// fmt.Println(sortedEdges(up, true), sortedEdges(down, true), sortedEdges(left, false), sortedEdges(right, false))
	edges += sortedEdges(up, true)
	edges += sortedEdges(down, true)
	edges += sortedEdges(left, false)
	edges += sortedEdges(right, false)
	return edges
}

func sortedEdges(edgeArray [][2]int, up bool) int {
	// fmt.Println(edgeArray)
	edges := 1
	if up {
		for i := range len(edgeArray) - 1 {
			// fmt.Println(edgeArray[i], edgeArray[i+1])
			if edgeArray[i+1][1] != edgeArray[i][1]+1 || edgeArray[i][0] != edgeArray[i+1][0] {
				edges++
			}
		}
	} else {
		for i := range len(edgeArray) - 1 {
			if edgeArray[i+1][0] != edgeArray[i][0]+1 || edgeArray[i][1] != edgeArray[i+1][1] {
				edges++
			}
		}
	}
	// fmt.Println(edges)
	return edges
}
