package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	data := parseInputReports("day3.txt")
	matches := getMulCalls(data)
	enabled := true
	part1answer := 0
	part2answer := 0
	for _, match := range matches {
		if match[0] != 'd' {
			part1answer += processMatch(match)
		}
		if match == "don't()" {
			enabled = false
			continue
		}
		if match == "do()" {
			enabled = true
			continue
		}
		if enabled {
			part2answer += processMatch(match)
		}
	}
	fmt.Println(part1answer)
	fmt.Println(part2answer)
	fmt.Println(time.Since(start))
}

func parseInputReports(raw string) string {
	file, _ := os.Open(raw)

	scanner := bufio.NewReader(file)
	var input strings.Builder
	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				input.WriteString(line)
				break
			}
			fmt.Println("Error")
			return ""
		}
		input.WriteString(line)
	}
	return input.String()
}

func getMulCalls(input string) []string {
	pattern := `mul\(\d+,\d+\)|do\(\)|don\'t\(\)`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllString(input, -1)
	return matches
}

func processMatch(input string) int {
	parts := strings.Split(input, ",")
	part1, _ := strconv.Atoi(parts[0][4:])
	part2, _ := strconv.Atoi(parts[1][:len(parts[1])-1])
	return part1 * part2
}
