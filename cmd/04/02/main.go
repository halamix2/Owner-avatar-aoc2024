package main

import (
	"fmt"
	"os"

	common "halamix2.pl/advent_of_code_24/cmd/04"
)

func main() {
	data, err := common.ParseInput("cmd/04/input.txt")
	if err != nil {
		fmt.Printf("parsing is broken: %d\n", err)
		os.Exit(1)
	}

	xmas := 0
	for x, column := range data {
		for y, singleRune := range column {
			if singleRune == 'A' {
				if x-1 >= 0 && x+1 < common.WIDTH && y-1 >= 0 && y+1 < common.HEIGHT {
					if crossFound(data, x, y) {
						xmas++
					}
				}
			}
		}
	}

	fmt.Printf("x-mases: %d\n", xmas)
}

func crossFound(d [][]rune, x, y int) bool {
	// starting from the centre
	if (d[x-1][y-1] != 'M' || d[x+1][y+1] != 'S') && (d[x-1][y-1] != 'S' || d[x+1][y+1] != 'M') {
		return false
	}
	if (d[x-1][y+1] != 'M' || d[x+1][y-1] != 'S') && (d[x-1][y+1] != 'S' || d[x+1][y-1] != 'M') {
		return false
	}
	return true
}
