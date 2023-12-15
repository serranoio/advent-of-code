package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindReflectionHEdge(t *testing.T) {
	assert := assert.New(t)

	groups := parseFile("adventday13.txt")

	num := findMirrors(groups[0], 0, 0)

	assert.Equal(0, num)
}
func TestFindReflectionH(t *testing.T) {
	assert := assert.New(t)

	reflectionPoint := 2
	groups := parseFile("test.txt")

	_, num := findReflectionH(groups[0], reflectionPoint, reflectionPoint+1, 0)
	assert.Equal(num, 1)

	// reflectionPoint
	_, num = findReflectionH(groups[2], reflectionPoint, reflectionPoint+1, 0)
	assert.Equal(num, 2)

	reflectionPoint = 4
	_, num = findReflectionH(groups[1], reflectionPoint, reflectionPoint+1, 0)
	assert.Equal(num, 4)
	// reflectionPoint = 2
	// isReflection = findReflectionV("#.##..##.", reflectionPoint, reflectionPoint+1)
	// assert.Equal(isReflection, false)

}
func TestFindReflectionV(t *testing.T) {
	assert := assert.New(t)

	reflectionPoint := 4

	isReflection, num := findReflectionV("#.##.###.", reflectionPoint, reflectionPoint+1, 0)

	assert.Equal(num, 1)
	assert.Equal(isReflection, true)

	reflectionPoint = 3
	_, num = findReflectionV("#.##..##.", reflectionPoint, reflectionPoint+1, 0)
	assert.Greater(num, 1)

	reflectionPoint = 2
	_, num = findReflectionV("#.##..##.", reflectionPoint, reflectionPoint+1, 0)
	assert.Equal(num, 1)

	// println(isReflection)
}

func TestFindMirrors(t *testing.T) {
	assert := assert.New(t)

	groups := parseFile("test.txt")

	count := findMirrors(groups[0], 0, 0)
	assert.Equal(300, count)

	count = findMirrors(groups[1], 0, 0)
	assert.Equal(100, count)

	count = findMirrors(groups[4], 0, 0)
	assert.Equal(100, count)

	count = findMirrors(groups[5], 0, 0)
	assert.Equal(100, count)

	// count := findMirrors(groups[3], 0, 0)
	// assert.Equal(400, count)
}

func TestDoc(t *testing.T) {
	assert := assert.New(t)

	groups := parseFile("adventday13.txt")

	sum := getSum(groups)
	assert.Equal(0, sum)
}

// if a row is off by MORE Than one
// create an off counter
