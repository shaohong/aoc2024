// https://adventofcode.com/2024/day/2

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func isReportSafe(levels []int) bool {

	nonPositive := make([]bool, len(levels)-1)
	diffs := make([]int, len(levels)-1)
	for i := 0; i < len(levels)-1; i++ {
		diffs[i] = levels[i+1] - levels[i]
		nonPositive[i] = math.Signbit(float64(diffs[i]))
	}

	for i, _ := range nonPositive {
		if nonPositive[0] != nonPositive[i] {
			return false
		}

		if !DiffWithInRange(diffs[i]) {
			return false
		}
	}

	return true
}

func DiffWithInRange(diff int) bool {
	return absInt(diff) <= 3 && absInt(diff) >= 1
}

func IsSafeAfterDamping(levels []int) bool {
	if isReportSafe(levels[1:]) {
		return true
	}

	if isReportSafe(levels[:len(levels)-1]) {
		return true
	}

	// using brute force to check all possible cases
	for i := 1; i < len(levels)-1; i++ {
		// remove level[i]
		newLevels := make([]int, len(levels)-1)
		copy(newLevels, levels[:i]) // copy the first i-1 levels
		copy(newLevels[i:], levels[i+1:])
		if isReportSafe(newLevels) {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nReports := 0
	nSafeReports := 0
	nSafeReportAfterDamping := 0
	for scanner.Scan() {
		text := scanner.Text()

		levelStrings := strings.Fields(text)
		levels := make([]int, len(levelStrings))
		for i, levelStr := range levelStrings {
			fmt.Sscanf(levelStr, "%d", &levels[i])
		}

		nReports++

		if isReportSafe(levels) {
			nSafeReports++
		} else {
			if IsSafeAfterDamping(levels) {
				nSafeReportAfterDamping++

			}
		}

	}

	fmt.Println("nReports:", nReports)

	fmt.Println("nSafeReports:", nSafeReports)

	fmt.Println("nSafeReportsIncludingDamping:", nSafeReportAfterDamping)

}
