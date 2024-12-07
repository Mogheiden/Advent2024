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
	bytesread, _ := os.ReadFile("day7.txt")
	day7data := strings.Split(string(bytesread), "\n")

	part1Answer := 0
	part2Answer := 0
	for _, line := range day7data {
		splitInput := strings.Split(line, ": ")
		total, _ := strconv.Atoi(splitInput[0])
		var operands []int
		splitOperands := strings.Split(splitInput[1], " ")
		for _, operand := range splitOperands {
			numOperand, _ := strconv.Atoi(operand)
			operands = append(operands, numOperand)
		}
		success := recursiveSolution(0, total, len(operands)-1, operands, 0)
		if success {
			part1Answer = total + part1Answer
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func recursiveSolution(depth int, goal int, maxDepth int, graph []int, cumulation int) bool {
	if depth == maxDepth && (cumulation+graph[maxDepth] == goal || cumulation*graph[maxDepth] == goal) {
		return true
	}

	if depth == maxDepth && !(cumulation+graph[maxDepth] == goal || cumulation*graph[maxDepth] == goal) {
		return false
	}

	if recursiveSolution(depth+1, goal, maxDepth, graph, cumulation+graph[depth]) {
		return true
	}
	return recursiveSolution(depth+1, goal, maxDepth, graph, cumulation*graph[depth])
}

func recursiveSolutionWithConcat(depth int, goal int, maxDepth int, graph []int, cumulation int) bool {
	if depth == maxDepth && (cumulation+graph[maxDepth] == goal || cumulation*graph[maxDepth] == goal) {
		return true
	}

	if depth == maxDepth && !(cumulation+graph[maxDepth] == goal || cumulation*graph[maxDepth] == goal) {
		return false
	}

	if recursiveSolution(depth+1, goal, maxDepth, graph, cumulation+graph[depth]) {
		return true
	}
	return recursiveSolution(depth+1, goal, maxDepth, graph, cumulation*graph[depth])
}
