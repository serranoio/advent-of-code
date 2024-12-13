package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/schwarmco/go-cartesian-product"
)

type Operator rune

const (
	ADD         Operator = '+'
	MULTIPLY    Operator = '*'
	CONCATENATE Operator = '|'
)

func readFile(fileName string) string {
	bytes, _ := os.ReadFile(fileName)

	return string(bytes)
}

var allCombinationsMap map[int][][]Operator

func generateCombinations(numLength int, part1 bool) [][]Operator {
	operatorCount := numLength - 1

	if allCombinations, ok := allCombinationsMap[operatorCount]; ok {
		return allCombinations
	}

	allCombinations := [][]interface{}{}

	for i := 0; i < operatorCount; i++ {
		var combination []interface{}
		if part1 {
			combination = []interface{}{ADD, MULTIPLY}
		} else {
			combination = []interface{}{ADD, MULTIPLY, CONCATENATE}
		}

		allCombinations = append(allCombinations, combination)
	}

	allProducts := [][]Operator{}
	products := cartesian.Iter(allCombinations...)
	for product := range products {
		operators := []Operator{}
		for _, num := range product {
			operator, _ := num.(Operator)
			operators = append(operators, operator)
		}
		allProducts = append(allProducts, operators)
	}

	allCombinationsMap[operatorCount] = allProducts

	return allProducts
}

func concatenateNumbers(num1, num2 int) int {
	newNum, _ := strconv.Atoi(fmt.Sprintf("%d%d", num1, num2))

	return newNum
}

func dropKickCombinations(total int, combinationStack []Operator, numsLeftInEquation []int, currentTotal int) int {
	operator := combinationStack[0]
	num := numsLeftInEquation[0]

	if operator == Operator(MULTIPLY) {
		currentTotal *= num
	} else if operator == Operator(ADD) {
		currentTotal += num
	} else {
		currentTotal = concatenateNumbers(currentTotal, num)
	}

	if currentTotal == total {
		return total
	}

	if len(numsLeftInEquation) == 1 {
		return 0
	}

	combinationStack = combinationStack[1:]
	numsLeftInEquation = numsLeftInEquation[1:]

	return dropKickCombinations(total, combinationStack, numsLeftInEquation, currentTotal)
}

func calculateCalibrationResult(total int, nums []int, part1 bool) int {
	combinations := generateCombinations(len(nums), part1)

	for _, combination := range combinations {
		calibrationResult := dropKickCombinations(total, combination, nums[1:], nums[0])
		if calibrationResult != 0 {
			return calibrationResult
		}
	}

	return 0
}

func constructEquations(contents string, part1 bool) int {

	totalCalibrationResult := 0

	for _, line := range strings.Split(contents, "\n") {
		bothSides := strings.Split(line, ":")
		totalString := bothSides[0]
		total, _ := strconv.Atoi(totalString)
		numsString := bothSides[1]
		numsString = numsString[1:]
		nums := []int{}
		for _, numString := range strings.Split(numsString, " ") {
			num, _ := strconv.Atoi(numString)
			nums = append(nums, num)
		}

		totalCalibrationResult += calculateCalibrationResult(total, nums, part1)
	}

	return totalCalibrationResult
}

func main() {
	contents := readFile("input.txt")

	allCombinationsMap = make(map[int][][]Operator)
	result := constructEquations(contents, true)
	fmt.Println(result)
	allCombinationsMap = make(map[int][][]Operator)
	result = constructEquations(contents, false)
	fmt.Println(result)

}
