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

	totalSimilarityScore := 0
	for _, a := range listA {
		multiplier := getNumberOfElements(listB, a)
		similarityScore := a * multiplier
		totalSimilarityScore += similarityScore
	}
	fmt.Printf("total similarity score: %d\n", totalSimilarityScore)
}

func getNumberOfElements(list []int, searched int) int {
	count := 0
	for _, el := range list {
		if el == searched {
			count++
		}
	}
	return count
}
