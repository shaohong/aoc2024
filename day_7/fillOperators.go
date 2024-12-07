// https://adventofcode.com/2024/day/7

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// we have a bunch of numbers and a goal. we'd like to find operators that when operated on the numbers, the result is the goal
type OperatorProblem struct {
	goal    int64
	numbers []int64
}

func (p OperatorProblem) String() string {
	return fmt.Sprintf("%d: %v", p.goal, p.numbers)
}

func IsOperatorProblemSolvable(operatorProblem OperatorProblem) bool {
	for _, operator := range []string{"+", "*", "||"} {
		var newNumber int64
		switch operator {
		case "+":
			newNumber = operatorProblem.numbers[0] + operatorProblem.numbers[1]
		case "*":
			newNumber = operatorProblem.numbers[0] * operatorProblem.numbers[1]
		case "||":
			newNumberStr := fmt.Sprintf("%d%d", operatorProblem.numbers[0], operatorProblem.numbers[1])
			a, error := strconv.Atoi(newNumberStr)
			if error != nil {
				panic(error)
			} else {
				newNumber = int64(a)
			}
		}

		// we only have these two numbers left
		if len(operatorProblem.numbers) == 2 {
			if newNumber == operatorProblem.goal {
				return true // go achevied
			} else {
				continue // try next operator
			}
		} else { // there were more then two numbers in the problem

			if newNumber > operatorProblem.goal {
				// no solution on this path
				continue
			} else {
				newProblem := OperatorProblem{operatorProblem.goal, append([]int64{newNumber}, operatorProblem.numbers[2:]...)}
				if IsOperatorProblemSolvable(newProblem) {
					return true
				} else {
					continue
				}
			}
		}
	}

	return false
}

func ParseOperatorProblem(line string) OperatorProblem {
	words := strings.Fields(line)

	// throw away the last ':' and convert to int
	goal, _ := strconv.Atoi(words[0][:len(words[0])-1])
	numbers := []int64{}
	for _, word := range words[1:] {
		number, _ := strconv.Atoi(word)
		numbers = append(numbers, int64(number))
	}

	return OperatorProblem{int64(goal), numbers}
}

func main() {

	operatorProblems := []OperatorProblem{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		operatorProblems = append(operatorProblems, ParseOperatorProblem(line))
	}

	fmt.Println("number of operator problems: ", len(operatorProblems))

	solvableProblems := []OperatorProblem{}

	for _, operatorProblem := range operatorProblems {
		if IsOperatorProblemSolvable(operatorProblem) {
			solvableProblems = append(solvableProblems, operatorProblem)
		}
	}

	var sum int64 = 0
	for _, p := range solvableProblems {
		sum += p.goal
	}

	fmt.Println("sum of solvable problems goals: ", sum)

}
