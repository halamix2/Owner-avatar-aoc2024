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
	possibilites := floor.MultiverseLoopPosibilities()

	fmt.Printf("Counted %d possibilites\n", possibilites)
}
