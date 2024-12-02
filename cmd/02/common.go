package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(filename string) ([][]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]int{}, fmt.Errorf("failed to load input: %s\n", err)
	}

	defer f.Close()

	reports := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		levels := strings.Fields(scanner.Text())
		parsedLevels := make([]int, 0)
		for _, level := range levels {
			parsedLevel, _ := strconv.Atoi(level)
			parsedLevels = append(parsedLevels, parsedLevel)
		}

		reports = append(reports, parsedLevels)

	}
	if err := scanner.Err(); err != nil {
		return [][]int{}, fmt.Errorf("failed to scan input: %s", err)
	}

	return reports, nil
}

func SafeDifference(a, b int) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return (diff > 0) && (diff < 4)
}

func GetBadReports(reports [][]int) [][]int {
	badReports := make([][]int, 0)
	for _, levels := range reports {
		if !ParseReport(levels) {
			badReports = append(badReports, levels)
		}
	}
	return badReports
}

func ParseReport(levels []int) bool {
	levelCount := len(levels)
	increasing := false
	goodReport := true
	for i, level := range levels {
		if i+1 < levelCount {
			nextLevel := levels[i+1]
			currentIncreasing := false

			// skip ones with unchanging levels
			if level == nextLevel {
				goodReport = false
				break
			}

			// chek if we have increasing numbers
			if level > nextLevel {
				currentIncreasing = true
				if i == 0 {
					// set inital increasing state
					increasing = true
				}
			}

			// skip if we've changed direction
			if increasing != currentIncreasing {
				goodReport = false
				break
			}

			// skip if data difference is too big
			if !SafeDifference(level, nextLevel) {
				goodReport = false
				break
			}
		}
	}
	return goodReport
}
