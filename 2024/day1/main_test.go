package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	totalDistance := part1("test.txt")

	assert.Equal(t, 11, totalDistance, "HELL YEAH BABY")
}

func TestPart2(t *testing.T) {
	totalSimilarity := part2("test.txt")

	assert.Equal(t, 31, totalSimilarity, "HELL YEAH BABY")
}

func TestSplitStringInTwo(t *testing.T) {

	str := "293847 2039487247"

	strs := splitStringInTwo(str)

	fmt.Println(strs)

	assert.Equal(t, strs[0], "293847", "HELL YEAH BABY")
	assert.Equal(t, strs[1], "2039487247", "HELL YEAH BABY")

}
