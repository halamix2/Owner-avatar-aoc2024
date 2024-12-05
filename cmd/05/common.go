package common

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type RuleSet []Rule

type Rule struct {
	pageNumber int
	moreThan   int
}

func (rules RuleSet) getSortingFunc() func(a int, b int) int {
	return func(a int, b int) int {
		both := []int{a, b}
		for _, rule := range rules {
			if !slices.Contains(both, rule.pageNumber) || !slices.Contains(both, rule.moreThan) {
				continue
			}
			if rule.IsCorrect(both) {
				continue
			} else if a == rule.pageNumber && b == rule.moreThan {
				return 1
			} else if b == rule.pageNumber && a == rule.moreThan {
				return -1
			}
		}
		return 0
	}
}

func (r Rule) IsCorrect(data []int) bool {
	pageNumberPosition := slices.Index(data, r.pageNumber)
	if pageNumberPosition < 0 {
		// nothing to do, move along
		return true
	}
	for i := 0; i <= pageNumberPosition; i++ {
		if data[i] == r.moreThan {
			return false
		}
	}

	return true
}

type Manual struct {
	data []int
}

func (m Manual) CompareAgaintsRuleset(rules []Rule) bool {
	for _, rule := range rules {
		if !rule.IsCorrect(m.data) {
			return false
		}
	}
	return true
}

func (m Manual) GetMiddle() int {
	return m.data[len(m.data)/2]
}

func (m Manual) Sort(rules RuleSet) Manual {
	sortedData := make([]int, len(m.data))
	_ = copy(sortedData, m.data)

	slices.SortFunc(sortedData, rules.getSortingFunc())
	slices.Reverse(sortedData)

	return Manual{data: sortedData}
}

func ParseInput(filename string) (RuleSet, []Manual, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load input: %s\n", err)
	}

	defer f.Close()

	rules := make([]Rule, 0)
	manuals := make([]Manual, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			elements := strings.Split(line, "|")
			rule, _ := strconv.Atoi(elements[0])
			compare, _ := strconv.Atoi(elements[1])
			rules = append(rules, Rule{pageNumber: rule, moreThan: compare})
		} else if strings.Contains(line, ",") {
			elements := strings.Split(line, ",")
			manualData := make([]int, 0)
			for _, el := range elements {
				num, _ := strconv.Atoi(el)
				manualData = append(manualData, num)
			}
			manuals = append(manuals, Manual{data: manualData})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to scan input: %s", err)
	}
	sort.Slice(rules, func(i, j int) bool {
		if rules[i].pageNumber == rules[j].pageNumber {
			return rules[i].moreThan < rules[j].moreThan
		}
		return rules[i].pageNumber < rules[j].pageNumber
	})
	return rules, manuals, nil
}
