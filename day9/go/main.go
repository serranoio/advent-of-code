package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func getAllLines(fileName string) []string {
	bytes, _ := os.ReadFile(path.Join("../", fileName))

	return strings.Split(string(bytes), "\n")
}

func allZeros(line []int) bool {
	for _, num := range line {
		if num != 0 {
			return false
		}
	}

	return true
}
func gatherDifferences(line []int, current, next int, remembered []int) []int {
	if next+1 > len(line) {
		return remembered
	}

	difference := line[next] - line[current]

	remembered = append(remembered, difference)

	return gatherDifferences(line, next, next+1, remembered)
}

func calculateDifferences(line []int, allLines [][]int) [][]int {
	if allZeros(line) {
		return allLines
	}
	// calculate differences
	nextLine := gatherDifferences(line, 0, 1, []int{})
	allLines = append(allLines, nextLine)

	return calculateDifferences(nextLine, allLines)
}

func calculateHistory2(line []int) int {
	differences := calculateDifferences(line, [][]int{line})

	top := differences[len(differences)-1]
	top = append([]int{0}, top...)

	for num := 1; num < len(differences); num++ {
		next := differences[len(differences)-1-num]

		nextValue := next[0] - top[0]
		// next = append(next, next[len(next)-1]+top[len(top)-1])

		next = append([]int{nextValue}, next...)

		top = next
	}

	return top[0]
}

func calculateHistory(line []int) int {
	differences := calculateDifferences(line, [][]int{line})

	top := differences[len(differences)-1]
	top = append([]int{0}, top...)

	for num := 1; num < len(differences); num++ {
		next := differences[len(differences)-1-num]

		// nextValue := next[0] + top[0]
		next = append(next, next[len(next)-1]+top[len(top)-1])

		// next = append([]int{nextValue}, next...)

		top = next
	}

	return top[len(top)-1]
}

func toDoubleDimensionArray(lines []string, partOne bool) []int {
	allLines := []int{}

	for _, line := range lines {
		intLine := []int{}
		for _, num := range strings.Split(line, " ") {
			numInt, _ := strconv.Atoi(num)

			intLine = append(intLine, numInt)
		}
		// here we calculate Historys!
		var history int
		if partOne {
			history = calculateHistory(intLine)
		} else {
			history = calculateHistory2(intLine)
		}
		allLines = append(allLines, history)
	}
	return allLines
}
func sum(all []int) int {
	total := 0

	for _, num := range all {
		total += num
	}

	return total
}

func main() {
	timer := time.Now()

	lines := getAllLines("adventday9.txt")

	all := toDoubleDimensionArray(lines, true)
	theSum := sum(all)
	fmt.Println("sum for part 1:", theSum)

	all = toDoubleDimensionArray(lines, false)
	theSum = sum(all)
	fmt.Println("sum for part 2:", theSum)

	fmt.Printf("Runtime: %d nanoseconds\n", time.Since(timer).Nanoseconds())
}
