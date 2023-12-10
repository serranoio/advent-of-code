package main

import (
	"fmt"
	"os"
	"strings"
)

type LR struct {
	left  string
	right string
}

func initLR(left string, right string) *LR {

	return &LR{
		left:  strings.TrimSpace(left),
		right: strings.TrimSpace(right),
	}
}

func parse(fileName string) (string, map[string]*LR) {
	bytes, _ := os.ReadFile(fileName)

	route := ""

	restStart := 0
	for num, byte := range bytes {
		if byte == 'L' ||
			byte == 'R' {
			route += string(byte)
		} else {
			restStart = num
			break
		}
	}

	allDirections := make(map[string]*LR)
	for _, line := range strings.Split(string(bytes[restStart:]), "\n") {
		if line == "" {
			continue
		}

		locationEqualsLR := strings.Split(line, "=")

		lr := strings.Split(locationEqualsLR[1], ",")

		LR := initLR(strings.Trim(strings.TrimSpace(lr[0]), "("), strings.Trim(lr[1], ")"))

		allDirections[strings.TrimSpace(locationEqualsLR[0])] = LR
	}

	return route, allDirections
}

func endLocation(currentLocation string, allDirections map[string]*LR, route string, targetLocation string, steps int) int {
	for _, move := range route {
		lr := allDirections[currentLocation]

		if move == 'L' {
			currentLocation = lr.left
		} else {
			currentLocation = lr.right
		}

		steps++

		if currentLocation == targetLocation {
			return steps
		}
	}

	return endLocation(currentLocation, allDirections, route, targetLocation, steps)
}

func part1() int {
	route, allDirections := parse("test.txt")
	steps := endLocation("AAA", allDirections, route, "ZZZ", 0)
	return steps
}

func getAllStarts(allDirections map[string]*LR) []string {
	allStarts := []string{}

	for k := range allDirections {
		if k[2] == 'A' {
			allStarts = append(allStarts, k)
		}
	}

	return allStarts
}

func allAtZ(currentLocations []string) bool {
	for _, location := range currentLocations {
		if location[2] != 'Z' {
			return false
		}
	}

	return true
}

// func endLocationPart2(currentLocations []string, allDirections map[string]*LR, route string, steps int) int {
// 	for _, move := range route {
// 		for num, location := range currentLocations {

// 			lr := allDirections[location]
// 			if move == 'L' {
// 				currentLocations[num] = lr.left
// 			} else {
// 				currentLocations[num] = lr.right
// 			}
// 			// check all
// 		}
// 		steps++

// 		if allAtZ(currentLocations) {
// 			return steps
// 		}
// 	}

// 	return endLocationPart2(currentLocations, allDirections, route, steps)
// }

func endLocationPart2(currentLocation string, allDirections map[string]*LR, route string, steps int) int {
	for _, move := range route {
		lr := allDirections[currentLocation]

		if move == 'L' {
			currentLocation = lr.left
		} else {
			currentLocation = lr.right
		}

		steps++

		if currentLocation[2] == 'Z' {
			return steps
		}
	}

	return endLocationPart2(currentLocation, allDirections, route, steps)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2() {
	route, allDirections := parse("adventday8.txt")

	allStarts := getAllStarts(allDirections)

	allSteps := []int{}

	for _, start := range allStarts {
		steps := endLocationPart2(start, allDirections, route, 0)
		allSteps = append(allSteps, steps)
	}

	num := LCM(allSteps[0], allSteps[1], allSteps...)

	fmt.Println(num)
	// endLocationPart2(allStarts, allDirections, route, 0)

}

func main() {

	part1()

	part2()

}
