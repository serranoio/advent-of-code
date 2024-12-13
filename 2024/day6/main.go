package main

import (
	"fmt"
	"os"
	"strings"
)

type BOARD rune

const (
	DOT      BOARD = '.'
	SHARP    BOARD = '#'
	VISITED  BOARD = 'X'
	LOCATION BOARD = '^'
)

var board [][]BOARD

type location struct {
	x int
	y int
}

var currentLocation location

type Direction int

const (
	UP    Direction = 1
	DOWN  Direction = 2
	LEFT  Direction = 3
	RIGHT Direction = 4
)

func readFile(filename string) string {
	bytes, _ := os.ReadFile(filename)

	return string(bytes)
}

func createBoard(contents string) {
	for y, line := range strings.Split(contents, "\n") {
		var boardLine []BOARD
		for x, char := range line {
			if BOARD(char) == LOCATION {
				currentLocation.x = x
				currentLocation.y = y
			}

			boardLine = append(boardLine, BOARD(char))
		}
		board = append(board, boardLine)
	}
}

var obsturctions []location

func doWeTurn(x int, y int) int {
	if !isInbouds(x, y) {
		return -1
	}

	if board[y][x] == SHARP {
		obsturctions = append(obsturctions, location{x, y})
		return 1
	}

	return 0
}
func printBoard() {

	for _, line := range board {
		for _, char := range line {
			fmt.Print(string(char) + " ")
		}
		fmt.Println()
	}
	fmt.Println()

}

func isInbouds(x int, y int) bool {
	if x < 0 || y < 0 || x >= len(board[0]) || y >= len(board) {
		return false
	}

	return true
}

func traverseBoard(currentDirection Direction) {
	if !isInbouds(currentLocation.x, currentLocation.y) {
		return
	}

	board[currentLocation.y][currentLocation.x] = VISITED
	// printBoard()

	switch currentDirection {
	case UP:
		turnVal := doWeTurn(currentLocation.x, currentLocation.y-1)
		if turnVal == 1 {
			currentLocation.x++
			traverseBoard(RIGHT)
			break
		} else if turnVal == -1 {
			return
		}
		currentLocation.y--
		traverseBoard(UP)
	case DOWN:
		turnVal := doWeTurn(currentLocation.x, currentLocation.y+1)
		if turnVal == 1 {
			currentLocation.x--
			traverseBoard(LEFT)
			break
		} else if turnVal == -1 {
			return
		}
		currentLocation.y++
		traverseBoard(DOWN)
	case LEFT:
		turnVal := doWeTurn(currentLocation.x-1, currentLocation.y)
		if turnVal == 1 {
			currentLocation.y--
			traverseBoard(UP)
			break
		} else if turnVal == -1 {
			return
		}
		currentLocation.x--
		traverseBoard(LEFT)
	case RIGHT:
		turnVal := doWeTurn(currentLocation.x+1, currentLocation.y)
		if turnVal == 1 {
			currentLocation.y++
			traverseBoard(DOWN)
			break
		} else if turnVal == -1 {
			return
		}
		currentLocation.x++
		traverseBoard(RIGHT)
	}
}

func countVisited() int {
	counter := 0
	for _, line := range board {
		for _, char := range line {
			if char == VISITED {
				counter++
			}
		}
	}

	return counter
}

type vector struct {
	xDirection int
	yDirection int
}

func isRectangle(points []location) bool {
	lastPoint := location{}

	xValues := []int{}
	yValues := []int{}

	firstLine := vector{
		yDirection: points[0].y - points[1].y,
		xDirection: points[0].x - points[1].x,
	}

	secondLine := vector{
		yDirection: points[0].y - points[2].y,
		xDirection: points[0].x - points[2].x,
	}

	thirdLine := vector{
		yDirection: points[1].y - points[2].y,
		xDirection: points[1].x - points[2].x,
	}

	allVectors := []vector{firstLine, secondLine, thirdLine}

	for _, vector := range allVectors {
		xValues = append(xValues, vector.xDirection)
		yValues = append(yValues, vector.yDirection)
	}

	xMap := make(map[int]int)
	yMap := make(map[int]int)

	for _, val := range xValues {
		if _, ok := xMap[val]; ok {
			xMap[val]++
		} else {
			xMap[val] = 1
		}
	}

	for _, val := range yValues {
		if _, ok := yMap[val]; ok {
			yMap[val]++
		} else {
			yMap[val] = 1
		}
	}

	for k, v := range xMap {
		if v == 2 {
			lastPoint.x = k
		}
	}

	for k, v := range yMap {
		if v == 2 {
			lastPoint.y = k
		}
	}

	fmt.Println(lastPoint)

	if board[lastPoint.y][lastPoint.x] == SHARP {
		return true
	}

	return false
}

func countRectangles() int {
	// 3, 5
	// 8, 5
	// 3, 2
	// 8, 2
	// *

	// every 3, we can discover a rectangle
	rectangleCount := 0

	// get the 4th theoretical point. If the point is marked as a sharp, rectangle.

	points := []location{}
	for i := 0; i < len(obsturctions); i++ {
		if i%3 == 0 && i != 0 {
			if isRectangle(points) {
				rectangleCount++
			}
			// can continue

		} else {
			points = append(points, obsturctions[i])
		}
	}

	return rectangleCount
}

func main() {
	contents := readFile("input.txt")
	createBoard(contents)
	traverseBoard(UP)
	printBoard()
	visited := countVisited()

	rectangleCount := countRectangles()

	fmt.Println(visited)
	fmt.Println(rectangleCount)

	// clog.Info("Hello, main!")
}
