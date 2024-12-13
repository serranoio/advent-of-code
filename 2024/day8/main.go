package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ernestosuarez/itertools"
)

type Coords struct {
	x int
	y int
}

var width int
var height int

func readFileIntoCombinations(fileName string) map[rune][]Coords {
	antennaLocations := make(map[rune][]Coords)

	bytes, _ := os.ReadFile(fileName)

	heightCount := 0
	widthCount := 0
	for y, line := range strings.Split(string(bytes), "\n") {
		heightCount++
		for x, r := range line {
			if y == 0 {
				widthCount++
			}
			if r == '.' {
				continue
			}
			if _, ok := antennaLocations[r]; ok {
				antennaLocations[r] = append(antennaLocations[r], Coords{x: x, y: y})
			} else {
				antennaLocations[r] = []Coords{{x: x, y: y}}
			}
		}
	}

	width = widthCount
	height = heightCount

	return antennaLocations
}

func generateCoordinateCombinations(coords []Coords) chan itertools.List {
	result := []interface{}{}

	for _, v := range coords {
		result = append(result, v)
	}

	results := itertools.CombinationsList(result, 2)

	return results
}

func calculateAntinodes(coordianteCombinations chan itertools.List) []Coords {
	antinodeCoordinates := []Coords{}

	for coords := range coordianteCombinations {
		coord1 := coords[0].(Coords)
		coord2 := coords[1].(Coords)

		yVec := coord2.y - coord1.y
		xVec := coord2.x - coord1.x

		// 2 to 1 is positive, so adding onto 1 would continue to make it positive
		antinodeCoordinates = append(antinodeCoordinates, Coords{
			x: coord2.x + xVec,
			y: coord2.y + yVec,
		})
		antinodeCoordinates = append(antinodeCoordinates, Coords{
			x: coord1.x - xVec,
			y: coord1.y - yVec,
		})
	}

	return antinodeCoordinates
}

func verifyCoordinatesAreOnBoard(antinodeCoordinates []Coords) int {
	anitnodeCount := 0
	for _, coord := range antinodeCoordinates {
		if coord.x >= 0 && coord.x < width && coord.y >= 0 && coord.y < height {
			anitnodeCount++
		}
	}

	return anitnodeCount
}

func removeDuplicates(allAntinodeCoordinates []Coords) []Coords {
	inResult := make(map[Coords]bool)
	results := []Coords{}

	for _, coord := range allAntinodeCoordinates {
		if _, ok := inResult[coord]; !ok {
			results = append(results, coord)
			inResult[coord] = true
		}

	}

	return results

}

func countAntinodes(locations map[rune][]Coords) int {

	allAntinodeCoordinates := []Coords{}

	for _, coords := range locations {
		coordianteCombinations := generateCoordinateCombinations(coords)
		anitnodeCoordinates := calculateAntinodes(coordianteCombinations)
		allAntinodeCoordinates = append(allAntinodeCoordinates, anitnodeCoordinates...)
	}

	allAntinodeCoordinates = removeDuplicates(allAntinodeCoordinates)

	return verifyCoordinatesAreOnBoard(allAntinodeCoordinates)
}

func main() {
	locations := readFileIntoCombinations("input.txt")
	count := countAntinodes(locations)

	fmt.Println(count)

}
