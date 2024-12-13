package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFileIntoStones(fileName string) map[int]int {
	bytes, _ := os.ReadFile(fileName)

	stones := make(map[int]int)

	for _, num := range strings.Split(string(bytes), " ") {
		n, _ := strconv.Atoi(num)

		if _, ok := stones[n]; ok {
			stones[n] += 1
		} else {
			stones[n] = 1
		}
	}

	return stones
}

func blinkBitch(stones map[int]int) map[int]int {
	newMap := make(map[int]int)

	for stone, count := range stones {
		if stone == 0 {

			newMap[1] += count

			continue
		}
		lenDigits := len(fmt.Sprint(stone))

		if lenDigits%2 == 0 {
			firstStone, _ := strconv.Atoi(fmt.Sprint(stone)[0 : lenDigits/2])
			secondStone, _ := strconv.Atoi(fmt.Sprint(stone)[lenDigits/2:])

			newMap[firstStone] += count
			newMap[secondStone] += count

			continue
		}

		newMap[stone*2024] += count

	}

	// for i := 0; i < len(stones); i++ {
	// 	if stones[i] == 0 {
	// 		stones[i] = 1
	// 		newStones = append(newStones, stones[i])
	// 		continue
	// 	}

	// 	lenDigits := len(fmt.Sprint(stones[i]))
	// 	if lenDigits%2 == 0 {

	// 		firstStone, _ := strconv.Atoi(fmt.Sprint(stones[i])[0 : lenDigits/2])
	// 		secondStone, _ := strconv.Atoi(fmt.Sprint(stones[i])[lenDigits/2:])

	// 		newStones = append(newStones, firstStone)
	// 		newStones = append(newStones, secondStone)

	// 		continue
	// 	}

	// 	stones[i] *= 2024

	// 	newStones = append(newStones, stones[i])
	// }

	return newMap
}

func blinkTimes(stones map[int]int, times int) int {

	for i := 0; i < times; i++ {
		stones = blinkBitch(stones)
	}

	count := 0
	for _, v := range stones {
		count += v
	}

	return count
}

func main() {
	stones := readFileIntoStones("input.txt")
	stoneCount := blinkTimes(stones, 75)

	fmt.Println(stoneCount)
}
