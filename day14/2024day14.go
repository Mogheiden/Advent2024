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
	bytesread, _ := os.ReadFile("day14.txt")
	day14data := strings.Split(string(bytesread), "\n")
	part1Answer := 1
	part2Answer := 0
	cycles := 100

	mapX := 101
	mapY := 103

	roboMap := make(map[[2]int]int)

	for _, vector := range day14data {
		startVector := strings.Split(vector, " ")
		startVector[0] = strings.TrimPrefix(startVector[0], "p=")
		startVector[1] = strings.TrimPrefix(startVector[1], "v=")
		startingCoords := strings.Split(startVector[0], ",")
		currentX, _ := strconv.Atoi(startingCoords[0])
		currentY, _ := strconv.Atoi(startingCoords[1])

		velocity := strings.Split(startVector[1], ",")
		velocityX, _ := strconv.Atoi(velocity[0])
		velocityY, _ := strconv.Atoi(velocity[1])
		for range cycles {
			currentX += velocityX
			currentY += velocityY
			if currentX >= mapX {
				currentX = currentX % mapX
			}
			if currentY >= mapY {
				currentY = currentY % mapY
			}
			if currentX < 0 {
				currentX += mapX
			}
			if currentY < 0 {
				currentY += mapY
			}
		}
		mapInput := [2]int{currentX, currentY}
		roboMap[mapInput]++
	}
	quadrantScore := [4]int{0}
	midpointX := mapX / 2
	midpointY := mapY / 2
	for key, value := range roboMap {
		if key[0] < midpointX && key[1] < midpointY {
			quadrantScore[0] += value
			continue
		}
		if key[0] > midpointX && key[1] < midpointY {
			quadrantScore[1] += value
			continue
		}
		if key[0] < midpointX && key[1] > midpointY {
			quadrantScore[2] += value
			continue
		}
		if key[0] > midpointX && key[1] > midpointY {
			quadrantScore[3] += value
			continue
		}
	}
	for _, mult := range quadrantScore {
		part1Answer *= mult
	}
	fmt.Println(part1Answer)
	var roboArray []Robot
	for _, vector := range day14data {
		startVector := strings.Split(vector, " ")
		startVector[0] = strings.TrimPrefix(startVector[0], "p=")
		startVector[1] = strings.TrimPrefix(startVector[1], "v=")
		startingCoords := strings.Split(startVector[0], ",")
		currentX, _ := strconv.Atoi(startingCoords[0])
		currentY, _ := strconv.Atoi(startingCoords[1])
		velocity := strings.Split(startVector[1], ",")
		velocityX, _ := strconv.Atoi(velocity[0])
		velocityY, _ := strconv.Atoi(velocity[1])
		newBot := Robot{currentX: currentX, currentY: currentY, velX: velocityX, velY: velocityY}
		roboArray = append(roboArray, newBot)

	}

	for {
		part2Map := make(map[[2]int]int)
		for i, robot := range roboArray {
			newBot := Robot{currentX: 0, currentY: 0, velX: robot.velX, velY: robot.velY}
			currentX := robot.currentX + robot.velX
			currentY := robot.currentY + robot.velY
			if currentX >= mapX {
				currentX = currentX % mapX
			}
			if currentY >= mapY {
				currentY = currentY % mapY
			}
			if currentX < 0 {
				currentX += mapX
			}
			if currentY < 0 {
				currentY += mapY
			}
			mapInput := [2]int{currentX, currentY}
			newBot.currentX = currentX
			newBot.currentY = currentY
			roboArray[i] = newBot
			part2Map[mapInput]++
		}
		part2Answer++
		if len(part2Map) == len(roboArray) {
			break
		}
	}
	fmt.Println(part2Answer)
	fmt.Println(time.Since(start))
}

type Robot struct {
	currentX int
	currentY int
	velX     int
	velY     int
}
