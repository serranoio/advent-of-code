package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFileIntoContents(fileName string) string {
	bytes, _ := os.ReadFile(fileName)

	return string(bytes)
}

func createBoard(contents string) [][]int {
	board := [][]int{}
	for _, line := range strings.Split(contents, "\n") {
		row := []int{}
		for _, r := range line {
			num, _ := strconv.Atoi(fmt.Sprintf("%c", r))
			row = append(row, num)
		}
		board = append(board, row)
	}

	return board
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

type Location struct {
	x int
	y int
}

func getNumAtLocation(location Location, board [][]int) int {
	if location.y < 0 || location.y >= len(board) || location.x < 0 || location.x >= len(board[0]) {
		return -1
	}

	return board[location.y][location.x]
}

func alreadyFound(found9Locations *[]Location, location Location) bool {
	for _, nineLocation := range *found9Locations {
		if nineLocation.x == location.x && nineLocation.y == location.y {
			return true
		}
	}
	return false
}

func launchTrailFinder(location Location, board [][]int, found9Locations *[]Location, part1 bool) {
	num := getNumAtLocation(location, board)

	if !part1 && num == 9 || num == 9 && !alreadyFound(found9Locations, location) {
		*found9Locations = append(*found9Locations, location)
		return
	}

	if num == getNumAtLocation(Location{location.x + 1, location.y}, board)-1 {
		launchTrailFinder(Location{location.x + 1, location.y}, board, found9Locations, part1)
	}
	if num == getNumAtLocation(Location{location.x - 1, location.y}, board)-1 {
		launchTrailFinder(Location{location.x - 1, location.y}, board, found9Locations, part1)
	}
	if num == getNumAtLocation(Location{location.x, location.y + 1}, board)-1 {
		launchTrailFinder(Location{location.x, location.y + 1}, board, found9Locations, part1)
	}
	if num == getNumAtLocation(Location{location.x, location.y - 1}, board)-1 {
		launchTrailFinder(Location{location.x, location.y - 1}, board, found9Locations, part1)
	}
}

func countTrailScore(board [][]int, part1 bool) int {
	trailScore := 0
	for y, row := range board {
		for x, c := range row {
			if c == 0 {
				found9Locations := []Location{}
				launchTrailFinder(Location{x: x, y: y}, board, &found9Locations, part1)
				trailScore += len(found9Locations)
			}
		}
	}

	return trailScore
}

func main() {
	contents := readFileIntoContents("input.txt")
	board := createBoard(contents)
	trailScore := countTrailScore(board, true)
	fmt.Println("part1: ", trailScore)

	contents = readFileIntoContents("input.txt")
	board = createBoard(contents)
	trailScore = countTrailScore(board, false)
	fmt.Println("part2: ", trailScore)

}
