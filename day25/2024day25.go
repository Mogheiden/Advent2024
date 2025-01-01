package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day25.txt")
	day25data := strings.Split(string(bytesread), "\n\n")

	locks := make(map[int][5]int)
	keys := make(map[int][5]int)

	for i, schematic := range day25data {
		key := [5]int{0}
		rows := strings.Split(schematic, "\n")

		lock := false
		if rows[0][0] == '#' {
			lock = true
		}
		for i := range 5 {
			for j := range 7 {
				if rows[j][i] == '#' {
					key[i]++
				}
			}
			key[i]--
		}

		if lock {
			locks[i] = key
		} else {
			keys[i] = key
		}
	}

	part1Answer := 0
	part2Answer := 0

	for _, keyValue := range keys {
		for _, lockValue := range locks {
			if compatible(keyValue, lockValue) {
				part1Answer++
			}
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func compatible(key [5]int, lock [5]int) bool {
	for i := range 5 {
		if key[i]+lock[i] >= 6 {
			return false
		}
	}
	// fmt.Println(lock, key)
	return true
}
