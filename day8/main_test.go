package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	route, allDirections := parse("adventday8.txt")

	steps := endLocation("AAA", allDirections, route, "ZZZ", 0)

	assert.Equal(11309, steps, "they should be equal")
}

func TestGetAllStarts(t *testing.T) {
	assert := assert.New(t)
	_, allDirections := parse("adventday8.txt")
	allStarts := getAllStarts(allDirections)

	assert.Equal(6, len(allStarts), "they should be equal")
}

func TestELP2(t *testing.T) {
	// assert := assert.New(t)
	part2()

	// assert.Equal(6, len(allStarts), "they should be equal")
}
