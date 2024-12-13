package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCoordinateCombinations(t *testing.T) {

	resultsChan := generateCoordinateCombinations([]Coords{
		{8, 1}, {5, 2}, {7, 3}, {4, 4},
	})

	count := 0
	for range resultsChan {
		count++
	}

	assert.Equal(t, count, 6, "WOw")

}

func TestMain(t *testing.T) {
	main()

}
