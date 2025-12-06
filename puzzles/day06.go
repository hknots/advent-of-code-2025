package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	Register(6, SolveDay06)
}

type Operator string

const (
	Additive Operator = "+"
	Multiply Operator = "*"
)

type mathProblem struct {
	operator Operator
	values   []int
}

func (mp *mathProblem) AppendValue(i int) {
	mp.values = append(mp.values, i)
}

func (mp *mathProblem) CalculateResult() int {
	if len(mp.values) == 0 {
		return 0
	}

	result := mp.values[0]

	for _, v := range mp.values[1:] {
		switch mp.operator {
		case "*":
			result = result * v
		case "+":
			result = result + v
		}
	}
	return result
}

// SolveDay06 Calculates the sum of all math problems
// Ref: https://adventofcode.com/2025/day/6
func SolveDay06(file *os.File) {
	var result int
	scanner := bufio.NewScanner(file)
	mathProblems := make(map[int]mathProblem)

	for scanner.Scan() {
		line := scanner.Text()
		_, err := createMathProblems(mathProblems, line)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	for _, mathProb := range mathProblems {
		result += mathProb.CalculateResult()
	}

	fmt.Printf("Result: %d\n", result)
}

func createMathProblems(mathProblems map[int]mathProblem, line string) (map[int]mathProblem, error) {
	strNumbers := strings.Fields(line)
	scanningNumbers := true

	for i, str := range strNumbers {
		mathProb := mathProblems[i]

		if scanningNumbers {
			value, err := strconv.Atoi(str)
			if err != nil {
				scanningNumbers = false

				// Check if it's an operator
				op, err := parseOperator(str)
				if err != nil {
					return map[int]mathProblem{}, err
				}
				mathProb.operator = op

			} else {
				mathProb.AppendValue(value)
			}
		} else {
			op, err := parseOperator(str)
			if err != nil {
				return map[int]mathProblem{}, err
			}
			mathProb.operator = op
		}

		mathProblems[i] = mathProb
	}

	return mathProblems, nil
}

func parseOperator(s string) (Operator, error) {
	var op Operator
	switch s {
	case "+":
		op = Additive
	case "*":
		op = Multiply
	default:
		return "", fmt.Errorf(fmt.Sprintf("unknown operator found: %s\n", s))
	}
	return op, nil
}
