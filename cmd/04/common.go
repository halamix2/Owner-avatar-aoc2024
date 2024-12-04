package common

import (
	"bufio"
	"fmt"
	"os"
)

const (
	WIDTH  = 140
	HEIGHT = 140
)

func ParseInput(filename string) ([][]rune, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]rune{}, fmt.Errorf("failed to load input: %s\n", err)
	}

	defer f.Close()
	data := make([][]rune, WIDTH)
	for i := 0; i < WIDTH; i++ {
		data[i] = make([]rune, HEIGHT)
	}

	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			data[x][y] = char
		}
		y++
	}

	return data, nil
}
