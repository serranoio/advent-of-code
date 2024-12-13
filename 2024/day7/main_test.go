package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {

	combinations := generateCombinations(3, true)
	fmt.Println(combinations)

	combinations = generateCombinations(8, true)
	fmt.Println(combinations)
}

func TestConcatenateNumbers(t *testing.T) {

	newNum := concatenateNumbers(12, 5)

	assert.Equal(t, 125, newNum)
}

func TestMain(t *testing.T) {

	main()
}
