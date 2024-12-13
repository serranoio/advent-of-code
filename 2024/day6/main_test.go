package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	main()

}

func TestIsRectangle(t *testing.T) {
	contents := readFile("input.txt")
	createBoard(contents)

	locations := []location{
		{4, 1},
		{8, 1},
		{8, 6},
	}

	rectangle := isRectangle(locations)

	fmt.Println(rectangle)

}
