package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	main()
}

func TestCalculatePerimeter(t *testing.T) {
	r := &Region{
		locations: []*Location{{x: 2, y: 1}, {x: 2, y: 2}, {x: 3, y: 2}, {x: 3, y: 3}},
		r:         'C',
	}

	perim := r.calculatePerimiter()

	assert.Equal(t, perim, 10)

	garden := readFileIntoGarden("test2.txt")
	regions := lookForRegions(garden)
	cost := calculateCost(regions)

	assert.Equal(t, cost, 772)

	garden = readFileIntoGarden("test3.txt")
	regions = lookForRegions(garden)
	cost = calculateCost(regions)

	assert.Equal(t, cost, 1930)

}

func TestGetSides(t *testing.T) {
	garden := readFileIntoGarden("test.txt")
	regions := lookForRegions(garden)

	sides := regions[0].getSides()

	fmt.Println(sides)
	// assert.Equal(t, cost, 772)

}
