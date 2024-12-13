package main

import (
	"fmt"
	"os"
	"strconv"
)

func readFileIntoString(fileName string) string {
	bytes, _ := os.ReadFile(fileName)

	return string(bytes)
}

func mapContentsIntoStorage(contents string) []int {
	storage := []int{}

	id := 0
	// 12345
	// 0 .. 111 .... 22222
	for pos, r := range contents {
		num, _ := strconv.Atoi(fmt.Sprintf("%c", r))
		for i := 0; i < num; i++ {
			if pos%2 == 0 {
				storage = append(storage, id)
			} else {
				storage = append(storage, -1)
			}
		}
		if pos%2 == 0 {
			id++
		}
	}

	return storage
}

func findIndex(str []int, direction int, location int) int {
	if location < 0 || location >= len(str) {
		return -1
	}

	if direction < 0 && str[location] != -1 {

		return location
	} else if direction > 0 && str[location] == -1 {
		return location
	}

	return findIndex(str, direction, location+direction)
}

func compactFile(storage []int) []int {
	// find index, going backward and forward
	untilCopmlete := true

	dotMemory := 0
	numMemory := len(storage) - 1

	for untilCopmlete {
		dotLocation := findIndex(storage, 1, dotMemory)
		numLocation := findIndex(storage, -1, numMemory)

		dotMemory = dotLocation
		numMemory = numLocation
		if dotMemory > numMemory {
			break
		}

		storage[dotLocation], storage[numLocation] = storage[numLocation], storage[dotLocation]
	}

	return storage
}

func calculateChecksum(file []int) int {
	checksum := 0

	for pos, r := range file {
		if r != -1 {
			checksum += r * pos
		}

	}

	return checksum
}

func main() {
	contents := readFileIntoString("input.txt")
	storage := mapContentsIntoStorage(contents)
	file := compactFile(storage)
	checksum := calculateChecksum(file)
	fmt.Println(contents, "***", storage, "****", file, checksum)

}
