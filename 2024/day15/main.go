package main

import (
	"fmt"
	"os"
	"strings"

	clog "github.com/charmbracelet/log"
)

type Location struct {
	x int
	y int
}

func readFileIntoBoard(fileName string) ([][]rune, []rune, Location) {
	bytes, _ := os.ReadFile(fileName)

	fileContents := string(bytes)

	var board [][]rune
	var location Location

	var moves []rune
	foundMoves := false

	for y, line := range strings.Split(fileContents, "\n") {

		if len(line) == 0 {
			foundMoves = true
		}

		if foundMoves {
			for _, movesLine := range strings.Split(line, "\n") {
				for _, row := range movesLine {
					moves = append(moves, row)
				}
			}
		} else {
			var row []rune
			for x, c := range line {
				if c == '@' {
					location = Location{
						y: y,
						x: x,
					}
				}

				row = append(row, c)
			}
			board = append(board, row)
		}
	}

	return board, moves, location
}

func printBoard(board [][]rune) {
	for _, line := range board {
		for _, c := range line {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

// one more will cause a whole line to move
func checkCanMove(board [][]rune, location Location, currentMove rune, hasMoved bool) ([][]rune, bool) {
	previousLocation := location

	switch currentMove {
	case '>':
		location.x += 1
		break
	case '<':
		location.x -= 1
		break
	case '^':
		location.y -= 1
		break
	case 'v':
		location.y += 1
		break
	}

	// bounds check
	if !(location.x >= 0 && location.y >= 0 && location.x < len(board[0]) && location.y < len(board)) {
		// invalid move
		return board, hasMoved
	}

	if board[location.y][location.x] == '#' {
		// invalid move
		return board, hasMoved
	}

	if board[location.y][location.x] == 'O' {
		board, hasMoved = checkCanMove(board, location, currentMove, hasMoved)
	}

	if board[location.y][location.x] == '.' {
		board[previousLocation.y][previousLocation.x], board[location.y][location.x] = board[location.y][location.x], board[previousLocation.y][previousLocation.x]
		hasMoved = true
	}

	return board, hasMoved
}

func moveRobot(board [][]rune, moves []rune, location Location) {
	if len(moves) == 0 {
		return
	}

	currentMove, moves := moves[0], moves[1:]
	board, hasMoved := checkCanMove(board, location, currentMove, false)

	if hasMoved {
		switch currentMove {
		case '>':
			location.x += 1
			break
		case '<':
			location.x -= 1
			break
		case '^':
			location.y -= 1
			break
		case 'v':
			location.y += 1
			break
		}
	}

	moveRobot(board, moves, location)
}

func implementGoodsPositioningSystem(board [][]rune) int {
	sumOfCoordinates := 0

	for y, line := range board {
		for x, c := range line {
			if c == 'O' {
				sumOfCoordinates += 100*y + x
			}

		}
	}

	return sumOfCoordinates
}

func main() {
	board, moves, location := readFileIntoBoard("test.txt")
	moveRobot(board, moves, location)
	sumOfCoordinates := implementGoodsPositioningSystem(board)
	fmt.Println(sumOfCoordinates)

	board, moves, location = readFileIntoBoard("larger.txt")
	moveRobot(board, moves, location)
	sumOfCoordinates = implementGoodsPositioningSystem(board)
	fmt.Println(sumOfCoordinates)

	board, moves, location = readFileIntoBoard("input.txt")
	moveRobot(board, moves, location)
	sumOfCoordinates = implementGoodsPositioningSystem(board)
	fmt.Println(sumOfCoordinates)

	clog.Info("Hello, main!")
}
