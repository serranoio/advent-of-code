package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	contents := readFile("test.txt")
	pageOrderingRules, _ := splitFile(contents)

	val := fixList(pageOrderingRules, []int{47, 75, 97, 29, 13})

	assert.Equal(t, val, 47, "they should be equal")

}
