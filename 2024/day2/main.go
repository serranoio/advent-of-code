package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	clog "github.com/charmbracelet/log"
)

func createReports(input string) [][]int {
	var reports [][]int

	for _, line := range strings.Split(input, "\n") {
		report := []int{}
		for _, char := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(char)
			report = append(report, num)
		}

		reports = append(reports, report)
	}

	return reports
}

func readFile(name string) string {
	bytes, _ := os.ReadFile(name)

	return string(bytes)
}

func checkSingleReport(report []int) int {
	isAscending := true
	// 1, 3
	if report[0]-report[1] < 0 {
		isAscending = false
	}

	for i := 0; i < len(report)-1; i++ {
		if !isAscending && !(report[i+1]-report[i] >= 1 && report[i+1]-report[i] <= 3) {
			return 0
		}
		if isAscending && !(report[i]-report[i+1] >= 1 && report[i]-report[i+1] <= 3) {
			return 0
		}
	}

	return 1
}

const (
	UP   direction = 1
	DOWN direction = 2
	NONE direction = 3
)

type direction int

// now, we have to determine direction as we go.
func checkSingleReportWithProblemDampener(report []int) int {
	dampen := 0

	isAscending := NONE

	// originalReport := report

	for i := 0; i < len(report)-1; i++ {
		// ascending
		// 1,2
		if report[i+1]-report[i] >= 1 && report[i+1]-report[i] <= 3 {

			if isAscending == DOWN {
				if dampen < 1 {
					if report[i-1] == report[i+1] {
						isAscending = UP
					} else {
						report = append(report[:i], report[i+1:]...)
						i--
						dampen++
					}

					continue
				}
				return 0
			}
			isAscending = UP
		} else if report[i]-report[i+1] >= 1 && report[i]-report[i+1] <= 3 {

			if isAscending == UP {
				if dampen < 1 {
					if report[i-1] == report[i+1] {
						isAscending = DOWN
					} else {
						report = append(report[:i], report[i+1:]...)
						i--
						dampen++
					}
					continue
				}
				return 0
			}
			isAscending = DOWN
		} else if dampen >= 1 {
			return 0
		} else {
			if i == 0 {
				report = append(report[:i], report[i+1:]...)
			} else {
				report = append(report[:i+1], report[i+2:]...)
			}

			i--
			dampen++
		}

	}

	return 1
}

func calculateReportSafety(reports [][]int, part1Bool bool) int {

	safeReportsCount := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, report := range reports {
		wg.Add(1)
		go func() {
			var isSafe int
			if part1Bool {
				isSafe = checkSingleReport(report)
			} else {
				isSafe = checkSingleReportWithProblemDampener(report)

				if isSafe == 0 {
					fmt.Printf("report %d isSafe=%b\n", report, isSafe)
				}
			}

			mu.Lock()
			safeReportsCount += isSafe
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	return safeReportsCount
}

func part1(name string) int {
	input := readFile(name)

	reports := createReports(input)

	return calculateReportSafety(reports, true)
}

func part2(name string) int {
	input := readFile(name)

	reports := createReports(input)

	return calculateReportSafety(reports, false)
}

func main() {
	clog.Info("Part 1", "Answer", part1("input.txt"))
	clog.Info("Part 2", "Answer", part2("input.txt"))
}
