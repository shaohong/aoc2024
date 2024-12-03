package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var thePattern *regexp.Regexp = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

// evaluate the `mul(a,b)` pattern, returns a*b
func EvaluateMul(pattern string) int {
	var a, b int
	fmt.Sscanf(pattern, "mul(%d,%d)", &a, &b)
	return a * b
}

func Part1() {
	scanner := bufio.NewScanner(os.Stdin)
	var total int64 = 0

	for scanner.Scan() {
		// scan input line by line
		line := scanner.Text()

		total += EvaluateText(line)

	}

	fmt.Println("total:", total)
}

// evaluate text containing the `mul(a,b)` pattern, returns the sum of all a*b
func EvaluateText(text string) int64 {
	matchingPatterns := thePattern.FindAllString(text, -1)

	var sum int64 = 0
	for _, pattern := range matchingPatterns {
		sum += int64(EvaluateMul(pattern))
	}
	return sum
}

func EvalulateLineWithStateMachine(line string) int64 {
	// by default, calculation is enabled, until we see the `dont't()` then calculating because false, until we see `do()` again`
	machineEnabled := true
	enableStr := "do()"
	dislableStr := "don't()"

	var lineResult int64 = 0

	for len(line) > 0 {
		if machineEnabled {

			// handle the string before the the next `don't()` pattern
			segmentEndIdx := strings.Index(line, dislableStr)
			if segmentEndIdx == -1 {
				// there is no disable anymore
				lineResult += EvaluateText(line)
				return lineResult
			} else {
				lineResult += EvaluateText(line[:segmentEndIdx])
				// modify line and continue the loop
				line = line[segmentEndIdx+len(dislableStr):]
				machineEnabled = false
				continue
			}

		} else {
			// find the first `do()` pattern
			segmentEndIdx := strings.Index(line, enableStr)
			if segmentEndIdx == -1 {
				// no enabling again, we are done!
				return lineResult
			} else {
				// modify line
				line = line[segmentEndIdx+len(enableStr):]
				machineEnabled = true
				continue
			}

		}

	}

	return lineResult
}

func Part2() {
	var sum int64 = 0

	block := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// scan input line by line
		line := scanner.Text()
		block += line

	}

	sum += EvalulateLineWithStateMachine(block)

	fmt.Println("Part2 sum:", sum)

}

func main() {
	// Part1()

	Part2()
}
