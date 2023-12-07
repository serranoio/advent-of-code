package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type Camel struct {
	bid   int
	hand  string
	value int
}

func initCamel(v string) *Camel {
	twoParts := strings.Split(v, " ")

	bid, _ := strconv.Atoi(twoParts[1])
	return &Camel{
		bid:  int(bid),
		hand: twoParts[0],
	}
}

func (c *Camel) calculateHandValue() {
	numbers := make(map[string]int)

	// create map
	for _, aRune := range c.hand {
		letter := string(aRune)
		if _, ok := numbers[letter]; ok {
			numbers[letter] += 1
		} else {
			numbers[letter] = 1
		}
	}

	if v, ok := numbers["J"]; ok {

		highestNumber := 0
		highestValue := ""

		for k, v := range numbers {
			if v > highestNumber && k != "J" {
				highestNumber = v
				highestValue = k
			}
		}

		numbers[highestValue] += v
		delete(numbers, "J")
	}

	// convert map to amount
	amountOfNumbers := make(map[int]int)
	for _, number := range numbers {
		if _, ok := amountOfNumbers[number]; ok {
			amountOfNumbers[number] += 1
		} else {
			amountOfNumbers[number] = 1
		}
	}

	if amountOfNumbers[5] == 1 {
		c.value = FIVE_OF_A_KIND
	} else if amountOfNumbers[4] == 1 {
		c.value = FOUR_OF_A_KIND
	} else if amountOfNumbers[3] == 1 && amountOfNumbers[2] == 1 {
		c.value = FULL_HOUSE
	} else if amountOfNumbers[3] == 1 {
		c.value = THREE_OF_A_KIND
	} else if amountOfNumbers[2] == 2 {
		c.value = TWO_PAIR
	} else if amountOfNumbers[2] == 1 ||
		numbers["J"] == 1 {
		c.value = ONE_PAIR
	} else {
		c.value = HIGH_CARD
	}
}

const (
	J = 1
	T = iota + 10
	Q
	K
	A
)

// const (
// 	T = iota + 10
// 	J
// 	Q
// 	K
// 	A
// )

var stringsToValue = map[rune]int{
	'T': T,
	'J': J,
	'Q': Q,
	'K': K,
	'A': A,
}

const (
	GREATER = iota
	EQUAL
	LESS
)

func insideMap(letter rune, secondLetter rune) int {
	letter1, ok1 := stringsToValue[letter]
	letter2, ok2 := stringsToValue[secondLetter]

	if !ok1 {
		letter1, _ = strconv.Atoi(string(letter))
	}
	if !ok2 {
		letter2, _ = strconv.Atoi(string(secondLetter))
	}

	if letter1 > letter2 {
		return GREATER
	} else if letter1 == letter2 {
		return EQUAL
	} else {
		return LESS
	}
}

func (c *Camel) compare(second *Camel) bool {
	if c.value > second.value {
		return true
	} else if c.value == second.value {
		for letter := 0; letter < 5; letter++ {
			value := insideMap(rune(c.hand[letter]), rune(second.hand[letter]))
			if value == GREATER {
				return true
			} else if value == LESS {
				return false
			}
		}
	}
	return false
}

func doSort(allCamels []*Camel) {
	sort.Slice(allCamels, func(i, j int) bool {
		return allCamels[i].compare(allCamels[j])
	})
}

func calculateWinnings(allCamels []*Camel) {

	length := len(allCamels)
	winnings := 0
	for num := 0; num < length; num++ {
		winnings += allCamels[num].bid * (length - num)
	}

	fmt.Println(winnings)
}

func main() {
	start := time.Now()

	input, _ := os.ReadFile("adventday6.txt")
	lines := strings.Split(string(input), "\n")

	allCamels := []*Camel{}

	for _, line := range lines {
		camel := initCamel(string(line))
		camel.calculateHandValue()
		allCamels = append(allCamels, camel)
	}

	doSort(allCamels)
	calculateWinnings(allCamels)
	// at beginning put first

	elapsed := time.Since(start)

	log.Printf("Took %s", elapsed)
}
