package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func countWaysToCreateDesign(patterns []string, design string) int {
	n := len(design)
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for _, pattern := range patterns {
			patternLength := len(pattern)
			if i >= patternLength && design[i-patternLength:i] == pattern {
				dp[i] += dp[i-patternLength]
			}
		}
	}

	return dp[n]
}

func getDesignCounts(patterns []string, designs []string) map[string]int {
	designCounts := make(map[string]int)
	for _, design := range designs {
		count := countWaysToCreateDesign(patterns, design)
		designCounts[design] = count
	}
	return designCounts
}

func main() {
	start := time.Now()
	file, err := os.Open("day19.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	patterns := make([]string, 0)
	designs := make([]string, 0)

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if i == 0 {
			patterns = strings.Split(line, ",")
			for j, pattern := range patterns {
				patterns[j] = strings.TrimSpace(pattern)
			}
			i++
		} else {
			designs = append(designs, line)
		}
	}

	designCounts := getDesignCounts(patterns, designs)
	totalPart1 := 0
	totalPart2 := 0
	for _, count := range designCounts {
		if count > 0 {
			totalPart1++
		}
		totalPart2 += count
	}
	fmt.Printf("Part 1: %d\n", totalPart1)
	fmt.Printf("Part 2: %d\n", totalPart2)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(time.Since(start))
}
