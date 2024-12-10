package main

import (
	"fmt"
	"os"

	common "halamix2.pl/advent_of_code_24/cmd/10"
)

func main() {
	heightMap, err := common.ParseInput("cmd/10/in.txt")
	if err != nil {
		fmt.Printf("failed to load data:%s\n", err)
		os.Exit(1)
	}

	sum := heightMap.TrailheadsSum()
	fmt.Printf("Sum of trailheads: %d = 36?\n", sum)
}
