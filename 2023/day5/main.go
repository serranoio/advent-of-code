package main

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/charmbracelet/log"
)

func splitFile(nameOfFile string) []string {
	bytes, _ := os.ReadFile(nameOfFile)

	return strings.Split(string(bytes), "\n")
}

type STD struct {
	rainj            []int
	sourceStart      []int
	destinationStart []int
}

func initSTD() *STD {
	return &STD{
		rainj:            []int{},
		sourceStart:      []int{},
		destinationStart: []int{},
	}
}

var seeds []int

func mapFileToSections() []string {
	allSections := []string{}

	mapString := ""
	for _, line := range splitFile("adventday5.txt") {
		if line == "" {
			allSections = append(allSections, mapString)
			mapString = ""
		}
		mapString += line + "\n"
	}
	allSections = append(allSections, mapString)

	return allSections
}

func getAllSeeds(row string) []int {
	allSeeds := []int{}
	onlySeeds := strings.Split(row, ":")[1]
	for _, seed := range strings.Split(onlySeeds, " ") {
		seedString := string(seed)
		if seedString == "" {
			continue
		}
		seedNum, _ := strconv.Atoi(strings.TrimSpace(seed))
		allSeeds = append(allSeeds, seedNum)
	}

	return allSeeds
}

// func (s *STD) createSTDNames(name string) {
// 	allWords := strings.Split(name, "-")

// 	s.destinationName = allWords[0]
// 	s.sourceName = strings.Split(allWords[2], " ")[0]
// }

func initSection(allSections []string, mapNumber int) (*STD, []string) {
	std := initSTD()
	section := allSections[mapNumber]
	partsOfSection := strings.Split(section, "\n")
	partsOfSection = partsOfSection[1 : len(partsOfSection)-1]
	// name := partsOfSection[0]
	// std.createSTDNames(name)

	return std, partsOfSection
}

func stringToNumber(number string) int {
	num, _ := strconv.Atoi(number)
	return num
}

// SEEDS ON RIGHT. SEEDS = SOURCE
// SOIL ON LEFT. SOIL = DESTINATION

func createStructsFromSections(allSections []string) []*STD {
	seeds = getAllSeeds(allSections[0])

	allSTDs := []*STD{}

	for mapNumber := 1; mapNumber < len(allSections); mapNumber++ {
		std, partsOfSection := initSection(allSections, mapNumber)
		// all numbers
		for partsNumbers := 1; partsNumbers < len(partsOfSection); partsNumbers++ {
			rangeLine := partsOfSection[partsNumbers]
			allNumbers := strings.Split(rangeLine, " ")

			std.rainj = append(std.rainj, stringToNumber(allNumbers[2]))
			std.sourceStart = append(std.sourceStart, stringToNumber(allNumbers[1]))
			std.destinationStart = append(std.destinationStart, stringToNumber(allNumbers[0]))
			// std.calculateRainj()
		}
		allSTDs = append(allSTDs, std)
	}
	return allSTDs
}

//	func (s *STD) calculateRainj() {
//		for mapping := 0; mapping < s.rainj; mapping++ {
//			s.stdRange[mapping+s.sourceStart] = mapping + s.destinationStart
//		}
//	}
func withinRange(std *STD, curSeed int) int {

	for nums := 0; nums < len(std.rainj); nums++ {
		// if its within here
		if curSeed >= std.sourceStart[nums] && curSeed < std.sourceStart[nums]+std.rainj[nums] {
			// then we try to map it to a destinationStart
			curSeed = curSeed - std.sourceStart[nums] + std.destinationStart[nums]
			break
		}
	}

	return curSeed
}

func greatness(allSTDs []*STD, seeds int) int {

	for _, std := range allSTDs {
		// this std maps to the next std
		seeds = withinRange(std, seeds)
	}

	return seeds
}

// Part 1: 621354867
func convertRange() []int {
	newRange := []int{}
	seeds := seeds
	for num := 0; num < len(seeds)-2; num += 2 {
		if seeds[num+2] < seeds[num]+seeds[num+1] {
			newRange = append(newRange, seeds[num])
			newRange = append(newRange, seeds[num+2]+seeds[num+3])
		}
	}

	return newRange
}

func execute() {
	allSTDs := createStructsFromSections(mapFileToSections())

	seeds = convertRange()
	seeds = convertRange()
	seeds = convertRange()

	allLocations := []int{}

	var wg sync.WaitGroup

	var lock sync.Mutex

	for num := 0; num < len(seeds); num += 2 {
		wg.Add(1)
		go func(num int) {
			log.Info("Current Seed", seeds[num])
			maxRange := seeds[num] + seeds[num+1]
			for seedsRange := seeds[num]; seedsRange < maxRange; seedsRange += 1 {
				value := greatness(allSTDs, seedsRange)
				lock.Lock()
				allLocations = append(allLocations, value)
				lock.Unlock()
			}

			log.Printf("Seed: %d Done", seeds[num])
			wg.Done()
		}(num)

	}
	wg.Wait()

	log.Info("Calculate min")
	curMinimum := 9999999999999999
	for _, location := range allLocations {
		if location < curMinimum {
			curMinimum = location
		}
	}

	log.Info(curMinimum)

}

func main() {

	execute()
}

// destination source range
// soil seed range
// seed-to-soil

// read from right to left
