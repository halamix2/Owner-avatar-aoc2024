package common

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

type Direction rune

func (d Direction) String() string {
	return string(d)
}

const (
	Up    Direction = '^'
	Right Direction = '>'
	Down  Direction = 'v'
	Left  Direction = '<'
)

type Tile struct {
	data    rune
	visited int
}

// Visit and return if we got loop
func (t *Tile) Visit() bool {
	t.data = 't'
	t.visited++
	return t.visited > 4
}

type Floor struct {
	data [][]*Tile
}

func (f *Floor) getGuardStatus() (int, int, Direction) {
	for y, row := range f.data {
		for x, tile := range row {
			direction := Direction(tile.data)
			if direction == Up || direction == Right || direction == Left || direction == Down {
				return x, y, direction
			}
		}
	}
	return -1, -1, Up
}

func (f *Floor) Print() {
	for _, row := range f.data {
		for _, tile := range row {
			fmt.Printf("%c", tile.data)
		}
		fmt.Println()
	}
}

func (f *Floor) getCopy() Floor {
	copiedFloorData := make([][]*Tile, 0)
	for _, column := range f.data {
		copedColumn := make([]*Tile, 0)
		for _, tile := range column {
			newTile := &Tile{data: tile.data, visited: tile.visited}
			copedColumn = append(copedColumn, newTile)
		}
		copiedFloorData = append(copiedFloorData, copedColumn)
	}

	return Floor{data: copiedFloorData}
}

func (f Floor) canPutObstacle(x, y int) bool {
	if f.data[y][x].data != '.' {
		// can't alter the original setup
		return false
	} else if y+1 >= len(f.data) {
		return true
	} else if f.data[y+1][x].data == '^' {
		return true
	}
	return true
}

func (f *Floor) MultiverseLoopPosibilities() int {
	var possibilites atomic.Int64
	var wg sync.WaitGroup
	possibilites.Store(0)
	size := len(f.data)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if !f.canPutObstacle(x, y) {
				// we can't put stuff in front of the guard
				continue
			}
			wg.Add(1)
			go func() {
				copiedFloor := f.getCopy()
				copiedFloor.data[y][x].data = '#'

				// fmt.Printf("checking multiverse %d,%d...\n", x, y)

				loop := copiedFloor.TraverseFloor()
				if loop {
					// fmt.Printf("multiverse %d %d confirmed\n", x, y)
					possibilites.Add(1)
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()

	return int(possibilites.Load())
}

func (f *Floor) TraverseFloor() bool {
	x, y, direction := f.getGuardStatus()
	size := len(f.data)
	for {
		switch direction {
		case Up:
			if y-1 < 0 {
				f.data[y][x].data = 't'
				return false
			} else if f.data[y-1][x].data == '#' {
				direction = Right
			} else {
				if f.data[y][x].Visit() {
					return true
				}
				f.data[y-1][x].data = rune(Up)
				y--
			}
		case Right:
			if x+1 >= size {
				f.data[y][x].data = 't'
				return false
			} else if f.data[y][x+1].data == '#' {
				direction = Down
			} else {
				if f.data[y][x].Visit() {
					return true
				}
				f.data[y][x+1].data = rune(Right)
				x++
			}
		case Down:
			if y+1 >= size {
				f.data[y][x].data = 't'
				return false
			} else if f.data[y+1][x].data == '#' {
				direction = Left
			} else {
				if f.data[y][x].Visit() {
					return true
				}
				f.data[y+1][x].data = rune(Down)
				y++
			}
		case Left:
			if x-1 < 0 {
				f.data[y][x].data = 't'
				return false
			} else if f.data[y][x-1].data == '#' {
				direction = Up
			} else {
				if f.data[y][x].Visit() {
					return true
				}
				f.data[y][x-1].data = rune(Left)
				x--
			}
		}
	}
}

func (f *Floor) GetTraversedTilesCount() int {
	return f.CountChars('t')
}

func (f *Floor) CountChars(char rune) int {
	count := 0
	for _, row := range f.data {
		for _, tile := range row {
			if tile.data == char {
				count++
			}
		}
	}

	return count
}

func ParseInput(filename string) (Floor, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Floor{}, fmt.Errorf("failed to load input: %s\n", err)
	}

	defer f.Close()

	floorData := make([][]*Tile, 0)
	// for i := range floor {
	// 	floor[i] = make([]rune, 0)
	// }

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		column := make([]*Tile, 0)
		for _, tile := range line {
			column = append(column, &Tile{data: tile, visited: 0})
		}
		floorData = append(floorData, column)
	}

	if err := scanner.Err(); err != nil {
		return Floor{}, fmt.Errorf("failed to scan input: %s", err)
	}

	return Floor{data: floorData}, nil
}
