package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day23.txt")
	day23data := strings.Split(string(bytesread), "\n")
	part1Answer := 0

	graph := make(map[string][]string)
	var ordered []string

	for _, line := range day23data {
		keyVal := strings.Split(line, "-")
		graph[keyVal[0]] = append(graph[keyVal[0]], keyVal[1])
		graph[keyVal[1]] = append(graph[keyVal[1]], keyVal[0])
	}

	for key := range graph {
		ordered = append(ordered, key)
	}

	var visited []string
	var triangles [][3]string
	for i := 0; i < len(ordered)-1; i++ {
		var visitedInner []string
		outerSliceVal := ordered[i]
		outerSlice := graph[ordered[i]]
		for j := i + 1; j < len(ordered); j++ {
			innerSlice := graph[ordered[j]]
			if !slices.Contains(innerSlice, outerSliceVal) {
				continue
			} else {
				for _, vertex := range innerSlice {
					if slices.Contains(visited, vertex) || slices.Contains(visitedInner, vertex) {
						continue
					}
					if slices.Contains(outerSlice, vertex) {
						triangle := [3]string{outerSliceVal, ordered[j], vertex}
						triangles = append(triangles, triangle)
					}
				}
			}
			visitedInner = append(visitedInner, ordered[j])
		}
		visited = append(visited, ordered[i])
	}

	// fmt.Println(triangles)

	for _, triangle := range triangles {
		for _, vertex := range triangle {
			if vertex[0] == 't' {
				part1Answer++
				break
			}
		}
	}
	// fmt.Println(ordered)
	var cliques [][]string

	graph2 := make(map[string]map[string]bool)

	for _, line := range day23data {
		keyVal := strings.Split(line, "-")
		if graph2[keyVal[0]] == nil {
			graph2[keyVal[0]] = make(map[string]bool)
		}
		if graph2[keyVal[1]] == nil {
			graph2[keyVal[1]] = make(map[string]bool)
		}
		graph2[keyVal[0]][keyVal[1]] = true
		graph2[keyVal[1]][keyVal[0]] = true
	}
	P := make(map[string]bool)
	X := make(map[string]bool)
	R := make(map[string]bool)

	for v := range graph2 {
		P[v] = true
	}

	bronKerbosch(R, P, X, graph2, &cliques)

	highestLen := 0
	var maxClique []string

	for _, clique := range cliques {
		if len(clique) > highestLen {
			highestLen = len(clique)
			maxClique = clique
		}
	}
	slices.Sort(maxClique)
	part2Answer := strings.Join(maxClique, ",")
	fmt.Println(maxClique)

	// fmt.Println(len(newMap))
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func bronKerbosch(R, P, X map[string]bool, adj map[string]map[string]bool, cliques *[][]string) {
	if len(P) == 0 && len(X) == 0 {
		clique := make([]string, 0, len(R))
		for vertex := range R {
			clique = append(clique, vertex)
		}
		*cliques = append(*cliques, clique)
	}
	for vertex := range P {
		newR := make(map[string]bool)
		for v := range R {
			newR[v] = true
		}
		newR[vertex] = true

		newP := make(map[string]bool)
		for v := range P {
			if adj[vertex][v] {
				newP[v] = true
			}
		}

		newX := make(map[string]bool)
		for v := range X {
			if adj[vertex][v] {
				newX[v] = true
			}
		}

		bronKerbosch(newR, newP, newX, adj, cliques)
		delete(P, vertex)
		X[vertex] = true
	}
}
