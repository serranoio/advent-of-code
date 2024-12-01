package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const firstRow = "Time:      7  15   30"

func TestSplitFile(t *testing.T) {
	assert := assert.New(t)

	split := splitFile("test.txt")

	assert.Equal(len(split), 2, "they should be equal")
	assert.Equal(split[0], firstRow, "they should be equal")
}

func TestConvertRowIntoIntArray(t *testing.T) {
	assert := assert.New(t)

	array := convertRowToIntArray(firstRow)

	assert.Equal(len(array), 3, "They should be equal")
	assert.Equal(array[0], 7, "They should be equal")
	assert.Equal(array[1], 15, "They should be equal")
	assert.Equal(array[2], 30, "They should be equal")
}

func TestRacesPart1(t *testing.T) {
	assert := assert.New(t)
	split := splitFile("test.txt")

	times := convertRowToIntArray(split[0])
	distances := convertRowToIntArray(split[1])

	allRaceMaxes := testRaces(times, distances)

	assert.Equal(allRaceMaxes[0], 4, "They should be equal")
	assert.Equal(allRaceMaxes[1], 8, "They should be equal")
	assert.Equal(allRaceMaxes[2], 9, "They should be equal")
}

func TestRacesPart2(t *testing.T) {
	assert := assert.New(t)
	split := splitFile("test.txt")

	times := convertRowToIntArray(split[0])
	distances := convertRowToIntArray(split[1])

	allRaceMaxes := testRaces(times, distances)

	assert.Equal(allRaceMaxes[0], 4, "They should be equal")
	assert.Equal(allRaceMaxes[1], 8, "They should be equal")
	assert.Equal(allRaceMaxes[2], 9, "They should be equal")
}

func TestJoinIntArray(t *testing.T) {
	assert := assert.New(t)

	number := joinIntArray([]int{15, 3, 20, 18})

	assert.Equal(number, 1532018, "They should be equal")
}
