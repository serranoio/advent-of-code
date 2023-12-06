package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllContents(t *testing.T) {
	assert := assert.New(t)

	split := allContentsSplit("test.txt")

	assert.Equal(len(split), 6, "they should be equal")
}

// successfully construct the hands
func TestCreateHandsIntArray(t *testing.T) {
	assert := assert.New(t)

	allCardHands := createHands(allContentsSplit("test.txt"))

	card := allCardHands[0].card
	hand := allCardHands[0].hand

	assert.Equal(len(card), 5, "they should be equal")
	assert.Equal(len(hand), 8, "they should be equal")
}

// func TestCalculateHands(t *testing.T) {
// 	assert := assert.New(t)

// 	allCardHands := createHands(allContentsSplit("test.txt"))

// 	assert.Equal(allCardHands[0].matchCount, 8, "they should be equal")
// 	assert.Equal(allCardHands[1].matchCount, 2, "they should be equal")
// 	assert.Equal(allCardHands[2].matchCount, 2, "they should be equal")
// 	assert.Equal(allCardHands[3].matchCount, 1, "they should be equal")
// 	assert.Equal(allCardHands[4].matchCount, 0, "they should be equal")
// 	assert.Equal(allCardHands[5].matchCount, 0, "they should be equal")
// }

// func TestTotal(t *testing.T) {
// 	assert := assert.New(t)

// 	allCardHands := createHands(allContentsSplit("test.txt"))
// 	ah := initACH(allCardHands)
// 	ah.calcTotalPart1()

// 	assert.Equal(ah.total, 9, "they should be equal")
// }

func TestCopyCopy(t *testing.T) {
	assert := assert.New(t)

	allCardHands := createHands(allContentsSplit("test.txt"))
	ah := initACH(allCardHands)

	ah.calcTotalPart2()

	assert.Equal(30, ah.total, "equal")
}
