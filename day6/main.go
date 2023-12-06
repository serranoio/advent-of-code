package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func splitFile(name string) []string {
	bytes, _ := os.ReadFile(name)

	return strings.Split(string(bytes), "\n")
}

func convertRowToIntArray(row string) []int {
	numbers := []int{}
	curString := ""
	for _, aRune := range row {
		if aRune == ' ' {
			number, err := strconv.Atoi(curString)
			if err == nil {
				numbers = append(numbers, number)
			}
			curString = ""
		} else {
			curString += string(aRune)
		}
	}

	number, err := strconv.Atoi(curString)
	if err == nil {
		numbers = append(numbers, number)
	}

	return numbers
}

func containsOdd(raceMaxes []int) int {
	for _, max := range raceMaxes {
		if max%2 != 0 {
			return 1
		}
	}
	return 0
}

// do odd numbers cause an odd number of races?

func testRaces(times []int, distances []int) []int {
	allRaceMaxes := []int{}

	for raceNumber := 0; raceNumber < len(times); raceNumber++ {
		raceMaxes := []int{}
		raceTime := times[raceNumber]
		raceDistance := distances[raceNumber]

		for holdTime := 1; holdTime < (raceTime+1)/2; holdTime++ {
			// if we hold one second, the boat will go 1 miles per hour
			// therefore holdTime = speed
			// rest of the duration is the
			timeLeft := raceTime - holdTime

			distanceTraveled := holdTime * timeLeft

			if distanceTraveled > raceDistance {
				raceMaxes = append(raceMaxes, distanceTraveled)
			}
		}

		addOne := containsOdd(raceMaxes)

		allRaceMaxes = append(allRaceMaxes, len(raceMaxes)*2+addOne)
	}

	return allRaceMaxes
}

func joinIntArray(array []int) int {
	allInts := ""
	for _, num := range array {
		allInts = fmt.Sprintf("%s%d", allInts, num)
	}

	num, _ := strconv.Atoi(allInts)

	return num
}

func testPartOne(partOne bool, name string) {
	split := splitFile(name)

	times := convertRowToIntArray(split[0])
	distances := convertRowToIntArray(split[1])

	if !partOne {
		// then they stay the same
		joinedTimesArray := joinIntArray(times)
		times = []int{joinedTimesArray}

		joinedDistancesArray := joinIntArray(distances)
		distances = []int{joinedDistancesArray}
	}

	allRaceMaxes := testRaces(times, distances)

	if partOne {
		multipliedJuntos := 1
		for _, maxRaces := range allRaceMaxes {
			multipliedJuntos *= maxRaces
		}

		log.Printf("All multiplied together %d", multipliedJuntos)
	} else {
		// part two -> show the amount of ways you can beat the record
		log.Printf("Ways: %d", allRaceMaxes[0])
	}
}

func main() {
	testPartOne(false, "adventday6.txt")
}
