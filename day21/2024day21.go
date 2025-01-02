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
	bytesread, _ := os.ReadFile("day21test.txt")
	day21data := strings.Split(string(bytesread), "\n")

	part1Answer := 0
	part2Answer := 0

	numberPad := [][]byte{{'7', '8', '9'}, {'4', '5', '6'}, {'1', '2', '3'}, {' ', '0', 'A'}}

	numberPadRelations := make(map[byte]map[byte][5]int)

	for i := range numberPad {
		for j := range numberPad[i] {
			numberPadRelations[numberPad[i][j]] = makeRelations(i, j, numberPad)
		}
	}

	directionPad := [][]byte{{' ', '^', 'A'}, {'<', 'v', '>'}}

	directionPadRelations := make(map[byte]map[byte][5]int)

	for i := range directionPad {
		for j := range directionPad[i] {
			directionPadRelations[directionPad[i][j]] = makeRelations(i, j, directionPad)
		}
	}
	for _, input := range day21data {
		baseRobotPos := 'A'
		secondRobotPos := 'A'
		thirdRobotPos := 'A'
		// fourthRobotPos := 'A'
		var finalString []string
		for _, char := range input {
			relation := numberPadRelations[byte(baseRobotPos)][byte(char)]
			vulnerability := vulnerablePath(byte(baseRobotPos), byte(char))
			// fmt.Println(relation, string(baseRobotPos), string(char))
			goldenPath := unwrapOptimalPath(relation, baseRobotPos, numberPadRelations, vulnerability)
			baseRobotPos = char
			// fmt.Println(goldenPath)
			for _, subchar := range goldenPath {
				subVuln := vulnerablePath(byte(secondRobotPos), byte(subchar))
				relation := directionPadRelations[byte(secondRobotPos)][byte(subchar)]
				subGoldenPath := unwrapOptimalPath(relation, secondRobotPos, directionPadRelations, subVuln)
				secondRobotPos = subchar
				for _, subsubchar := range subGoldenPath {
					subsubVuln := vulnerablePath(byte(thirdRobotPos), byte(subsubchar))
					relation := directionPadRelations[byte(thirdRobotPos)][byte(subsubchar)]
					subsubGoldenPath := unwrapOptimalPath(relation, thirdRobotPos, directionPadRelations, subsubVuln)
					thirdRobotPos = subsubchar
					finalString = append(finalString, subsubGoldenPath)
				}
				// fmt.Println(subGoldenPath)
			}
		}
		joinedString := strings.Join(finalString, "")
		lenHalf := len(joinedString)
		intHalf, _ := strconv.Atoi(input[:len(input)-1])
		fmt.Println(lenHalf)
		part1Answer += intHalf * lenHalf
	}

	// distance := distanceFinder(input, baseRobotPos, numberPadRelations)

	// fmt.Println(distance)

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

func vulnerablePath(start byte, end byte) int {

	if start == '0' && end == '4' {
		return 3
	}

	if ((start == '7' || start == '4' || start == '1') && (end == '0' || end == 'A')) || (start == '<' && (end == '^' || end == 'A')) {
		return 1
	}

	if ((end == '7' || end == '4' || end == '1') && (start == '0' || start == 'A')) || (end == '<' && (start == '^' || start == 'A')) {
		return 2
	}
	return 0
}

func makeRelations(yCoord int, xCoord int, toMap [][]byte) map[byte][5]int {
	relationMap := make(map[byte][5]int)
	for i := range toMap {
		for j := range toMap[i] {
			if toMap[i][j] == ' ' {
				continue
			}
			var relation [5]int
			relation[4] = 1
			yDist := i - yCoord
			if yDist < 0 {
				relation[UP] = yDist * -1
			} else {
				relation[DOWN] = yDist
			}
			xDist := j - xCoord
			if xDist < 0 {
				relation[LEFT] = xDist * -1
			} else {
				relation[RIGHT] = xDist
			}
			relationMap[toMap[i][j]] = relation
		}
	}
	return relationMap
}

func distanceFinder(input string, current rune, relationMap map[byte]map[byte][5]int) int {
	distance := 0
	for _, char := range input {
		distance += arraySum(relationMap[byte(current)][byte(char)])
		current = char
	}
	return distance
}

func arraySum(array [5]int) int {
	retVal := 0
	for _, val := range array {
		retVal += val
	}
	return retVal
}

func unwrapOptimalPath(array [5]int, position rune, relationMap map[byte]map[byte][5]int, vulnerability int) string {
	strings := makeStrings(array)

	if vulnerability == 3 {
		return "^<^A"
	}

	if vulnerability == 1 {
		return strings[0]
	}
	if vulnerability == 2 {
		return strings[1]
	}
	xFirst := distanceFinder(strings[0], position, relationMap)
	yFirst := distanceFinder(strings[1], position, relationMap)

	if xFirst > yFirst {
		return strings[1]
	} else {
		return strings[0]
	}

}

