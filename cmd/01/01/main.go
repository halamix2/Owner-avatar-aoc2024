package main

import (
	"fmt"
	"os"

	common "halamix2.pl/advent_of_code_24/cmd/01"
)

func main() {
	listA, listB, err := common.ParseFileToSortedLists("cmd/01/input.txt")
	if err != nil {
		fmt.Printf("couldn't parse input file: %s\n", err)
		os.Exit(1)
	}

	totalDiff := 0
	for i, a := range listA {
		b := listB[i]
		diff := a - b
		if diff < 0 {
			diff = -diff
		}
		totalDiff += diff
	}
	fmt.Printf("total difference: %d\n", totalDiff)
}
