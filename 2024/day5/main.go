package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitFile(name string) (map[int][]int, [][]int) {

	pageOrderingRules := make(map[int][]int)

	updates := [][]int{}

	for _, line := range strings.Split(name, "\n") {
		if strings.Contains(line, "|") {
			nums := strings.Split(line, "|")

			num0, _ := strconv.Atoi(nums[0])
			num1, _ := strconv.Atoi(nums[1])

			if _, ok := pageOrderingRules[num0]; !ok {
				pageOrderingRules[num0] = []int{num1}
			} else {
				pageOrderingRules[num0] = append(pageOrderingRules[num0], num1)
			}
		} else {
			nums := strings.Split(line, ",")

			update := []int{}
			for _, num := range nums {
				numInt, _ := strconv.Atoi(num)

				update = append(update, numInt)
			}
			updates = append(updates, update)
		}
	}

	updates = updates[1:]

	return pageOrderingRules, updates
}

func readFile(fileName string) string {
	bytes, _ := os.ReadFile(fileName)

	return string(bytes)
}

func appearInList(list []int, num int, update []int) int {
	if len(update) == 0 {
		return 1
	}

	for _, item := range update {
		found := false
		for _, listNum := range list {
			if item == listNum {
				found = true
			}
		}

		if found {
			return 1
		} else {
			return 0
		}
	}

	return 0
}

func calculateUpdate(pageOrderingRules map[int][]int, update []int) int {
	for i, num := range update {
		// 75
		// [list]
		list := pageOrderingRules[num]
		if appearInList(list, num, update[i+1:]) == 0 {
			return 0
		}
		// entire list on the right has to exist on the left

	}

	return 1
}

func fixList(pageOrderingRules map[int][]int, update []int) int {

	for i := 0; i < len(update); i++ {
		num := update[i]
		rules := pageOrderingRules[num]
		behindList := update[i+1:]

		begI := i
		if appearInList2(rules, num, behindList) == 0 {
			update[begI], update[begI+1] = update[begI+1], update[begI]
			i = -1
		}
	}

	return update[len(update)/2]
}

func appearInList2(rules []int, num int, update []int) int {
	if len(update) == 0 {
		return 1
	}

	for _, item := range update {
		found := false
		for _, rule := range rules {
			if item == rule {
				found = true
			}
		}

		if found {
			return 1
		} else {
			return 0
		}
	}

	return 0
}

func calculateUpdates(pageOrderingRules map[int][]int, updates [][]int) (int, int) {
	firstPart := 0
	secondPart := 0
	for _, update := range updates {
		if calculateUpdate(pageOrderingRules, update) == 1 {
			firstPart += update[len(update)/2]
		} else {
			secondPart += fixList(pageOrderingRules, update)
		}
	}

	return firstPart, secondPart
}

func bothParts(name string) (int, int) {
	contents := readFile(name)

	pageOrderingRules, updates := splitFile(contents)

	return calculateUpdates(pageOrderingRules, updates)
}

func main() {
	part1, part2 := bothParts("input.txt")

	fmt.Println(part1)
	fmt.Println(part2)
}
