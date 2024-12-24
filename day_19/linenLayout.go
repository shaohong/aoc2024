// https://adventofcode.com/2024/day/19
package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func ParseInput(input string) (patterns []string, designs []string) {
	parts := strings.Split(input, "\n\n")

	patterns = strings.Split(parts[0], ",")
	for i, pattern := range patterns {
		patterns[i] = strings.TrimSpace(pattern)
	}
	designs = strings.Split(parts[1], "\n")
	return patterns, designs
}

// Is it possible to build the design with the list of patterns?
func IsDesignPossible(design string, patterns []string) bool {
	for _, pattern := range patterns {
		if len(pattern) > len(design) {
			continue
		}

		if pattern == design {
			return true
		}

		if strings.HasPrefix(design, pattern) {
			// check if the rest of the design is possible
			restOfDesign := design[len(pattern):]
			if IsDesignPossible(restOfDesign, patterns) {
				return true
			}
		}
	}
	return false
}

func patternToDict(patterns []string) map[string][]string {
	patternDict := map[string][]string{}
	for _, pattern := range patterns {
		startingChar := pattern[0:1]
		patternDict[startingChar] = append(patternDict[startingChar], pattern)
	}

	for _, patterns := range patternDict {
		slices.SortFunc(patterns, func(a, b string) int {
			return len(b) - len(a)
		})
	}
	return patternDict
}

var DebugLog bool
var KnownDeadEnds = map[string]bool{}
var KnownCounts = map[string]int{}

func IsDesignPossible2(design string, patternDict map[string][]string) bool {
	if _, ok := KnownDeadEnds[design]; ok {
		return false
	}

	if DebugLog {
		fmt.Println("IsDesignPossible2 -- design: ", design)
	}
	firstChar := design[0:1]
	for _, pattern := range patternDict[firstChar] {
		if len(pattern) > len(design) {
			continue
		}

		if pattern == design {
			return true
		}

		if strings.HasPrefix(design, pattern) {
			if DebugLog {
				fmt.Printf("IsDesignPossible2 -- design: %q has prefix %q", design, pattern)
			}
			// check if the rest of the design is possible
			restOfDesign := design[len(pattern):]
			if IsDesignPossible2(restOfDesign, patternDict) {
				return true
			}
		}
	}

	KnownDeadEnds[design] = true
	return false
}

func CountDesignPosibilities(design string, patternDict map[string][]string) int {
	if _, ok := KnownDeadEnds[design]; ok {
		return 0
	}

	if _, ok := KnownCounts[design]; ok {
		return KnownCounts[design]
	}

	firstChar := design[0:1]
	count := 0
	for _, pattern := range patternDict[firstChar] {
		// go through all the patterns and check possibilities
		if len(pattern) > len(design) {
			continue
		}

		if pattern == design {
			count++
		} else {
			if strings.HasPrefix(design, pattern) {
				// check if the rest of the design is possible
				restOfDesign := design[len(pattern):]
				if DebugLog {
					fmt.Printf("CountDesignPosibilities -- design: %q has prefix %q, continue to check %q\n", design, pattern, restOfDesign)

				}
				subCount := CountDesignPosibilities(restOfDesign, patternDict)
				if subCount > 0 {
					count += subCount
				} else {
					if DebugLog {
						fmt.Printf("adding to known dead ends: %q\n", restOfDesign)
					}
					KnownDeadEnds[restOfDesign] = true
				}
			}
		}

	}

	KnownCounts[design] = count
	return count
}

func main() {
	// Read the input
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	patterns, designs := ParseInput(string(data))

	slices.Sort(patterns)
	fmt.Println("patterns: ", patterns)
	fmt.Println("designs: ", designs)

	// partition the patterns by starting character
	patternDict := patternToDict(patterns)
	for k, v := range patternDict {
		fmt.Println(k, ":", v)
	}

	possibleDesigns := 0
	for i, design := range designs {
		if IsDesignPossible2(design, patternDict) {
			possibleDesigns++
			fmt.Println(i, "design: ", design, "possible")
		} else {
			fmt.Println(i, "design: ", design, "impossible")
		}
	}

	// Print the result
	fmt.Println("# of possible design", possibleDesigns)

	fmt.Println("--- Part 2 ---")
	totalPossibility := 0
	clear(KnownDeadEnds)
	DebugLog = false
	for i, design := range designs {
		fmt.Println(i, "design: ", design)
		count := CountDesignPosibilities(design, patternDict)
		fmt.Println(i, "design: ", design, "count:", count)
		totalPossibility += count
	}

	fmt.Println("totalPossibility:", totalPossibility)
}
