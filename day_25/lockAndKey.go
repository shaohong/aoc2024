// https://adventofcode.com/2024/day/25

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Schematic struct {
	pattern [][]rune
	heights []int
}

func (s Schematic) IsLock() bool {
	// check the top row
	for _, c := range s.pattern[0] {
		if c != '#' {
			return false
		}
	}

	return true
}

func (s Schematic) IsKey() bool {
	return !s.IsLock()
}

func (s Schematic) String() string {

	result := ""
	// for _, row := range s.pattern {
	// 	result += string(row) + "\n"
	// }
	if s.IsLock() {
		result += "Lock: "
	} else {
		result += "Key: "
	}
	for _, height := range s.heights {
		result += fmt.Sprintf("%v ", height)
	}
	return result
}

func (s Schematic) ParseHeights() {
	for col := 0; col < len(s.pattern[0]); col++ {
		height := 0
		startingRow := 1
		endingRow := len(s.pattern) - 1
		dr := +1

		if s.IsKey() {
			startingRow = len(s.pattern) - 1 - 1
			endingRow = 0
			dr = -1
		}
		for row := startingRow; row != endingRow; row += dr {
			if s.pattern[row][col] == '#' {
				height++
			}
		}
		s.heights[col] = height
	}
}

func ParseInput(data string) []Schematic {
	schematicsMaps := strings.Split(data, "\n\n")
	// Parse the patterns
	schematics := make([]Schematic, len(schematicsMaps))
	for i, schematicsMap := range schematicsMaps {

		lines := strings.Split(schematicsMap, "\n")
		pattern := make([][]rune, len(lines))
		for i, line := range lines {
			pattern[i] = []rune(line)
		}
		newSchematic := Schematic{pattern: pattern, heights: make([]int, len(pattern[0]))}
		newSchematic.ParseHeights()
		schematics[i] = newSchematic
	}
	return schematics
}

func LockAndKeyFit(lock Schematic, key Schematic) bool {
	maxHeight := len(lock.pattern) - 2 // do not count the top and bottom rows
	// fmt.Println("Max height: ", maxHeight)

	for i, h := range lock.heights {
		if h+key.heights[i] > maxHeight {
			return false
		}
	}
	return true
}

func main() {
	// Read the input
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	schematics := ParseInput(string(data))
	// Print the schematics
	for _, schematic := range schematics {
		fmt.Println(schematic.String())
	}

	locks := map[int]Schematic{}
	keys := map[int]Schematic{}

	for i, schematic := range schematics {
		if schematic.IsLock() {
			locks[i] = schematic
		} else {
			keys[i] = schematic
		}
	}

	// looking for fits between locks and keys
	totalFits := 0
	for _, lock := range locks {
		for _, key := range keys {
			if LockAndKeyFit(lock, key) {
				totalFits++
			}
		}
	}

	fmt.Println(totalFits)

}
