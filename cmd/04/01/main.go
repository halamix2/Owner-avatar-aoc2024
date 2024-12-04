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

	horizontals := 0
	diagonalsToRIght := 0
	diagonalsToLeft := 0
	verticals := 0
	for x, column := range data {
		for y, singleRune := range column {
			// skip checks on other chars to speed things up
			if singleRune == 'X' || singleRune == 'S' {
				if x+3 < common.WIDTH && y+3 < common.HEIGHT {
					// diagonals top left-bottom right
					if findDiagonalsToRight(data, x, y) {
						diagonalsToRIght++
					}
				}
				if x-3 >= 0 && y+3 < common.HEIGHT {
					// diagonals top right-bottom left
					if findDiagonalsToLeft(data, x, y) {
						diagonalsToLeft++
					}
				}
				if y+3 < common.HEIGHT {
					if findVertical(data, x, y) {
						verticals++
					}
				}
				if x+3 < common.WIDTH {
					if findHorizontal(data, x, y) {
						horizontals++
					}
				}
			}
		}
	}

	total := horizontals + diagonalsToRIght + diagonalsToLeft + verticals
	fmt.Printf("Found total %d, including %d horizontal, %d vertical, %d /diagonal and %d \\diagonal\n", total, horizontals, verticals, diagonalsToLeft, diagonalsToRIght)
}

func findDiagonalsToRight(data [][]rune, x, y int) bool {
	if data[x][y] == 'X' && data[x+1][y+1] == 'M' && data[x+2][y+2] == 'A' && data[x+3][y+3] == 'S' {
		return true
	} else if data[x][y] == 'S' && data[x+1][y+1] == 'A' && data[x+2][y+2] == 'M' && data[x+3][y+3] == 'X' {
		return true
	}
	return false
}

func findDiagonalsToLeft(data [][]rune, x, y int) bool {
	if data[x][y] == 'X' && data[x-1][y+1] == 'M' && data[x-2][y+2] == 'A' && data[x-3][y+3] == 'S' {
		return true
	} else if data[x][y] == 'S' && data[x-1][y+1] == 'A' && data[x-2][y+2] == 'M' && data[x-3][y+3] == 'X' {
		return true
	}
	return false
}

func findVertical(data [][]rune, x, y int) bool {
	if data[x][y] == 'X' && data[x][y+1] == 'M' && data[x][y+2] == 'A' && data[x][y+3] == 'S' {
		return true
	} else if data[x][y] == 'S' && data[x][y+1] == 'A' && data[x][y+2] == 'M' && data[x][y+3] == 'X' {
		return true
	}
	return false
}

func findHorizontal(data [][]rune, x, y int) bool {
	if data[x][y] == 'X' && data[x+1][y] == 'M' && data[x+2][y] == 'A' && data[x+3][y] == 'S' {
		return true
	} else if data[x][y] == 'S' && data[x+1][y] == 'A' && data[x+2][y] == 'M' && data[x+3][y] == 'X' {
		return true
	}
	return false
}
