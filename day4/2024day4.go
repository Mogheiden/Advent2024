package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	start := time.Now()
	bytesread, _ := os.ReadFile("day4.txt")
	day4data := strings.Split(string(bytesread), "\n")

	part1Answer := 0
	part2Answer := 0

	for i:= 0; i < len(day4data);i++{
		for j:=0; j < len(day4data[0]); j++{
			if day4data[i][j] == 'm' || day4data[i][j] == 'a'{
				continue
			}
			substring := make([]string, 0)
			if j < len(day4data[0]) - 3{
				for k:= 0; k < 4; k++{
					substring = append(substring,string(day4data[i][j+k]))
				}	
				if joiner(substring){
					part1Answer ++
				}			
			}
			substring = make([]string, 0)
			if i < len(day4data) - 3{
				for k:= 0; k < 4; k++{
					substring = append(substring,string(day4data[i+k][j]))
				}	
				if joiner(substring){
					part1Answer ++
				}			
			}
			substring = make([]string, 0)
			substring1 := make([]string,0)
			if i < len(day4data) - 3 &&  j<len(day4data[0])-3{
				for k:= 0; k < 4; k++{
					substring = append(substring,string(day4data[i+k][j+k]))
					substring1 = append(substring1, string(day4data[i+3-k][j+k]))
				}	
				if joiner(substring){
					part1Answer++
				}
				if joiner(substring1){
					part1Answer++
				}		
			}
		}
	}
	// fmt.Println(part1Answer)

	for i:= 0; i<len(day4data) - 2; i++{
		for j:= 0; j<len(day4data[0]) - 2; j++{
			if day4data[i][j] == 'x' || day4data[i][j] == 'a'{
				continue
			}
			xSlice1 := make([]string, 0)
			xSlice2 := make([]string, 0)

			for k:= 0; k < 3; k++{
				xSlice1 = append(xSlice1,string(day4data[i+k][j+k]))
				xSlice2 = append(xSlice2, string(day4data[i+2-k][j+k]))
			}
			if doublejoiner(xSlice1,xSlice2){
				part2Answer++
			}
			// fmt.Println(i,j)
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func joiner(data []string) bool{
	joined := strings.Join(data, "")
	if joined == "XMAS" || joined == "SAMX"{
		return true
	}
	return false
}

func doublejoiner(data1 []string, data2 []string) bool{
	joined1 := strings.Join(data1, "")
	joined2 := strings.Join(data2, "")
	if (joined1 == "MAS" || joined1 == "SAM") && (joined2 == "MAS" || joined2 == "SAM"){
		return true
	}
	return false
}