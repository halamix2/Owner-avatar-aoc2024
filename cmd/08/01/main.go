package main

import (
	"fmt"
	"os"

	common "halamix2.pl/advent_of_code_24/cmd/08"
)

func main() {
	antennaMap, err := common.ParseInput("cmd/08/input.txt")
	if err != nil {
		fmt.Printf("failed to load data:%s\n", err)
		os.Exit(1)
	}

	antennaMap.CalculateAntinodes(false)

	antennaMap.Print()
	fmt.Println()

	count := antennaMap.CountAntinodes()

	fmt.Printf("total antinodes: %d\n", count)
}
