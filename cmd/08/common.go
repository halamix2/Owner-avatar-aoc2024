package common

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"

	"klaidliadon.dev/next"
)

/*
freq - lowercase, uppercase, letter
bad signal - on antinodes
    perfectly in line with two antennas of the same frequency
    one antenna twice as fas away as the other - 1/3 and 2/3 OUTWARD of the way I guess
    basically # A A in 1/3 discances, or the antinode distance from any antenna is equat to the distance between antennas
    antinodes CAN occur on top of another antenna

*/

type Point struct {
	x int
	y int
}

func getAntinodes(a, b Point, all bool, width, height int) []Point {
	antinodes := make([]Point, 0)
	distX := a.x - b.x
	distY := a.y - b.y

	if all {
		i := 0
		for {
			x := a.x + (distX * i)
			y := a.y + (distY * i)
			if x < 0 || x >= width || y < 0 || y >= height {
				break
			}
			antinodes = append(antinodes, Point{x: x, y: y})
			i++
		}

		i = 0
		for {
			x := b.x - (distX * i)
			y := b.y - (distY * i)
			if x < 0 || x >= width || y < 0 || y >= height {
				break
			}
			antinodes = append(antinodes, Point{x: x, y: y})
			i++
		}
	} else {
		antinodes = append(antinodes, Point{x: a.x + distX, y: a.y + distY})
		antinodes = append(antinodes, Point{x: b.x - distX, y: b.y - distY})
	}

	return antinodes
}

type Tile struct {
	data     rune
	antinode bool
}

type AntennaMap struct {
	data [][]*Tile
}

func (a *AntennaMap) Print() {
	for _, row := range a.data {
		for _, tile := range row {
			if tile.antinode {
				fmt.Printf("#")
			} else {
				fmt.Printf("%c", tile.data)
			}
		}
		fmt.Println()
	}
}

func (a *AntennaMap) PointInMap(p Point) bool {
	height := len(a.data)
	width := len(a.data[0])
	return p.x >= 0 && p.x < width && p.y >= 0 && p.y < height
}

func (a *AntennaMap) MarkAntinode(p Point) {
	a.data[p.y][p.x].antinode = true
}

func (a *AntennaMap) getAllFrequencies() []rune {
	frequencies := make([]rune, 0)
	for _, row := range a.data {
		for _, tile := range row {
			if !slices.Contains(frequencies, tile.data) && tile.data != '.' {
				frequencies = append(frequencies, tile.data)
			}
		}
	}

	return frequencies
}

func (a *AntennaMap) getAntennasByFrequency(freq rune) []interface{} {
	antennas := make([]interface{}, 0)
	for y, row := range a.data {
		for x, tile := range row {
			if tile.data == freq {
				antennas = append(antennas, Point{x: x, y: y})
			}
		}
	}
	return antennas
}

func (a *AntennaMap) CalculateAntinodes(all bool) {
	height := len(a.data)
	width := len(a.data[0])

	frequencies := a.getAllFrequencies()
	fmt.Printf("got %d frequencies...\n", len(frequencies))
	for _, freq := range frequencies {
		fmt.Printf("checking %c frequency\n", freq)
		antennas := a.getAntennasByFrequency(freq)

		// combinations over antennae
		for antennaPair := range next.Combination(antennas, 2, false) {
			//convert to something more usable
			ap := make([]Point, 2)
			for i := range 2 {
				ap[i] = reflect.ValueOf(antennaPair).Index(i).Interface().(Point)
			}
			fmt.Printf("checking antinodes of %v\n", ap)
			// calculate distance
			antinodes := getAntinodes(ap[0], ap[1], all, width, height)
			for _, antinode := range antinodes {
				if a.PointInMap(antinode) {
					a.MarkAntinode(antinode)
				}
			}
		}
		//kombinatoryka?
	}
}

func (a *AntennaMap) CountAntinodes() int {
	count := 0
	for _, row := range a.data {
		for _, tile := range row {
			if tile.antinode {
				count++
			}
		}
	}
	return count
}

func ParseInput(filename string) (AntennaMap, error) {
	f, err := os.Open(filename)
	if err != nil {
		return AntennaMap{}, fmt.Errorf("failed to load input: %s\n", err)
	}

	antennaData := make([][]*Tile, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		column := make([]*Tile, 0)
		for _, tile := range line {
			column = append(column, &Tile{data: tile, antinode: false})
		}
		antennaData = append(antennaData, column)
	}
	if err := scanner.Err(); err != nil {
		return AntennaMap{}, fmt.Errorf("failed to scan input: %s", err)
	}
	return AntennaMap{data: antennaData}, nil
}
