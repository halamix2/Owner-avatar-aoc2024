package common

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"klaidliadon.dev/next"
)

func operators() []rune {
	// as close to const slice as we'll get
	return []rune{'+', '*', '|'}
}

type Equation struct {
	result   int
	operands []int
}

func (e Equation) String() string {
	return fmt.Sprintf("%d\t=\t%d", e.result, e.operands)
}

func (e Equation) getResult(operators []rune) int {
	result := e.operands[0]
	e.operands = e.operands[1:]

	for len(operators) > 0 {
		switch operators[0] {
		case '+':
			result += e.operands[0]
		case '*':
			result *= e.operands[0]
		case '|':
			for i, operator := range operators {
				if operator == '|' {
					result, _ = strconv.Atoi(fmt.Sprintf("%d%d", result, e.operands[i]))

					break
				}

			}
		}
		if len(operators) <= 1 {
			break
		}
		operators = operators[1:]
		e.operands = e.operands[1:]
	}

	return result
}

func (e Equation) IsCorrect(operations []interface{}) bool {
	// I didn't want to implement permutations myself, oh well
	for currentOperators := range next.Permutation(operations, len(e.operands)-1, true) {

		//convert to soemthing more usable
		operators := make([]rune, len(e.operands)-1)
		for i := range operators {
			operators[i] = rune(reflect.ValueOf(currentOperators).Index(i).Interface().(int32))
		}
		result := e.getResult(operators)
		if e.result == result {
			return true
		}
	}

	return false
}

func (e Equation) GetResult() int {
	return e.result
}

func ParseInput(filename string) ([]Equation, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []Equation{}, fmt.Errorf("failed to load input: %s\n", err)
	}
	defer f.Close()

	equations := make([]Equation, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		elements := strings.Fields(scanner.Text())
		operands := make([]int, 0)
		for i, element := range elements {
			if i == 0 {
				continue
			}
			operand, _ := strconv.Atoi(element)
			operands = append(operands, operand)
		}

		result, _ := strconv.Atoi(strings.ReplaceAll(elements[0], ":", ""))
		equations = append(equations, Equation{result: result, operands: operands})
	}

	if err := scanner.Err(); err != nil {
		return []Equation{}, fmt.Errorf("failed to scan input: %s", err)
	}
	return equations, nil
}
