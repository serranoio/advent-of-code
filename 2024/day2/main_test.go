package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSingleReportWithProblemDampener(t *testing.T) {

	isValid := checkSingleReportWithProblemDampener([]int{1, 3, 2, 4, 5})
	assert.Equal(t, isValid, 1, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{1, 2, 7, 8, 9})
	assert.Equal(t, isValid, 0, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{7, 6, 4, 2, 1})
	assert.Equal(t, isValid, 1, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{9, 7, 6, 2, 1})
	assert.Equal(t, isValid, 0, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{10, 6, 4, 3, 2})
	assert.Equal(t, isValid, 1, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{8, 6, 7, 6, 5})
	assert.Equal(t, isValid, 1, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{7, 5, 6, 5, 4})
	assert.Equal(t, isValid, 1, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{7, 6, 5, 4, 0})
	assert.Equal(t, isValid, 1, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{7, 6, 5, 1, 0})
	assert.Equal(t, isValid, 0, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{8, 9, 10, 2, 1})
	assert.Equal(t, isValid, 0, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{1, 5, 6, 7, 8})
	assert.Equal(t, isValid, 1, "they should be equal")

	isValid = checkSingleReportWithProblemDampener([]int{75, 74, 75, 78, 80, 83})
	assert.Equal(t, isValid, 1, "they should be equal")

	// numbers := []struct {
	// 	input    []int
	// 	expected int
	// }{
	// 	{[]int{35, 32, 31, 28, 25}, 1},
	// 	{[]int{79, 78, 75, 72, 69}, },
	// 	{[]int{33, 35, 37, 38, 39, 42, 44}, 0},
	// 	{[]int{22, 21, 19, 18, 15, 12, 9, 8}, 0},
	// 	{[]int{66, 65, 64, 63, 62, 61, 59, 58}, 0},
	// 	{[]int{88, 85, 84, 81, 79, 76, 74}, 0},
	// 	{[]int{73, 71, 68, 66, 65, 62}, 0},
	// 	{[]int{66, 67, 69, 71, 73}, 0},
	// 	{[]int{51, 49, 47, 46, 45, 44}, 0},
	// 	{[]int{29, 26, 24, 21, 18, 17, 16}, 0},
	// }
	// for _, tc := range numbers {
	// 	isValid = checkSingleReportWithProblemDampener(tc.input)
	// 	assert.Equal(t, isValid, 0, "they should be equal")

	// }

}

func TestCheckSingleReport(t *testing.T) {

	isValid := checkSingleReport([]int{1, 3, 5, 7, 9})
	assert.Equal(t, isValid, 1, "they should be equal")

	isValid = checkSingleReport([]int{1, 3, 5, 7, 11})
	assert.Equal(t, isValid, 0, "they should be equal")

}

func TestPart1(t *testing.T) {
	answer := part1("test.txt")
	assert.Equal(t, answer, 2, "they should be equal")
}

func TestPart2(t *testing.T) {
	part2("input.txt")
	// assert.Equal(t, answer, 4, "they should be equal")
}
