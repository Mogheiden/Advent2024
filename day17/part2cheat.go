package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Registers struct {
	A, B, C int
}

type Instruction struct {
	Opcode, Operand int
}

func parseInput(s []string) (Registers, []Instruction) {
	registers := Registers{}
	_, _ = fmt.Sscanf(s[0], "Register A: %d", &registers.A)
	_, _ = fmt.Sscanf(s[1], "Register B: %d", &registers.B)
	_, _ = fmt.Sscanf(s[2], "Register C: %d", &registers.C)
	var instructions []Instruction
	strInstructions := strings.Split(s[4][9:], ",")
	for idx := 0; idx < len(strInstructions); idx += 2 {
		opcode, _ := strconv.Atoi(strInstructions[idx])
		operand, _ := strconv.Atoi(strInstructions[idx+1])
		instructions = append(instructions, Instruction{
			Opcode:  opcode,
			Operand: operand,
		})
	}
	return registers, instructions
}

func run(instruction Instruction, registers *Registers, pc *int) string {
	var combo int
	switch instruction.Operand {
	case 0:
		combo = 0
	case 1:
		combo = 1
	case 2:
		combo = 2
	case 3:
		combo = 3
	case 4:
		combo = registers.A
	case 5:
		combo = registers.B
	case 6:
		combo = registers.C
	default:
		fmt.Println("Halt!")
	}
	var out string
	switch instruction.Opcode {
	case 0: // adv
		registers.A = registers.A / int(math.Pow(2, float64(combo)))
		*pc++
	case 1: // blx
		registers.B = registers.B ^ instruction.Operand
		*pc++
	case 2: // bts
		registers.B = combo % 8
		*pc++
	case 3: // jnz
		if registers.A != 0 {
			*pc = instruction.Operand / 2
		} else {
			*pc++
		}
	case 4: // bxc
		registers.B = registers.B ^ registers.C
		*pc++
	case 5: // out
		out = strconv.Itoa(combo % 8)
		*pc++
	case 6: // bdv
		registers.B = registers.A / int(math.Pow(2, float64(combo)))
		*pc++
	case 7: // cdv
		registers.C = registers.A / int(math.Pow(2, float64(combo)))
		*pc++
	}
	return out
}

func getOutput(s []string) string {
	registers, instructions := parseInput(s)
	var (
		pc     int = 0
		output bytes.Buffer
	)
	for pc < len(instructions) {
		out := run(instructions[pc], &registers, &pc)
		if out != "" {
			output.WriteString(out)
			output.WriteString(",")
		}
	}
	return output.String()[:len(output.String())-1]
}

func backPropagation(instructions []Instruction, position, initVal int) int {
	intIns := make([]int, len(instructions)*2)
	for idx, instruction := range instructions {
		intIns[idx*2] = instruction.Opcode
		intIns[idx*2+1] = instruction.Operand
	}
	// Checking instructions we can see A (as index) is divided by 8 on 0, 3 instruction in each iteration.
	for idx := 0; idx < 8; idx++ {
		registers := Registers{
			A: initVal*8 + idx,
			B: 0,
			C: 0,
		}
		pc := 0
		var output []int
		for pc < len(instructions) {
			out := run(instructions[pc], &registers, &pc)
			if out != "" {
				n, _ := strconv.Atoi(out)
				output = append(output, n)
			}
		}
		ok := true
		for j := position; j < len(intIns); j++ {
			if intIns[j] != output[j-position] {
				ok = false
				break
			}
		}
		if ok {
			if position == 0 {
				return initVal*8 + idx
			}
			if val := backPropagation(instructions, position-1, initVal*8+idx); val != -1 {
				return val
			}
		}
	}
	return -1
}

func findARegister(s []string) int {
	_, instructions := parseInput(s)
	return backPropagation(instructions, len(instructions)*2-1, 0)
}

func main() {
	bytesread, _ := os.ReadFile("cheatinput.txt")
	opcode := strings.Split(string(bytesread), "\n")

	fmt.Println(getOutput(opcode))
	fmt.Println(findARegister(opcode))
}
