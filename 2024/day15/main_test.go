package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) { assert.Equal(t, 123, 123, "") }

func TestCheckCanMove(t *testing.T) {
	board := [][]rune{
		[]rune{'.', '.', '@', 'O', '.', '.', '#'},
	}
	location := Location{
		x: 2,
		y: 0,
	}
	currentMove := '>'
	board, hasMoved := checkCanMove(board, location, currentMove, false)
	boardResult := [][]rune{
		[]rune{'.', '.', '.', '@', 'O', '.', '#'},
	}
	assert.Equal(t, boardResult, board)
	assert.Equal(t, true, hasMoved)

	board = [][]rune{
		[]rune{'.', '.', '@', 'O', 'O', '.', '#'},
	}
	location = Location{
		x: 2,
		y: 0,
	}
	currentMove = '>'
	board, hasMoved = checkCanMove(board, location, currentMove, false)
	boardResult = [][]rune{
		[]rune{'.', '.', '.', '@', 'O', 'O', '#'},
	}
	assert.Equal(t, boardResult, board)
	assert.Equal(t, true, hasMoved)

	board = [][]rune{
		[]rune{'.', '.', '@', '#', 'O', '.', '#'},
	}
	location = Location{
		x: 2,
		y: 0,
	}
	currentMove = '>'
	board, hasMoved = checkCanMove(board, location, currentMove, false)
	boardResult = [][]rune{
		[]rune{'.', '.', '@', '#', 'O', '.', '#'},
	}
	assert.Equal(t, false, hasMoved)
	assert.Equal(t, boardResult, board)

}
