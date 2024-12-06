package main

import (
	"fmt"
	"os"

	common "halamix2.pl/advent_of_code_24/cmd/06"
)

func main() {
	floor, err := common.ParseInput("cmd/06/input.txt")
	if err != nil {
		fmt.Printf("failed to load data:%s\n", err)
		os.Exit(1)
	}

	loop := floor.TraverseFloor()
	if loop {
		fmt.Printf("cound not escape maze: %s", err)
	}

	traversedCount := floor.GetTraversedTilesCount()
	columnCount := floor.CountChars('#')

	floor.Print()
	fmt.Printf("Counted %d columns, and %d traversed tiles\n", columnCount, traversedCount)
}
