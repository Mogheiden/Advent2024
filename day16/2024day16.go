package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day16.txt")
	day16data := strings.Split(string(bytesread), "\n")

	part1Answer := 0
	part2Answer := 0

	var startPoint [2]int
	var finishPoint [2]int
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < len(day16data); i++ {
		for j := 0; j < len(day16data[0]); j++ {
			if day16data[i][j] == 'E' {
				finishPoint = [2]int{i, j}
			}
			if day16data[i][j] == 'S' {
				startPoint = [2]int{i, j}
			}
		}
	}

	var convertedMap [][]string

	for i := range day16data {
		row := strings.Split(day16data[i], "")
		convertedMap = append(convertedMap, row)
	}

	visitedNodes := make(map[[2]int][]*coordPoint)
	pq := NewPriorityQueue()

	heap.Push(pq, &coordPoint{direction: RIGHT, coord: startPoint, score: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*coordPoint)
		currentCoord := current.coord
		_, contains := visitedNodes[currentCoord]
		if contains {
			continue
		}
		if tupleEquals(currentCoord, finishPoint) {
			part1Answer = current.score
		}
		for i := range directions {
			step := tupleAddition(currentCoord, directions[i])
			if day16data[step[0]][step[1]] == '#' {
				continue
			}
			if i == current.direction {
				heap.Push(pq, &coordPoint{direction: i, coord: step, score: current.score + 1})
			} else {
				heap.Push(pq, &coordPoint{direction: i, coord: step, score: current.score + 1001})
			}
		}
		visitedNodes[currentCoord] = append(visitedNodes[currentCoord], current)
	}
	goldenPaths := findAllOptimalPaths(day16data, part1Answer, coordPoint{direction: RIGHT, coord: startPoint, score: 0}, finishPoint)

	goldCoords := make(map[[2]int]bool)
	for _, path := range goldenPaths {
		for _, coord := range path {
			goldCoords[coord] = true
		}
	}

	part2Answer = len(goldCoords) + 1

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func findAllOptimalPaths(m []string, targetScore int, start coordPoint, end [2]int) [][][2]int {
	queue := []coordPoint{start}
	visited := make(map[[3]int]int)
	var paths [][][2]int
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentCoord := current.coord
		if current.score > targetScore {
			continue
		}

		key := [3]int{currentCoord[0], currentCoord[1], current.direction}

		if score, exists := visited[key]; exists && score < current.score {
			continue
		}
		visited[key] = current.score

		if tupleEquals(end, currentCoord) && current.score == targetScore {
			paths = append(paths, current.path)
			continue
		}
		nextPos := tupleAddition(currentCoord, directions[current.direction])
		if m[nextPos[0]][nextPos[1]] != '#' {
			newPath := make([][2]int, len(current.path))
			copy(newPath, current.path)
			queue = append(queue, coordPoint{direction: current.direction, coord: nextPos, score: current.score + 1, path: append(newPath, nextPos)})
		}
		for i := range 4 {
			queue = append(queue, coordPoint{
				coord:     currentCoord,
				direction: i,
				score:     current.score + 1000,
				path:      current.path,
			})
		}
	}
	return paths
}

const (
	UP    int = 0
	DOWN  int = 1
	LEFT  int = 2
	RIGHT int = 3
)

type coordPoint struct {
	direction int
	coord     [2]int
	score     int
	index     int
	path      [][2]int
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

func tupleEquals(first [2]int, second [2]int) bool {
	return first[0] == second[0] && first[1] == second[1]
}

func tupleAddition(first [2]int, second [2]int) [2]int {
	return [2]int{first[0] + second[0], first[1] + second[1]}
}
