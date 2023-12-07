package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newCamel(hand string) *Camel {
	camel := &Camel{
		bid:  1,
		hand: hand,
	}

	camel.calculateHandValue()
	return camel
}

func TestOne(t *testing.T) {
	assert := assert.New(t)

	camels := []*Camel{
		newCamel("22345"),
		newCamel("32345"),
		newCamel("42345"),
	}

	doSort(camels)

	assert.Equal("42345", camels[0].hand, "they should be equal")
	assert.Equal("32345", camels[1].hand, "they should be equal")
	assert.Equal("22345", camels[2].hand, "they should be equal")
}

func TestFiveOfAKind(t *testing.T) {
	assert := assert.New(t)

	camels := []*Camel{
		newCamel("J4444"),
		newCamel("44444"),
		newCamel("JJJJ4"),
		newCamel("JJJ44"),
	}

	assert.Equal(FIVE_OF_A_KIND, camels[0].value, "they should be equal")
	assert.Equal(FIVE_OF_A_KIND, camels[1].value, "they should be equal")
	assert.Equal(FIVE_OF_A_KIND, camels[2].value, "they should be equal")
	assert.Equal(FIVE_OF_A_KIND, camels[3].value, "they should be equal")
}

func TestFourOfAKind(t *testing.T) {
	assert := assert.New(t)

	camels := []*Camel{
		newCamel("4444A"),
		newCamel("22JJA"),
		newCamel("AAAJ3"),
	}

	doSort(camels)

	assert.Equal("AAAJ3", camels[0].hand, "they should be equal")
	assert.Equal("4444A", camels[1].hand, "they should be equal")
	assert.Equal("22JJA", camels[2].hand, "they should be equal")
}

func TestFullHouse(t *testing.T) {
	assert := assert.New(t)

	camels := []*Camel{
		newCamel("AAA55"),
		newCamel("AAJ55"),
		newCamel("555JJ"),
	}

	doSort(camels)

	assert.Equal(FIVE_OF_A_KIND, camels[0].value, "they should be equal")
	assert.Equal(FULL_HOUSE, camels[1].value, "they should be equal")
	assert.Equal(FULL_HOUSE, camels[2].value, "they should be equal")
}

func TestThree(t *testing.T) {
	assert := assert.New(t)

	camels := []*Camel{
		newCamel("AJJ12"),
		newCamel("AAJ56"),
		newCamel("JJJ56"),
	}

	// doSort(camels)

	assert.Equal(THREE_OF_A_KIND, camels[0].value, "they should be equal")
	assert.Equal(THREE_OF_A_KIND, camels[1].value, "they should be equal")
	assert.Equal(FOUR_OF_A_KIND, camels[2].value, "they should be equal")
}

func TestOnePair(t *testing.T) {
	assert := assert.New(t)

	camels := []*Camel{
		newCamel("AJ456"),
		newCamel("J4567"),
		newCamel("J1234"),
	}

	doSort(camels)

	assert.Equal(ONE_PAIR, camels[0].value, "they should be equal")
	assert.Equal(ONE_PAIR, camels[1].value, "they should be equal")
	assert.Equal(ONE_PAIR, camels[2].value, "they should be equal")
}
