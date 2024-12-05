package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day4.txt")
	day4data := strings.Split(string(bytesread), "\n")

	part1Answer := 0
	part2Answer := 0

	for i := 0; i < len(day4data); i++ {
		for j := 0; j < len(day4data[0]); j++ {
			if day4data[i][j] == 'm' || day4data[i][j] == 'a' {
				continue
			}
			var substring [4]byte
			if j < len(day4data[0])-3 {
				for k := 0; k < 4; k++ {
					substring[k] = day4data[i][j+k]
				}
				if joiner(substring) {
					part1Answer++
				}
			}
			if i < len(day4data)-3 {
				for k := 0; k < 4; k++ {
					substring[k] = day4data[i+k][j]
				}
				if joiner(substring) {
					part1Answer++
				}
			}

			var substring2 [4]byte
			if i < len(day4data)-3 && j < len(day4data[0])-3 {
				for k := 0; k < 4; k++ {
					substring[k] = day4data[i+k][j+k]
					substring2[k] = day4data[i+3-k][j+k]
				}
				if joiner(substring) {
					part1Answer++
				}
				if joiner(substring2) {
					part1Answer++
				}
			}
		}
	}
	// fmt.Println(part1Answer)

	for i := 0; i < len(day4data)-2; i++ {
		for j := 0; j < len(day4data[0])-2; j++ {
			if day4data[i][j] == 'x' || day4data[i][j] == 'a' || day4data[i][j+2] == 'x' || day4data[i][j+2] == 'a' {
				continue
			}

			var x1 [4]byte
			var x2 [4]byte

			for k := 0; k < 3; k++ {
				x1[k] = day4data[i+k][j+k]
				x2[k] = day4data[i+2-k][j+k]
			}
			if xmasjoin(x1) && xmasjoin(x2) {
				part2Answer++
			}
			// fmt.Println(i,j)
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func joiner(data [4]byte) bool {
	forward := "XMAS"
	backward := "SAMX"
	valid := true
	for i := 0; i < 4; i++ {
		if data[i] != forward[i] {
			valid = false
			break
		}
	}
	if valid {
		return true
	}
	valid = true
	for i := 0; i < 4; i++ {
		if data[i] != backward[i] {
			valid = false
			break
		}
	}
	if valid {
		return true
	}
	return false
}

func xmasjoin(data [4]byte) bool {
	forward := "MAS"
	backward := "SAM"

	data1Valid := true

	for i := 0; i < 3; i++ {
		if data[i] != forward[i] {
			data1Valid = false
			break
		}
	}

	if !data1Valid {
		data1Valid = true
		for i := 0; i < 3; i++ {
			if data[i] != backward[i] {
				data1Valid = false
				break
			}

		}
	}
	return data1Valid
}
