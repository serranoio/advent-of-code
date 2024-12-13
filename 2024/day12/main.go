package main

import (
	"fmt"
	"os"
	"strings"
)

type Location struct {
	x int
	y int
}

type Region struct {
	locations []*Location
	r         rune
}

func readFileIntoGarden(fileName string) [][]rune {
	bytes, _ := os.ReadFile(fileName)

	var garden [][]rune
	for _, line := range strings.Split(string(bytes), "\n") {
		var gardenLine []rune
		for _, r := range line {
			gardenLine = append(gardenLine, r)
		}
		garden = append(garden, gardenLine)
	}

	return garden
}

func regionAlreadyHasLocation(region *Region, location Location) bool {
	for _, regionLocation := range region.locations {
		if regionLocation.y == location.y && regionLocation.x == location.x {
			return true
		}
	}

	return false
}

func NewRegion(r rune) *Region {
	return &Region{
		locations: []*Location{},
		r:         r,
	}
}

func (r *Region) calculateArea() int {
	return len(r.locations)
}

func isInArea(location Location, locations []*Location) bool {
	for _, regionLocation := range locations {
		if regionLocation.y == location.y && location.x == regionLocation.x {
			return true
		}
	}

	return false
}

func (r *Region) calculatePerimiter() int {
	count := 0

	for _, c := range r.getPerimiter() {
		count += c
	}

	return count
}

func sideDirection(side []*Location, location Location) bool {
	if len(side) == 1 {
		return true
	}

	if side[1].x == side[0].x && location.x == side[0].x {
		return true
	}

	if side[1].y == side[0].y && location.y == side[0].y {
		return true
	}

	return false
}

func isLocationInSide(side []*Location, location Location) bool {
	for _, sideLocation := range side {
		// calculate direction && is in direction

		if sideLocation.x == location.x && location.y == sideLocation.y {
			return false
		}
	}

	return true
}

func (r *Region) getSides() [][]*Location {
	perimeterLocations := r.getPerimiter()
	sides := [][]*Location{}

	for location := range perimeterLocations {

		// if point is linear to any points in a side
		// if point makes up new side
		needsNewSide := true

		for k, side := range sides {
			if sideDirection(side, location) {
				sides[k] = append(sides[k], &location)
				needsNewSide = false
				break
			}

			if isLocationInSide(side, location) {
				needsNewSide = false
			}
		}

		if needsNewSide {
			sides = append(sides, []*Location{&location})
		}
	}

	return sides
}

func (r *Region) getPerimiter() map[Location]int {
	perimeterLocations := make(map[Location]int)

	for _, location := range r.locations {
		if !isInArea(Location{x: location.x, y: location.y + 1}, r.locations) {
			perimeterLocations[Location{location.x, location.y + 1}] += 1
		}
		if !isInArea(Location{x: location.x, y: location.y - 1}, r.locations) {
			perimeterLocations[Location{location.x, location.y - 1}] += 1
		}
		if !isInArea(Location{x: location.x - 1, y: location.y}, r.locations) {
			perimeterLocations[Location{location.x - 1, location.y}] += 1
		}
		if !isInArea(Location{x: location.x + 1, y: location.y}, r.locations) {
			perimeterLocations[Location{location.x + 1, location.y}] += 1
		}
	}

	return perimeterLocations
}

func locationOnBoard(garden [][]rune, location Location) bool {
	if location.x >= 0 && location.y >= 0 && location.x < len(garden[0]) && location.y < len(garden) {
		return true
	}
	return false
}

func searchForRegion(region *Region, garden [][]rune, currentLocation Location) {
	if !locationOnBoard(garden, currentLocation) {
		return
	}

	if !regionAlreadyHasLocation(region, currentLocation) && region.r == garden[currentLocation.y][currentLocation.x] {
		region.locations = append(region.locations, &currentLocation)
	} else {
		return
	}

	searchForRegion(region, garden, Location{x: currentLocation.x + 1, y: currentLocation.y})
	searchForRegion(region, garden, Location{x: currentLocation.x - 1, y: currentLocation.y})
	searchForRegion(region, garden, Location{x: currentLocation.x, y: currentLocation.y + 1})
	searchForRegion(region, garden, Location{x: currentLocation.x, y: currentLocation.y - 1})
}

func locationIsInAnyRegion(location Location, regions []*Region) bool {
	for _, region := range regions {
		for _, regionLocation := range region.locations {
			if regionLocation.y == location.y && regionLocation.x == location.x {
				return true
			}
		}
	}

	return false
}

func lookForRegions(garden [][]rune) []*Region {
	regions := []*Region{}
	for y, line := range garden {
		for x, r := range line {
			if !locationIsInAnyRegion(Location{y: y, x: x}, regions) {
				region := NewRegion(r)

				searchForRegion(region, garden, Location{y: y, x: x})
				regions = append(regions, region)
			}

		}
	}
	return regions
}

func calculateCost(regions []*Region) int {
	cost := 0
	for _, region := range regions {
		perim := region.calculatePerimiter()
		area := region.calculateArea()

		cost += area * perim
	}

	return cost
}

func main() {
	garden := readFileIntoGarden("input.txt")

	regions := lookForRegions(garden)

	cost := calculateCost(regions)

	fmt.Println(cost)

}
