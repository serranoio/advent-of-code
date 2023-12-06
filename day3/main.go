package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	PERIOD = iota
	NUMBER
	SYMBOL
)

type space struct {
	content int
	number  string
}

var allNumbers = [10]string{"1", "0", "2", "3", "4", "5", "6", "7", "8", "9"}
var numberTriggered bool

func includes(char string) bool {
	for _, number := range allNumbers {
		if number == char {
			return true
		}
	}
	return false
}

var gear map[int]map[int]int

func fillSpaces(spaces [][]*space) [][]*space {
	contents, _ := os.ReadFile("test.txt")

	for _, line := range strings.Split(string(contents), "\n") {
		spaceRow := []*space{}
		for _, char := range line {
			spaceContent := space{}
			if char == '.' {
				spaceContent.content = PERIOD
			} else if includes(string(char)) {
				spaceContent.content = NUMBER
				spaceContent.number = string(char)
			} else {
				spaceContent.number = string(char)
				spaceContent.content = SYMBOL
			}
			spaceRow = append(spaceRow, &spaceContent)
		}
		spaces = append(spaces, spaceRow)
	}

	return spaces
}

func containsGear(spaces [][]*space, row int, column int) {
	if spaces[row][column].number != "*" || numberTriggered {
		return
	}

	if rowMap, ok := gear[row]; !ok {
		newMap := make(map[int]int)
		gear[row] = newMap

		if _, ok := newMap[column]; ok {
			newMap[column] += 1
		} else {
			newMap[column] = 1
		}
	} else {
		if _, ok := rowMap[column]; ok {
			rowMap[column] += 1
		} else {
			rowMap[column] = 1
		}
	}

	numberTriggered = false
}

func isAdjacent(spaces [][]*space, row int, column int) bool {
	if row-1 > 0 {
		if spaces[row-1][column].content == SYMBOL {
			containsGear(spaces, row-1, column)
			return true
		}
	}
	if row-1 > 0 && column-1 > 0 {
		if spaces[row-1][column-1].content == SYMBOL {
			containsGear(spaces, row-1, column-1)
			return true
		}
	}
	if row-1 > 0 && column+1 < len(spaces) {
		if spaces[row-1][column+1].content == SYMBOL {
			containsGear(spaces, row-1, column+1)
			return true
		}
	}

	if column-1 > 0 {
		if spaces[row][column-1].content == SYMBOL {
			containsGear(spaces, row, column-1)
			return true
		}
	}
	if column+1 < len(spaces) {
		if spaces[row][column+1].content == SYMBOL {
			containsGear(spaces, row, column+1)
			return true
		}
	}

	if column+1 < len(spaces) && row+1 < len(spaces) {
		if spaces[row+1][column+1].content == SYMBOL {
			containsGear(spaces, row+1, column+1)
			return true
		}
	}
	if row+1 < len(spaces) {
		if spaces[row+1][column].content == SYMBOL {
			containsGear(spaces, row+1, column)
			return true
		}
	}
	if column-1 > 0 && row+1 < len(spaces) {
		if spaces[row+1][column-1].content == SYMBOL {
			containsGear(spaces, row+1, column-1)
			return true
		}
	}

	return false
}

func main() {
	spaces := [][]*space{}
	spaces = fillSpaces(spaces)
	numberList := []string{}
	gear = make(map[int]map[int]int)
	numberTriggered = false

	for row := range spaces {

		newNumber := ""
		startNumber := false
		numberIsAdjacentToSymbol := false
		for column := range spaces[row] {
			if spaces[row][column].content == NUMBER {
				startNumber = true
				numberTriggered = startNumber
				newNumber += spaces[row][column].number
				if isAdjacent(spaces, row, column) {
					numberIsAdjacentToSymbol = true
				}
			} else {
				if startNumber {
					if numberIsAdjacentToSymbol {
						numberList = append(numberList, newNumber)
						numberIsAdjacentToSymbol = false
					}
					numberTriggered = startNumber
					startNumber = false
					newNumber = ""
				}
			}

		}
		if numberIsAdjacentToSymbol {
			numberList = append(numberList, newNumber)
			numberIsAdjacentToSymbol = false
			startNumber = false
			newNumber = ""
		}
		fmt.Println("\n")
	}

	gearMap := gear

	fmt.Println(gearMap)

	for _, rowMap := range gear {
		for _, val := range rowMap {
			fmt.Println(val)
		}
	}

}
