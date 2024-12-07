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
			part2Answer = total + part2Answer
		} else {
			success2 := recursiveSolutionWithConcatenation(0, total, len(operands)-2, operands, 0)
			if success2 {
				part2Answer = total + part2Answer
			}
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func recursiveSolution(depth int, goal int, maxDepth int, graph []int, cumulation int) bool {
	// fmt.Println(depth, maxDepth)
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

func recursiveSolutionWithConcatenation(depth int, goal int, maxDepth int, graph []int, cumulation int) bool {
	var concatVal int
	var plusVal int
	var multVal int

	if depth == 0 {
		concatVal = concatenate(graph[0], graph[1])
		plusVal = graph[0] + graph[1]
		multVal = graph[0] * graph[1]
	} else {
		concatVal = concatenate(cumulation, graph[depth+1])
		plusVal = cumulation + graph[depth+1]
		multVal = cumulation * graph[depth+1]
	}
	// fmt.Println(depth, concatVal)
	if depth == maxDepth && (plusVal == goal || multVal == goal || concatVal == goal) {
		return true
	}

	if depth == maxDepth && !(plusVal == goal || multVal == goal || concatVal == goal) {
		return false
	}

	if recursiveSolutionWithConcatenation(depth+1, goal, maxDepth, graph, plusVal) {
		return true
	}
	if recursiveSolutionWithConcatenation(depth+1, goal, maxDepth, graph, multVal) {
		return true
	}
	return recursiveSolutionWithConcatenation(depth+1, goal, maxDepth, graph, concatVal)
}

func concatenate(numOne int, numTwo int) int {
	combination, _ := strconv.Atoi(strconv.Itoa(numOne) + strconv.Itoa(numTwo))
	return combination
}
