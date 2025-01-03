package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	bytesread, _ := os.ReadFile("day24.txt")
	day24data := strings.Split(string(bytesread), "\n\n")
	initials := strings.Split(day24data[0], "\n")
	relations := strings.Split(day24data[1], "\n")
	instructions := make(map[string][]instructionSet)

	determinedGates := make(map[string]int)

	xVal := 0
	yVal := 0

	for _, initial := range initials {
		instruction := strings.Split(initial, ": ")
		if instruction[0][0] == 'x' {
			intPart, _ := strconv.Atoi(instruction[0][1:])
			instrVal, _ := strconv.Atoi(instruction[1])
			xVal |= instrVal << intPart
		}
		if instruction[0][0] == 'y' {
			intPart, _ := strconv.Atoi(instruction[0][1:])
			instrVal, _ := strconv.Atoi(instruction[1])
			yVal |= instrVal << intPart
		}

		if instruction[1] == "1" {
			determinedGates[instruction[0]] = 1
		} else {
			determinedGates[instruction[0]] = 0
		}
	}
	for _, relation := range relations {
		parsed := strings.Split(relation, " ")
		first := parsed[0]
		second := parsed[2]
		instruction := parsed[1]
		result := parsed[4]

		instructions[result] = append(instructions[result], instructionSet{first: first, second: second, operation: instruction})
	}

	for len(instructions) > 0 {
		for key, value := range instructions {
			if _, exists := determinedGates[key]; exists {
				delete(instructions, key)
				continue
			}
			for _, relation := range value {
				val1, exists1 := determinedGates[relation.first]
				val2, exists2 := determinedGates[relation.second]
				if exists1 && exists2 {
					determinedGates[key] = parse(val1, val2, relation.operation)
					delete(instructions, key)
					break
				}
			}
		}
	}

	initialXORs := make(map[string]int)
	initialAnds := make(map[string]int)
	// var secondaryXORmap []string
	// var secondaryAndmap []string
	remainders := make(map[string]int)

	gateTypes := make(map[string]instructionSet)

	var wronguns []string

	for _, relation := range relations {
		parsed := strings.Split(relation, " ")
		first := parsed[0]
		second := parsed[2]
		instruction := parsed[1]
		result := parsed[4]
		gateTypes[result] = instructionSet{first: first, second: second, operation: instruction}
	}

	for key, value := range gateTypes {
		if value.first[1:] == value.second[1:] {
			number, _ := strconv.Atoi(value.second[1:])
			switch value.operation {
			case "XOR":
				initialXORs[key] = number
				// delete(gateTypes, key)
			case "AND":
				initialAnds[key] = number
				// delete(gateTypes, key)
			}
		}
	}
	for key, value := range gateTypes {
		if value.operation == "OR" {
			if gateTypes[value.first].operation == "AND" && gateTypes[value.second].operation == "AND" {
				val1, exist1 := initialAnds[value.first]
				val2, exist2 := initialAnds[value.second]
				if exist1 {
					remainders[key] = val1
				} else if exist2 {
					remainders[key] = val2
				} else {
					panic("oh no")
				}
			} else {
				if gateTypes[value.first].operation != "AND" {
					wronguns = append(wronguns, value.first)
				}
				if gateTypes[value.second].operation != "AND" {
					wronguns = append(wronguns, value.second)
				}
			}
		}
		if key[0] == 'z' && value.operation != "XOR" && key != "z45" {
			wronguns = append(wronguns, key)
		}
	}
	for key, value := range gateTypes {
		if value.operation == "XOR" {
			if _, exists := initialXORs[key]; exists {
				continue
			} else {
				operationArray := [2]string{value.first, value.second}
				for _, operand := range operationArray {
					if gateTypes[operand].operation == "XOR" {
						_, exists = initialXORs[operand]
						if !exists {
							wronguns = append(wronguns, operand)
						}
					} else if gateTypes[operand].operation != "OR" && key != "z01" {
						wronguns = append(wronguns, operand)
					}
				}
			}
		}
	}
	sort.Strings(wronguns)

	part1Answer := 0
	part2Answer := strings.Join(wronguns, ",")

	for key, value := range determinedGates {
		if key[0] == 'z' {
			intPart, _ := strconv.Atoi(key[1:])
			part1Answer |= value << intPart
		}
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func parse(first int, second int, instruction string) int {
	switch instruction {
	case "AND":
		return first & second
	case "OR":
		return first | second
	case "XOR":
		return first ^ second
	default:
		panic("oops!")
	}
}

type instructionSet struct {
	first     string
	second    string
	operation string
}

type keyValue struct {
	key   string
	value int
}
