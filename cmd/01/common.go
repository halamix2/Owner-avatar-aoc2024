package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseFileToSortedLists(filePath string) ([]int, []int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return []int{}, []int{}, fmt.Errorf("failed to load input: %s\n", err)
	}

	listA := make([]int, 0)
	listB := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) != 2 {
			return []int{}, []int{}, errors.New("failed to split input")
		}
		a, err := strconv.Atoi(numbers[0])
		if err != nil {
			return []int{}, []int{}, err
		}
		b, err := strconv.Atoi(numbers[1])
		if err != nil {
			return []int{}, []int{}, err
		}
		listA = append(listA, a)
		listB = append(listB, b)
	}
	if err := scanner.Err(); err != nil {
		return []int{}, []int{}, fmt.Errorf("failed to scan input: %s", err)
	}
	sort.Ints(listA)
	sort.Ints(listB)
	return listA, listB, f.Close()
}
