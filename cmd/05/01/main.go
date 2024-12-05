package main

import (
	"fmt"
	"os"

	common "halamix2.pl/advent_of_code_24/cmd/05"
)

func main() {
	rules, manuals, err := common.ParseInput("cmd/05/input.txt")
	if err != nil {
		fmt.Printf("failed to load data:%s\n", err)
		os.Exit(1)
	}

	sum := 0

	for _, manual := range manuals {
		if manual.CompareAgaintsRuleset(rules) {
			sum += manual.GetMiddle()
		}
	}

	fmt.Printf("Sum: %d\n", sum)

}
