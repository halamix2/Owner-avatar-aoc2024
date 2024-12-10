package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
0-9 lowest-highest
get longest low-to-hight string
even, gradual
0-9, step 1 (no diagonals)

trailhead - 0
trailhead score - number of reachable nines

*/

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("{%d, %d}", p.x, p.y)
}

type Map struct {
	data [][]int
}

func (m *Map) getNeighbors(p Point, val int) []Point {
	points := make([]Point, 0)
	if p.y+1 < len(m.data) && m.data[p.y+1][p.x] == val {
		points = append(points, Point{x: p.x, y: p.y + 1})
	}
	if p.x+1 < len(m.data[0]) && m.data[p.y][p.x+1] == val {
		points = append(points, Point{x: p.x + 1, y: p.y})
	}
	if p.y-1 >= 0 && m.data[p.y-1][p.x] == val {
		points = append(points, Point{x: p.x, y: p.y - 1})
	}
	if p.x-1 >= 0 && m.data[p.y][p.x-1] == val {
		points = append(points, Point{x: p.x - 1, y: p.y})
	}

	return points
}

func (m *Map) recurTrailblazer() int {
	return -1
}

func (m *Map) getScore(p Point) int {
	score := 0
	// currentPoint := p
	// for i := 1; i < 10; i++ {
	neighbors := m.getNeighbors(p, 1)

	// for each neighbor check its neighbors, up to val 9
	for _, n := range neighbors {
		currNeigh := m.getNeighbors(n, 2)
		_ = currNeigh
	}
	// count all nines

	// }

	return score
}

func (m *Map) getAllPoints(wanted int) []Point {
	points := make([]Point, 0)
	for y, column := range m.data {
		for x, data := range column {
			if data == wanted {
				points = append(points, Point{x: x, y: y})
			}
		}
	}

	return points
}

func (m *Map) TrailheadsSum() int {
	sum := 0

	zeroes := m.getAllPoints(0)
	for _, zero := range zeroes {
		fmt.Printf("Looking for %s = ", zero)
		score := m.getScore(zero)
		fmt.Printf("%d\n", score)
		sum += score
	}

	return sum
}

func ParseInput(filename string) (Map, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Map{}, fmt.Errorf("failed to load input: %s\n", err)
	}

	heightMap := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		column := make([]int, 0)
		for _, char := range line {
			height, _ := strconv.Atoi(string(char))
			column = append(column, height)
		}
		heightMap = append(heightMap, column)
	}
	if err := scanner.Err(); err != nil {
		return Map{}, fmt.Errorf("failed to scan input: %s", err)
	}
	return Map{data: heightMap}, nil
}
