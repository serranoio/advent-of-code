package main

import (
	"fmt"
	"os"
	"strings"
)

func parseFile(fileName string) [][]string {
	bytes, _ := os.ReadFile(fileName)

	allGroups := [][]string{}
	allLines := []string{}

	for _, line := range strings.Split(string(bytes), "\n") {
		if line == "" {
			allGroups = append(allGroups, allLines)
			allLines = []string{}
		} else {
			allLines = append(allLines, line)
			line = ""
		}
	}

	allGroups = append(allGroups, allLines)

	return allGroups
}

func equalStrings(leftString, rightString string) int {
	counter := 0

	for pos := 0; pos < len(leftString); pos++ {
		if leftString[pos] != rightString[pos] {
			counter++
		}
	}

	return counter
}

func findReflectionH(lines []string, posLeft, posRight, errors int) (bool, int) {
	if posLeft < 0 || posRight >= len(lines) { // im gonna be fucked by an edge case lol
		return true, 0
	}

	if lines[posLeft] != lines[posRight] {
		errors += equalStrings(lines[posLeft], lines[posRight])
		if errors > 1 {
			return false, errors

		}
	}

	posLeft -= 1
	posRight += 1
	return findReflectionH(lines, posLeft, posRight, errors)
}

func findReflectionV(line string, posLeft, posRight, errors int) (bool, int) {
	if posLeft < 0 || posRight >= len(line) { // im gonna be fucked by an edge case lol
		return true, errors
	}

	if line[posLeft] != line[posRight] {
		errors += 1
		if errors > 1 {
			return false, errors
		}
	}

	posLeft -= 1
	posRight += 1
	return findReflectionV(line, posLeft, posRight, errors)
}

func findMirrors(group []string, reflectionPointX, reflectionPointY int) int {
	// find vertical reflections
	errors := 0
	foundReflection := true
	if reflectionPointX >= len(group[0])-1 {
		foundReflection = false
	} else {
		for _, line := range group {
			didNotFind, verticalErrors := findReflectionV(line, reflectionPointX, reflectionPointX+1, 0)
			errors += verticalErrors

			if errors > 1 {
				break
			}
			if !didNotFind {
				foundReflection = false
				break
			}
		}
	}

	if errors == 1 {
		return reflectionPointX + 1
	} else {
		foundReflection = true
	}
	errors = 0
	isReflection, errors := findReflectionH(group, reflectionPointY, reflectionPointY+1, 0)
	if reflectionPointY >= len(group)-1 || !isReflection {
		foundReflection = false
	}

	if foundReflection || errors == 1 {
		return (reflectionPointY + 1) * 100
	}

	return findMirrors(group, reflectionPointX+1, reflectionPointY+1)
}

func getSum(lines [][]string) int {

	sum := 0
	for _, group := range lines {
		sum += findMirrors(group, 0, 0)
	}

	return sum
}

func main() {
	lines := parseFile("adventday13.txt")
	fmt.Println(getSum(lines))
}
