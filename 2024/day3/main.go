package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	clog "github.com/charmbracelet/log"
)

func readFile(name string) string {
	bytes, _ := os.ReadFile(name)

	return string(bytes)
}

func stringToInt(str string) int {
	num, _ := strconv.Atoi(str)

	return num
}

func parseStringToMul(str string, regex string) [][]int {
	rx, _ := regexp.Compile(regex)

	matches := rx.FindAllString(str, -1)

	allMatches := []string{}
	rx, _ = regexp.Compile(`\d{1,3},\d{1,3}`)
	for _, match := range matches {
		allMatches = append(allMatches, rx.FindAllString(match, -1)...)
	}

	numbers := [][]int{}
	rx, _ = regexp.Compile(`\d{1,3}`)

	for _, match := range allMatches {
		numberSet := []int{}

		matches := rx.FindAllString(match, -1)

		numberSet = append(numberSet, stringToInt(matches[0]))
		numberSet = append(numberSet, stringToInt(matches[1]))

		numbers = append(numbers, numberSet)
	}

	return numbers
}

func mulAll(numbers [][]int) int {
	totalCount := 0
	for _, numberSet := range numbers {
		totalCount += numberSet[0] * numberSet[1]
	}

	return totalCount
}

func part1(fileName string) int {
	contents := readFile(fileName)
	matches := parseStringToMul(contents, `mul\(\d{1,3},\d{1,3}\)`)

	mul := mulAll(matches)

	return mul
}

func parseStringToMulPart2(str string, regex string) [][]int {
	rx, _ := regexp.Compile(regex)
	matches := rx.FindAllString(str, -1)

	matchesString := strings.Join(matches, "")

	rx, _ = regexp.Compile(`don't\(\)`)
	dontIndexes := rx.FindAllIndex([]byte(matchesString), -1)

	rx, _ = regexp.Compile(`do\(\)`)
	doIndexes := rx.FindAllIndex([]byte(matchesString), -1)

	constructNewString := ""
	negativeSpace := [][]int{}
	for i := 0; i < len(dontIndexes); i++ {
		// all we need to do is to pop a number from do, and dont index
		dontIndex := dontIndexes[i][1]

		j := 0

		doIndex := 0
		for ; j < len(doIndexes); j++ {
			doIndex = doIndexes[j][0]

			if doIndex > dontIndex {
				break
			}
		}

		if doIndex > len(matchesString) || dontIndex > len(matchesString) {
			break
		}

		negativeSpace = append(negativeSpace, []int{dontIndex, doIndex})
	}

	for i := 0; i < len(negativeSpace); i++ {
		beg := negativeSpace[i][0]
		end := negativeSpace[i][1]

		if beg > end {
			end = len(matchesString)
		}

		if beg > len(matchesString) {
			end = len(matchesString)
			break
		}
		if end > len(matchesString) {
			end = len(matchesString)

		}

		constructNewString += matchesString[beg:end]
	}

	fmt.Println(constructNewString)

	// allMatches := []string{}
	// rx, _ = regexp.Compile(`\d{1,3},\d{1,3}`)
	// for _, match := range matches {
	// 	allMatches = append(allMatches, rx.FindAllString(match, -1)...)
	// }

	// numbers := [][]int{}
	// rx, _ = regexp.Compile(`\d{1,3}`)

	// for _, match := range allMatches {
	// 	numberSet := []int{}

	// 	matches := rx.FindAllString(match, -1)

	// 	numberSet = append(numberSet, stringToInt(matches[0]))
	// 	numberSet = append(numberSet, stringToInt(matches[1]))

	// 	numbers = append(numbers, numberSet)
	// }

	return [][]int{}
}
func part2(fileName string) int {
	contents := readFile(fileName)
	matches := parseStringToMulPart2(contents, `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)|`)

	mul := mulAll(matches)

	fmt.Println(mul)

	return mul
}

func main() {
	clog.Info("Hello, main!")
	fmt.Printf("Answer to part 1: %d\n", part1("input.txt"))
	fmt.Printf("Answer to part 2: %d\n", part2("input.txt"))
}
