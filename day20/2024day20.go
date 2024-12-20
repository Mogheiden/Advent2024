package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day20.txt")
	day20data := strings.Split(string(bytesread), "\n")
	part1Answer := 0
	part2Answer := 0

	// fmt.Println(len(newMap))
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}
