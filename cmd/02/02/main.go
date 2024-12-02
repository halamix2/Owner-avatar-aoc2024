package main

import (
	"fmt"
	"os"

	common "halamix2.pl/advent_of_code_24/cmd/02"
)

/*
Input data:
each line is a report
report consists of levels (ints), sperated by spaces

safe reports - data is only increasing or decreasing
adjacent levels must differ by <1, 3>

Output:
How many reports are safe
*/

func main() {
	reports, err := common.ParseInput("cmd/02/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %s", err)
		os.Exit(1)
	}

	badReports := common.GetBadReports(reports)
	safeReports := len(reports) - len(badReports)

	// try to fix bad reports by looping each report with each level removed, until it becomes good

	fixedReports := 0
	for _, levels := range badReports {
		isFixed := false
		for skipped := range levels {
			levelsShort := getLevelsSkipped(levels, skipped)
			if common.ParseReport(levelsShort) {
				isFixed = true
				break
			}
		}
		if isFixed {
			fixedReports++
		}
	}

	fmt.Printf("Safe reports: %d, corrected reports: %d, total dampened: %d\n", safeReports, fixedReports, safeReports+fixedReports)
}

func getLevelsSkipped(levels []int, skipped int) []int {
	levelsSkipped := make([]int, 0)
	for i, level := range levels {
		if i != skipped {
			levelsSkipped = append(levelsSkipped, level)
		}
	}
	return levelsSkipped
}
