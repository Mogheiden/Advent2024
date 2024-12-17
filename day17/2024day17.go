package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	bytesread, _ := os.ReadFile("day17.txt")
// 	opcode := strings.Split(string(bytesread), ",")

// 	goldenCopy := strings.Join(opcode, ",")
// 	fmt.Println(goldenCopy)

// 	registerA := 46323429
// 	registerB := 0
// 	registerC := 0

// 	// registerA := 729
// 	// registerB := 0
// 	// registerC := 0
// 	pc := 0
// 	var output string

// 	for pc < len(opcode) {
// 		instruction, err := strconv.Atoi(opcode[pc])
// 		operand, err2 := strconv.Atoi(opcode[pc+1])
// 		if err == nil && err2 == nil {
// 			switch instruction {
// 			case 0:
// 				comboVal := 1 << getComboOperand(operand, registerA, registerB, registerC)
// 				registerA = registerA / comboVal
// 				pc += 2
// 			case 1:
// 				registerB ^= operand
// 				pc += 2
// 			case 2:
// 				registerB = getComboOperand(operand, registerA, registerB, registerC) % 8
// 				pc += 2
// 			case 3:
// 				if registerA == 0 {
// 					pc += 2
// 				} else {
// 					pc = operand
// 				}
// 			case 4:
// 				registerB ^= registerC
// 				pc += 2
// 			case 5:
// 				comboVal := getComboOperand(operand, registerA, registerB, registerC) % 8
// 				output += strconv.Itoa(comboVal)
// 				output += ","
// 				pc += 2
// 			case 6:
// 				comboVal := 1 << getComboOperand(operand, registerA, registerB, registerC)
// 				registerB = registerA / comboVal
// 				pc += 2
// 			case 7:
// 				comboVal := 1 << getComboOperand(operand, registerA, registerB, registerC)
// 				registerC = registerA / comboVal
// 				pc += 2
// 			}
// 		}
// 	}
// 	registerA = 1
// 	registerB = 0
// 	registerC = 0

// 	output = ""

// 	fmt.Println(output)
// 	fmt.Println(time.Since(start))
// }

// func getComboOperand(comboNum int, regA int, regB int, regC int) int {
// 	switch comboNum {
// 	case 0:
// 		return 0
// 	case 1:
// 		return 1
// 	case 2:
// 		return 2
// 	case 3:
// 		return 3
// 	case 4:
// 		return regA
// 	case 5:
// 		return regB
// 	case 6:
// 		return regC
// 	default:
// 		panic("shouldn't be here")
// 	}
// }
