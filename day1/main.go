package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var allNumbers = [10]string{"1", "0", "2", "3", "4", "5", "6", "7", "8", "9"}

var allNumbersMap = map[string]string{}

func getNumber(num string, myMap map[string]string) string {
	if myMap[num] == "" {
		return num
	}
	return myMap[num]
}

func convertToTwoDigit(number []string, myMap map[string]string) string {
	if len(number) == 0 {
		return ""
	}

	if len(number) == 1 {
		return getNumber(number[0], myMap) + "" + getNumber(number[0], myMap)
	}

	num := getNumber(number[0], myMap) + "" + getNumber(number[len(number)-1], myMap)

	return num
}

func initMap() map[string]string {
	return map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}
}

func main() {
	bytes, _ := os.ReadFile("advent1.txt")

	myMap := initMap()

	var doubleArray [][]string

	for _, line := range strings.Split(string(bytes), "\n") {
		// numbers := ""

		var lineArray []string

		for i, _ := range line {
			for k, v := range myMap {
				if string(line[i]) == v {
					lineArray = append(lineArray, string(line[i]))
					continue
				}

				notGood := false
				for curLetter, letterInMap := range k {
					if i+curLetter >= len(line) {
						notGood = true
						break
					}

					if !(string(line[i+curLetter]) == string(letterInMap)) {
						notGood = true
						break
					}
					// get current place in line
					// make sure theyre all equal to the letters in here
				}
				if !notGood {
					lineArray = append(lineArray, k)
				}
			}
			// END THIS
		}

		doubleArray = append(doubleArray, lineArray)
	}

	counter := 0
	for _, row := range doubleArray {
		twoDigits := convertToTwoDigit(row, myMap)

		num, _ := strconv.Atoi(twoDigits)
		counter += num
	}

	fmt.Println(counter)

}