func makeStrings(array [5]int) [2]string {
	var xFirst string
	var yFirst string

	for range array[UP] {
		yFirst += "^"
	}
	for range array[DOWN] {
		yFirst += "v"
	}
	for range array[LEFT] {
		xFirst += "<"
	}
	for range array[RIGHT] {
		xFirst += ">"
	}
	for range array[UP] {
		xFirst += "^"
	}
	for range array[DOWN] {
		xFirst += "v"
	}
	for range array[LEFT] {
		yFirst += "<"
	}
	for range array[RIGHT] {
		yFirst += ">"
	}

	xFirst += "A"
	yFirst += "A"
	group := [2]string{xFirst, yFirst}
	return group
}

const (
	UP      int = 0
	DOWN    int = 1
	LEFT    int = 2
	RIGHT   int = 3
	CONFIRM int = 4
)

// package main

// import (
// 	_ "embed"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"unicode"
// )

// type sequenceKey struct {
// 	sequence string
// 	depth    int
// }

// var sequenceCache = make(map[sequenceKey]int)

// func main() {
// 	codes := readInput()

// 	// part 1
// 	complexityScore := 0
// 	for _, code := range codes {
// 		complexityScore += calculateScore(code, 3) // 3 robots
// 	}
// 	fmt.Printf("Part 1 - Total complexity score: %d\n", complexityScore)

// 	// part 2
// 	complexityScore = 0
// 	for _, code := range codes {
// 		complexityScore += calculateScore(code, 26) // 26 robots
// 	}
// 	fmt.Printf("Part 2 - Total complexity score: %d\n", complexityScore)
// }

// func calculateScore(code string, robots int) int {
// 	numericCode, _ := codeToInteger(code)
// 	length := getSequenceLength(code, robots)
// 	return numericCode * length
// }

// func getSequenceLength(targetSequence string, depth int) int {
// 	key := sequenceKey{sequence: targetSequence, depth: depth}
// 	if value, exists := sequenceCache[key]; exists {
// 		return value
// 	}

// 	length := 0
// 	if depth == 0 {
// 		length = len(targetSequence)
// 	} else {
// 		current := 'A'
// 		for _, next := range targetSequence {
// 			len := getMoveCount(current, next, depth)
// 			current = next
// 			length += len
// 		}
// 	}

// 	sequenceCache[key] = length
// 	return length
// }

// func getMoveCount(current, next rune, depth int) int {
// 	if current == next {
// 		return 1
// 	}
// 	newSequence := paths[buttonPair{first: current, second: next}]
// 	return getSequenceLength(newSequence, depth-1)
// }

// func codeToInteger(input string) (int, error) {
// 	var numericPart strings.Builder
// 	for _, char := range input {
// 		if unicode.IsDigit(char) {
// 			numericPart.WriteRune(char)
// 		}
// 	}

// 	return strconv.Atoi(numericPart.String())
// }

// func readInput() []string {
// 	bytesread, _ := os.ReadFile("day21.txt")
// 	day21data := strings.Split(string(bytesread), "\n")

// 	return day21data
// }

// type buttonPair struct {
// 	first  rune
// 	second rune
// }

