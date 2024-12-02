package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	// "time"
)

func main(){
	// start := time.Now()
	bytesread, _ := os.ReadFile("day2.txt")
	day2data := strings.Split(string(bytesread), "\n")

	part1Answer := 0
	part2Answer := 0


	for _, report := range day2data {
		reportList := strings.Split(report, " ")
		var numberedReport []int
		for _, value := range reportList{
			numValue, _ := strconv.Atoi(value)
			numberedReport = append(numberedReport, numValue)
		}
		if valid(numberedReport){
			part1Answer++
			part2Answer++
			continue
		}
		for i := range numberedReport {
			if valid(append(slices.Clone(numberedReport[:i]), numberedReport[i+1:]...)) {
				part2Answer++
				break
			}
		}
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

func absolutediff (number1 int,  number2 int) int {
	diff := number1 - number2
	if diff < 0{
		diff *= -1
	}
	return diff
}

func valid(r []int) bool{
	ascending := r[0] > r[1]
	for i := 1; i < len(r); i++ {
		diff := absolutediff(r[i], r[i-1])
		if diff < 1 || diff > 3 || (ascending != (r[i] < r[i-1])){
			return false
		}
	}
	// fmt.Println(r)
	return true
}
