package main

import (
	"fmt"
	"os"

	common "halamix2.pl/advent_of_code_24/cmd/07"
)

func main() {
	equations, err := common.ParseInput("cmd/07/input.txt")
	if err != nil {
		fmt.Printf("failed to load data:%s\n", err)
		os.Exit(1)
	}

	sum := 0
	for _, equation := range equations {
		if equation.IsCorrect([]interface{}{'+', '*'}) {
			sum += equation.GetResult()
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
