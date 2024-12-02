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

	badReportsCount := len(common.GetBadReports(reports))
	safeReports := len(reports) - badReportsCount

	fmt.Printf("Safe reports: %d\n", safeReports)
}
