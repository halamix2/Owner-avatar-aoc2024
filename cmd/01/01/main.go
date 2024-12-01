package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("cmd/01/01/input.txt")
	if err != nil {
		fmt.Printf("failed to load input: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	listA := make([]int, 0)
	listB := make([]int, 0)

	i := 0
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) != 2 {
			fmt.Printf("failed to split input")
			os.Exit(1)
		}
		a, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic("owo")
		}
		b, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic("uwu")
		}
		listA = append(listA, a)
		listB = append(listB, b)
		i++
	}

	sort.Ints(listA)
	sort.Ints(listB)

	totalDiff := 0
	for i, a := range listA {
		b := listB[i]
		diff := a - b
		if diff < 0 {
			diff = -diff
		}
		totalDiff += diff
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("failed to scan input: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("total difference: %d\n", totalDiff)
}
