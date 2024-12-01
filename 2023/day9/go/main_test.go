package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFile(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, len(getAllLines("test.txt")), "they should be equal")
}

func TestGatherDifferences(t *testing.T) {
	assert := assert.New(t)

	differences := gatherDifferences([]int{0, 3, 6, 9, 12, 15}, 0, 1, []int{})

	assert.Equal([]int{3, 3, 3, 3, 3}, differences, "they should be equal")

	differences = gatherDifferences(differences, 0, 1, []int{})

	assert.Equal([]int{0, 0, 0, 0}, differences, "they should be equal")
}

func TestCalculateDifferences(t *testing.T) {
	assert := assert.New(t)
	var start = []int{0, 3, 6, 9, 12, 15}

	differences := calculateDifferences(start, [][]int{start})

	assert.Equal([]int{3, 3, 3, 3, 3}, differences[1], "equal!")
	assert.Equal([]int{0, 0, 0, 0}, differences[2], "equal!")
}

func TestCalculateHistory(t *testing.T) {
	assert := assert.New(t)

	history := calculateHistory([]int{0, 3, 6, 9, 12, 15})

	assert.Equal(18, history, "Equal!")

	history = calculateHistory([]int{1, 3, 6, 10, 15, 21})

	assert.Equal(28, history, "Equal!")

	history = calculateHistory([]int{10, 13, 16, 21, 30, 45})

	assert.Equal(68, history, "Equal!")
}
func TestCalculateHistory2(t *testing.T) {
	assert := assert.New(t)

	history := calculateHistory2([]int{0, 3, 6, 9, 12, 15})

	assert.Equal(-3, history, "Equal!")

	history = calculateHistory2([]int{1, 3, 6, 10, 15, 21})

	assert.Equal(0, history, "Equal!")

	history = calculateHistory2([]int{10, 13, 16, 21, 30, 45})

	assert.Equal(5, history, "Equal!")
}

func TestSumAll(t *testing.T) {

	assert := assert.New(t)

	lines := getAllLines("test.txt")
	all := toDoubleDimensionArray(lines, true)

	total := sum(all)

	assert.Equal(114, total, "Equal!")

	all = toDoubleDimensionArray(lines, false)

	total = sum(all)

	assert.Equal(2, total, "Equal!")

}