// var paths = map[buttonPair]string{
// 	{'A', '0'}: "<A",
// 	{'0', 'A'}: ">A",
// 	{'A', '1'}: "^<<A",
// 	{'1', 'A'}: ">>vA",
// 	{'A', '2'}: "<^A",
// 	{'2', 'A'}: "v>A",
// 	{'A', '3'}: "^A",
// 	{'3', 'A'}: "vA",
// 	{'A', '4'}: "^^<<A",
// 	{'4', 'A'}: ">>vvA",
// 	{'A', '5'}: "<^^A",
// 	{'5', 'A'}: "vv>A",
// 	{'A', '6'}: "^^A",
// 	{'6', 'A'}: "vvA",
// 	{'A', '7'}: "^^^<<A",
// 	{'7', 'A'}: ">>vvvA",
// 	{'A', '8'}: "<^^^A",
// 	{'8', 'A'}: "vvv>A",
// 	{'A', '9'}: "^^^A",
// 	{'9', 'A'}: "vvvA",
// 	{'0', '1'}: "^<A",
// 	{'1', '0'}: ">vA",
// 	{'0', '2'}: "^A",
// 	{'2', '0'}: "vA",
// 	{'0', '3'}: "^>A",
// 	{'3', '0'}: "<vA",
// 	{'0', '4'}: "^<^A",
// 	{'4', '0'}: ">vvA",
// 	{'0', '5'}: "^^A",
// 	{'5', '0'}: "vvA",
// 	{'0', '6'}: "^^>A",
// 	{'6', '0'}: "<vvA",
// 	{'0', '7'}: "^^^<A",
// 	{'7', '0'}: ">vvvA",
// 	{'0', '8'}: "^^^A",
// 	{'8', '0'}: "vvvA",
// 	{'0', '9'}: "^^^>A",
// 	{'9', '0'}: "<vvvA",
// 	{'1', '2'}: ">A",
// 	{'2', '1'}: "<A",
// 	{'1', '3'}: ">>A",
// 	{'3', '1'}: "<<A",
// 	{'1', '4'}: "^A",
// 	{'4', '1'}: "vA",
// 	{'1', '5'}: "^>A",
// 	{'5', '1'}: "<vA",
// 	{'1', '6'}: "^>>A",
// 	{'6', '1'}: "<<vA",
// 	{'1', '7'}: "^^A",
// 	{'7', '1'}: "vvA",
// 	{'1', '8'}: "^^>A",
// 	{'8', '1'}: "<vvA",
// 	{'1', '9'}: "^^>>A",
// 	{'9', '1'}: "<<vvA",
// 	{'2', '3'}: ">A",
// 	{'3', '2'}: "<A",
// 	{'2', '4'}: "<^A",
// 	{'4', '2'}: "v>A",
// 	{'2', '5'}: "^A",
// 	{'5', '2'}: "vA",
// 	{'2', '6'}: "^>A",
// 	{'6', '2'}: "<vA",
// 	{'2', '7'}: "<^^A",
// 	{'7', '2'}: "vv>A",
// 	{'2', '8'}: "^^A",
// 	{'8', '2'}: "vvA",
// 	{'2', '9'}: "^^>A",
// 	{'9', '2'}: "<vvA",
// 	{'3', '4'}: "<<^A",
// 	{'4', '3'}: "v>>A",
// 	{'3', '5'}: "<^A",
// 	{'5', '3'}: "v>A",
// 	{'3', '6'}: "^A",
// 	{'6', '3'}: "vA",
// 	{'3', '7'}: "<<^^A",
// 	{'7', '3'}: "vv>>A",
// 	{'3', '8'}: "<^^A",
// 	{'8', '3'}: "vv>A",
// 	{'3', '9'}: "^^A",
// 	{'9', '3'}: "vvA",
// 	{'4', '5'}: ">A",
// 	{'5', '4'}: "<A",
// 	{'4', '6'}: ">>A",
// 	{'6', '4'}: "<<A",
// 	{'4', '7'}: "^A",
// 	{'7', '4'}: "vA",
// 	{'4', '8'}: "^>A",
// 	{'8', '4'}: "<vA",
// 	{'4', '9'}: "^>>A",
// 	{'9', '4'}: "<<vA",
// 	{'5', '6'}: ">A",
// 	{'6', '5'}: "<A",
// 	{'5', '7'}: "<^A",
// 	{'7', '5'}: "v>A",
// 	{'5', '8'}: "^A",
// 	{'8', '5'}: "vA",
// 	{'5', '9'}: "^>A",
// 	{'9', '5'}: "<vA",
// 	{'6', '7'}: "<<^A",
// 	{'7', '6'}: "v>>A",
// 	{'6', '8'}: "<^A",
// 	{'8', '6'}: "v>A",
// 	{'6', '9'}: "^A",
// 	{'9', '6'}: "vA",
// 	{'7', '8'}: ">A",
// 	{'8', '7'}: "<A",
// 	{'7', '9'}: ">>A",
// 	{'9', '7'}: "<<A",
// 	{'8', '9'}: ">A",
// 	{'9', '8'}: "<A",
// 	{'<', '^'}: ">^A",
// 	{'^', '<'}: "v<A",
// 	{'<', 'v'}: ">A",
// 	{'v', '<'}: "<A",
// 	{'<', '>'}: ">>A",
// 	{'>', '<'}: "<<A",
// 	{'<', 'A'}: ">>^A",
// 	{'A', '<'}: "v<<A",
// 	{'^', 'v'}: "vA",
// 	{'v', '^'}: "^A",
// 	{'^', '>'}: "v>A",
// 	{'>', '^'}: "<^A",
// 	{'^', 'A'}: ">A",
// 	{'A', '^'}: "<A",
// 	{'v', '>'}: ">A",
// 	{'>', 'v'}: "<A",
// 	{'v', 'A'}: "^>A",
// 	{'A', 'v'}: "<vA",
// 	{'>', 'A'}: "^A",
// 	{'A', '>'}: "vA",
// }
