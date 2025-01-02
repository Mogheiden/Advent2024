package main

import (
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
	bytesread, _ := os.ReadFile("day19.txt")
	day19data := strings.Split(string(bytesread), "\n\n")
	patterns := strings.Split(day19data[0], ", ")
	designs := strings.Split(day19data[1], "\n")

	designCounts := getDesignCounts(patterns, designs)
	part1Answer := 0
	part2Answer := 0
	for _, count := range designCounts {
		if count > 0 {
			part1Answer++
		}
		part2Answer += count
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}
