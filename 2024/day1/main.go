package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(name string) string {
	bytes, _ := os.ReadFile(name)

	return string(bytes)
}

func splitStringInTwo(str string) []string {
	strs := []string{}

	newStr := ""
	foundSpace := false
	for i := 0; i < len(str); i++ {

		if str[i] != ' ' {
			newStr += string(str[i])
		}

		if str[i] == ' ' && !foundSpace {
			foundSpace = true
			strs = append(strs, newStr)
			newStr = ""
		}
	}

	strs = append(strs, newStr)

	return strs
}

func distanceCalculation(leftList, rightList []int) int {
	distances := []int{}

	for i := 0; i < len(leftList); i++ {
		num := math.Abs(float64(leftList[i] - rightList[i]))
		distances = append(distances, int(num))
	}

	totalDistance := 0

	for _, distance := range distances {
		totalDistance += distance
	}

	return totalDistance
}

func parseString(str string) ([]int, []int) {

	leftList := []int{}
	rightList := []int{}

	for _, line := range strings.Split(str, "\n") {
		// here, we need to split the string into two
		numString := splitStringInTwo(line)

		leftString := string(numString[0])
		rightString := string(numString[1])

		leftNum, _ := strconv.Atoi(leftString)
		rightNum, _ := strconv.Atoi(rightString)

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	// sort both

	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList
}

func calculateSimilarityScore(leftList, rightList []int) int {

	// take number in left list, see how many times it appears in right list
	// multiply

	totalSimilarityScore := 0
	for i := 0; i < len(leftList); i++ {
		leftNum := leftList[i]

		foundCounter := 0
		for j := 0; j < len(rightList); j++ {
			if leftNum == rightList[j] {
				foundCounter++
			}
		}

		totalSimilarityScore += foundCounter * leftNum

	}

	return totalSimilarityScore
}

func part2(fileName string) int {
	str := readFile(fileName)
	leftList, rightList := parseString(str)

	distance := calculateSimilarityScore(leftList, rightList)

	return distance
}

func part1(fileName string) int {
	str := readFile(fileName)
	leftList, rightList := parseString(str)

	distance := distanceCalculation(leftList, rightList)

	return distance
}

func main() {
	fmt.Printf("Answer to part 1: %d\n", part1("input.txt"))
	fmt.Printf("Answer to part 2: %d\n", part2("input.txt"))

}
