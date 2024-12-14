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
	bytesread, _ := os.ReadFile("day13.txt")
	day13data := strings.Split(string(bytesread), "\n\n")
	part1Answer := 0
	part2Answer := 0

	var gameStructs []Game

	for _, game := range day13data {
		var gameRecord Game
		instructions := strings.Split(game, "\n")
		for instruction := range instructions {
			xCoord := 0
			coords := strings.Split(instructions[instruction], ",")
			yCoord, _ := strconv.Atoi(coords[1][3:])
			if instruction < 2 {
				xCoord, _ = strconv.Atoi(coords[0][12:])
			} else {
				xCoord, _ = strconv.Atoi(coords[0][9:])
			}
			if instruction == 0 {
				gameRecord.ax = xCoord
				gameRecord.ay = yCoord
			}
			if instruction == 1 {
				gameRecord.bx = xCoord
				gameRecord.by = yCoord
			}
			if instruction == 2 {
				gameRecord.x = xCoord
				gameRecord.y = yCoord
			}
		}
		// fmt.Println(gameRecord)
		gameStructs = append(gameStructs, gameRecord)
	}
	for _, game := range gameStructs {
		part1Answer += calculateTokens(game)
	}
	for _, game := range gameStructs {
		game.x = game.x + 10000000000000
		game.y = game.y + 10000000000000
		part2Answer += calculateTokens(game)
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

type Game struct {
	ax int
	ay int
	bx int
	by int
	x  int
	y  int
}

func calculateTokens(machine Game) int {
	// use Cramer's rule to solve the system of equations
	a, b := solveEquation(machine)

	// if no valid solution, return 0
	if a <= 0 || b <= 0 {
		return 0
	}
	// fmt.Println(a, b)
	// calculate the total tokens (3 per A press, 1 per B press)
	return (3 * a) + b
}

func solveEquation(m Game) (int, int) {
	// determinant of coefficients matrix
	d := m.ax*m.by - m.bx*m.ay
	d1 := m.x*m.by - m.y*m.bx
	d2 := m.y*m.ax - m.x*m.ay

	// check if we have whole number sols
	if d1%d != 0 || d2%d != 0 {
		return 0, 0
	}
	return d1 / d, d2 / d
}
