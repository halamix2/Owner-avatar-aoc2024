package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("cmd/03/input.txt")
	if err != nil {
		fmt.Printf("failed to load input: %s\n", err)
		os.Exit(1)
	}
	mulRegexp := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	numberRegexp := regexp.MustCompile(`\d+`)
	sum := 0
	hits := mulRegexp.FindAll(data, -1)
	enable := true
	for _, hit := range hits {
		fmt.Printf("%s\n", string(hit))
		if string(hit) == "do()" {
			fmt.Println("yep")
			enable = true
		} else if string(hit) == "don't()" {
			enable = false
		} else if enable {
			numbers := numberRegexp.FindAll(hit, -1)
			if len(numbers) != 2 {
				fmt.Println("numebrs len not equal 2!")
				os.Exit(1)
			}
			multiply := 1
			for _, n := range numbers {
				number, _ := strconv.Atoi(string(n))
				multiply *= number
			}
			sum += multiply
		}
	}
	fmt.Printf("sum: %d\n", sum)
}
