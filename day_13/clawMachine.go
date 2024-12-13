// https://adventofcode.com/2024/day/13

package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Puzzle struct {
	buttonA [2]uint64
	buttonB [2]uint64
	prize   [2]uint64
}

func (p Puzzle) String() string {
	return fmt.Sprintf("buttonA: %v, buttonB: %v, prize: %v", p.buttonA, p.buttonB, p.prize)
}

func (p Puzzle) Solve() [2]uint64 {
	// solve the puzzle
	b := mat.NewDense(2, 1, []float64{
		float64(p.prize[0]),
		float64(p.prize[1]),
	})

	a := mat.NewDense(2, 2, []float64{
		float64(p.buttonA[0]), float64(p.buttonB[0]),
		float64(p.buttonA[1]), float64(p.buttonB[1]),
	})

	var x mat.Dense
	// solve a*x=b
	err := x.Solve(a, b)

	if err == nil {
		// Print the result using the formatter.
		fx := mat.Formatted(&x, mat.Prefix("    "), mat.Squeeze())
		fmt.Printf("x = %.1f\n", fx)

		// buttonA press
		pa := uint64(math.Round(x.At(0, 0)))
		pb := uint64(math.Round(x.At(1, 0)))

		if pa*p.buttonA[0]+pb*p.buttonB[0] == p.prize[0] && pa*p.buttonA[1]+pb*p.buttonB[1] == p.prize[1] {
			return [2]uint64{pa, pb}
		} else {
			return [2]uint64{0, 0}
		}
	}
	return [2]uint64{0, 0}
}

func ParseProblem(problem string) Puzzle {
	lines := strings.Split(problem, "\n")

	var puzzle Puzzle

	fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &puzzle.buttonA[0], &puzzle.buttonA[1])
	fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &puzzle.buttonB[0], &puzzle.buttonB[1])
	fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &puzzle.prize[0], &puzzle.prize[1])

	return puzzle
}

func ParseProblem_2(problem string) Puzzle {
	lines := strings.Split(problem, "\n")

	var puzzle Puzzle

	fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &puzzle.buttonA[0], &puzzle.buttonA[1])
	fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &puzzle.buttonB[0], &puzzle.buttonB[1])
	fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &puzzle.prize[0], &puzzle.prize[1])

	puzzle.prize[0] += 10000000000000
	puzzle.prize[1] += 10000000000000

	return puzzle
}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	problems := strings.Split(string(data), "\n\n")

	fmt.Println("we have", len(problems), "problems")

	buttonAPressed := uint64(0)
	buttonBPressed := uint64(0)
	for i, problem := range problems {

		puzzle := ParseProblem_2(problem)

		solution := puzzle.Solve()

		if solution[0] == 0 && solution[1] == 0 {
			fmt.Println("puzzle", i, ":", puzzle, "has no solution")
			continue
		}

		buttonAPressed += solution[0]
		buttonBPressed += solution[1]

	}

	buttonACost := 3
	buttonBCost := 1
	fmt.Println("token price:", buttonAPressed*uint64(buttonACost)+buttonBPressed*uint64(buttonBCost))

}
