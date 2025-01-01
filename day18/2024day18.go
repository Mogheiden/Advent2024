package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const mapSize = 71

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day18.txt")
	day18data := strings.Split(string(bytesread), "\n")

	var bitMap [mapSize][mapSize]string

	for i := range mapSize {
		for j := range mapSize {
			bitMap[i][j] = "."
		}
	}

	i := 0

	for i < 1024 {
		brickCoord := parseBrick(day18data[i])
		bitMap[brickCoord[0]][brickCoord[1]] = "#"
		i++
	}

	startPos := [2]int{0, 0}
	endPos := [2]int{mapSize - 1, mapSize - 1}

	part1Answer := 0
	part2Answer := 0

	part1Answer = aStar(bitMap, startPos, endPos)

	path := aStar2(bitMap, startPos, endPos)

	for len(path) > 3 {
		brickCoord := parseBrick(day18data[i])
		bitMap[brickCoord[0]][brickCoord[1]] = "#"
		if path[brickCoord] {
			path = aStar2(bitMap, startPos, endPos)
		}
		i++
	}
	fmt.Println(day18data[i-1])
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func aStar(trailMap [mapSize][mapSize]string, startingCoord [2]int, endingCoord [2]int) int {
	visitedCoords := make(map[[2]int]bool)
	pq := NewPriorityQueue()

	heap.Push(pq, &coordPoint{coord: startingCoord, score: 0, distance: 0})
	neighbours := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	backtrack := make(map[[2]int][2]int)
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*coordPoint)
		distance := current.distance
		coord := current.coord
		for i := range neighbours {
			neighbour := [2]int{coord[0] + neighbours[i][0], coord[1] + neighbours[i][1]}
			if !outofBounds(neighbour) || visitedCoords[neighbour] {
				continue
			}
			if trailMap[neighbour[0]][neighbour[1]] == "#" {
				continue
			}
			if tupleEquals(neighbour, endingCoord) {
				backtrack[neighbour] = coord
				return distance + 1
			}
			heap.Push(pq, &coordPoint{coord: neighbour, score: manhattanDistance(neighbour, endingCoord) + distance, distance: distance + 1})
			backtrack[neighbour] = coord
		}
		visitedCoords[coord] = true
	}
	return 0
}

func aStar2(trailMap [mapSize][mapSize]string, startingCoord [2]int, endingCoord [2]int) map[[2]int]bool {
	visitedCoords := make(map[[2]int]bool)
	path := make(map[[2]int]bool)
	pq := NewPriorityQueue()

	heap.Push(pq, &coordPoint{coord: startingCoord, score: 0, distance: 0})
	neighbours := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	backtrack := make(map[[2]int][2]int)
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*coordPoint)
		distance := current.distance
		coord := current.coord
		for i := range neighbours {
			neighbour := [2]int{coord[0] + neighbours[i][0], coord[1] + neighbours[i][1]}
			if !outofBounds(neighbour) || visitedCoords[neighbour] {
				continue
			}
			if trailMap[neighbour[0]][neighbour[1]] == "#" {
				continue
			}
			if tupleEquals(neighbour, endingCoord) {
				backtrack[neighbour] = coord
				path = unspool(backtrack, coord)
				return path
			}
			heap.Push(pq, &coordPoint{coord: neighbour, score: manhattanDistance(neighbour, endingCoord) + distance, distance: distance + 1})
			backtrack[neighbour] = coord
		}
		visitedCoords[coord] = true
	}
	return path
}

func unspool(backtrack map[[2]int][2]int, current [2]int) map[[2]int]bool {
	current, available := backtrack[current]
	pathMap := make(map[[2]int]bool)
	for available {
		current, available = backtrack[current]
		pathMap[current] = true
	}
	return pathMap
}

func tupleEquals(first [2]int, second [2]int) bool {
	return first[0] == second[0] && first[1] == second[1]
}

func outofBounds(coord [2]int) bool {
	return coord[0] >= 0 && coord[1] >= 0 && coord[0] < mapSize && coord[1] < mapSize
}

type coordPoint struct {
	coord    [2]int
	score    int
	distance int
	index    int
}

type priorityQueue []*coordPoint

// Len is the number of elements in the collection.
func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*coordPoint)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the highest priority element from the queue.
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// NewPriorityQueue initializes and returns a new priority queue.
func NewPriorityQueue() *priorityQueue {
	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

func manhattanDistance(first [2]int, second [2]int) int {
	yDiff := first[0] - second[0]
	if yDiff < 1 {
		yDiff *= -1
	}
	xDiff := first[1] - second[1]
	if xDiff < 1 {
		xDiff *= -1
	}
	return xDiff + yDiff
}

func parseBrick(input string) [2]int {
	line := input
	coord := strings.Split(line, ",")
	coordY, _ := strconv.Atoi(coord[0])
	coordX, _ := strconv.Atoi(coord[1])
	return [2]int{coordX, coordY}
}
