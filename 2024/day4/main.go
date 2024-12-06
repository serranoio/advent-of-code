package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(name string) string {
	bytes, _ := os.ReadFile(name)

	return string(bytes)
}

type XMAS rune

const (
	X XMAS = 'X'
	M XMAS = 'M'
	A XMAS = 'A'
	S XMAS = 'S'
)

func searchXMAS(x int, y int, horizontal int, vertical int, coords [][]rune, currentRune rune) int {
	if x+horizontal < 0 || y+vertical < 0 || x+horizontal >= len(coords[0]) || y+vertical >= len(coords) {
		return 0
	}

	switch currentRune {
	case rune(X):
		if coords[y+vertical][x+horizontal] == rune(M) {
			return searchXMAS(x+horizontal, y+vertical, horizontal, vertical, coords, rune(M))
		}
		return 0
	case rune(M):
		if coords[y+vertical][x+horizontal] == rune(A) {
			return searchXMAS(x+horizontal, y+vertical, horizontal, vertical, coords, rune(A))
		}
		return 0
	case rune(A):
		if coords[y+vertical][x+horizontal] == rune(S) {
			return 1
		}
		return 0

		// case rune(S):
		// return 1
	}

	return 0
}

func searchAtEveryPoint(coords [][]rune) int {
	count := 0
	for y, line := range coords {
		for x, c := range line {
			if c == rune(X) {
				count += searchXMAS(x, y, 1, 1, coords, c)
				count += searchXMAS(x, y, -1, 1, coords, c)
				count += searchXMAS(x, y, 1, -1, coords, c)
				count += searchXMAS(x, y, -1, -1, coords, c)
				count += searchXMAS(x, y, 0, 1, coords, c)
				count += searchXMAS(x, y, 1, 0, coords, c)
				count += searchXMAS(x, y, 0, -1, coords, c)
				count += searchXMAS(x, y, -1, 0, coords, c)
			}
		}
	}

	return count
}

func convertFileToCoordsArray(name string) [][]rune {
	var board [][]rune

	input := readFile(name)

	for _, line := range strings.Split(input, "\n") {
		var runeLine []rune
		for _, r := range line {
			runeLine = append(runeLine, r)
		}
		board = append(board, runeLine)
	}

	return board
}

func part2(name string) int {

	return 0
}
func part1(name string) int {

	arr := convertFileToCoordsArray(name)
	count := searchAtEveryPoint(arr)

	fmt.Println(count)

	return count
}
func main() {

	part1("input.txt")
}
