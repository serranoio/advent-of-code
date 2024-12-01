package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type CardHand struct {
	card       []int
	hand       []int
	matchCount int
	cardNumber int
	copies     int
}

type AllHands struct {
	allHands []CardHand
	total    int
}

func initCardHand() *CardHand {
	return &CardHand{
		card:       []int{},
		hand:       []int{},
		matchCount: 0,
		cardNumber: 0,
		copies:     1,
	}
}

func convertStringToIntArray(content string) []int {

	allInts := []int{}

	for _, num := range strings.Split(content, " ") {
		if num == "" {
			// split could be at a double space junction,
			// causing extra num. hence, continue
			continue
		}

		newInt, _ := strconv.Atoi(num)
		allInts = append(allInts, newInt)
	}
	return allInts
}

func (c *CardHand) createMatchCount() {

	for _, numCard := range c.card {
		for _, numHand := range c.hand {
			// from part 1:
			// if numHand == numCard && c.matchCount == 0 {
			// 	c.matchCount = 1
			// } else if numHand == numCard {
			// 	c.matchCount *= 2
			// }

			if numHand == numCard {
				c.matchCount += 1
			}
		}
	}
}

func allContentsSplit(nameOfFile string) []string {
	contents, _ := os.ReadFile(nameOfFile)

	return strings.Split(string(contents), "\n")
}

func createHands(content []string) []CardHand {
	allCardHands := []CardHand{}

	for _, line := range content {
		cardHand := initCardHand()

		cardAndHand := strings.Split(line, "|")
		cardPart := strings.Split(cardAndHand[0], ":")

		// Card 5: -> ["Card", "5"] -> "5" => 5
		cardHand.cardNumber, _ = strconv.Atoi(strings.Split(cardPart[0], " ")[1])
		cardNumbers := cardPart[1]
		handNumbers := cardAndHand[1]

		cardHand.card = convertStringToIntArray(cardNumbers)
		cardHand.hand = convertStringToIntArray(handNumbers)
		cardHand.createMatchCount()

		allCardHands = append(allCardHands, *cardHand)
	}

	return allCardHands
}

// from challenge one
func (aH *AllHands) calcTotalPart1() {
	for _, cardHand := range aH.allHands {
		aH.total += cardHand.matchCount
	}
}

func initACH(cardHands []CardHand) *AllHands {
	return &AllHands{
		total:    0,
		allHands: cardHands,
	}
}

func copyCopy(aH *AllHands, cardNumber int) int {
	total := 1
	if cardNumber >= len(aH.allHands) {
		return total
	}

	for num := 0; num < aH.allHands[cardNumber].matchCount; num++ {
		total += copyCopy(aH, num+1+cardNumber)
	}

	return total
}

func (aH *AllHands) calcTotalPart2() {
	for num := 0; num < len(aH.allHands); num++ {
		aH.total += copyCopy(aH, num)
	}
}

func main() {
	allCardHands := createHands(allContentsSplit("adventday4.txt"))

	ah := initACH(allCardHands)
	ah.calcTotalPart2()

	log.Info(ah.total)
}
