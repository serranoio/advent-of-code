package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitFile(t *testing.T) {
	assert := assert.New(t)

	split := splitFile("test.txt")

	assert.Equal(len(split), 33, "they should be equal")
}

const seedsString = "seeds: 79 14 55 13\n"

func TestMapFileToSections(t *testing.T) {
	assert := assert.New(t)

	allMaps := mapFileToSections()

	assert.Equal(len(allMaps), 7, "they should be equal")
	assert.Equal(allMaps[0], seedsString, "they should be equal")
}

func TestGetAllSeeds(t *testing.T) {
	assert := assert.New(t)

	allSeeds := getAllSeeds(seedsString)

	assert.Equal(len(allSeeds), 4, "they should be equal")
}

func TestCreateStructsFromSections(t *testing.T) {
	// assert := assert.New(t)

	// allSeeds := getAllSeeds(seedsString)
}

func TestCreateSTDName(t *testing.T) {
	assert := assert.New(t)
	std := initSTD()

	std.createSTDNames("seed-to-soil map:")

	assert.Equal(std.sourceName, "soil", "they should be equal")
	assert.Equal(std.destinationName, "seed", "they should be equal")
}

func TestCreateRainj(t *testing.T) {
	assert := assert.New(t)
	std := initSTD()

	std.rainj = 2
	std.sourceStart = 98
	std.destinationStart = 50
	std.calculateRainj()

	assert.Equal(std.stdRange[98], 50, "they should be equal")
	assert.Equal(std.stdRange[99], 51, "they should be equal")
}
