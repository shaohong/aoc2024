// https://adventofcode.com/2024/day/4

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Location struct {
	row, col int
}

func LinesTo2dSlices(lines []string) [][]rune {
	result := make([][]rune, 0)

	for _, line := range lines {
		result = append(result, []rune(line))
	}

	return result
}

// given the matrix and the location of an 'X', find the 'XMAS' pattern in the horizontal, vertical, and diagonal directions
func CountXMASPatternsGivenX(matrix [][]rune, xLoc Location) int {
	// the boundary value for row and column
	rowMin, colMin := 0, 0
	rowMax, colMax := len(matrix)-1, len(matrix[0])-1

	patternsCoordinates := [][]Location{
		{Location{xLoc.row, xLoc.col}, Location{xLoc.row, xLoc.col + 1}, Location{xLoc.row, xLoc.col + 2}, Location{xLoc.row, xLoc.col + 3}},
		{Location{xLoc.row, xLoc.col}, Location{xLoc.row + 1, xLoc.col + 1}, Location{xLoc.row + 2, xLoc.col + 2}, Location{xLoc.row + 3, xLoc.col + 3}},
		{Location{xLoc.row, xLoc.col}, Location{xLoc.row + 1, xLoc.col}, Location{xLoc.row + 2, xLoc.col}, Location{xLoc.row + 3, xLoc.col}},
		{Location{xLoc.row, xLoc.col}, Location{xLoc.row + 1, xLoc.col - 1}, Location{xLoc.row + 2, xLoc.col - 2}, Location{xLoc.row + 3, xLoc.col - 3}},

		{Location{xLoc.row, xLoc.col}, Location{xLoc.row, xLoc.col - 1}, Location{xLoc.row, xLoc.col - 2}, Location{xLoc.row, xLoc.col - 3}},
		{Location{xLoc.row, xLoc.col}, Location{xLoc.row - 1, xLoc.col - 1}, Location{xLoc.row - 2, xLoc.col - 2}, Location{xLoc.row - 3, xLoc.col - 3}},
		{Location{xLoc.row, xLoc.col}, Location{xLoc.row - 1, xLoc.col}, Location{xLoc.row - 2, xLoc.col}, Location{xLoc.row - 3, xLoc.col}},
		{Location{xLoc.row, xLoc.col}, Location{xLoc.row - 1, xLoc.col + 1}, Location{xLoc.row - 2, xLoc.col + 2}, Location{xLoc.row - 3, xLoc.col + 3}},
	}

	count := 0
	// for each pattern, check if the pattern is valid, fetch the characters according to pattern coordinates and see if it matches 'XMAS'
	for _, patternCoordinates := range patternsCoordinates {
		// check if the pattern is within the boundary
		validPattern := true
		for _, loc := range patternCoordinates {
			if loc.row < rowMin || loc.row > rowMax || loc.col < colMin || loc.col > colMax {
				validPattern = false
				break
			}
		}

		if !validPattern {
			continue
		}

		// fetch the characters according to pattern coordinates
		patternStr := ""
		for _, loc := range patternCoordinates {
			patternStr += string(matrix[loc.row][loc.col])
		}

		if patternStr == "XMAS" {
			count++
		}
	}

	return count
}

func CountAllXMASPatterns(matrix [][]rune) int {
	// find the location of 'X' in the matrix
	xLocations := make([]Location, 0)

	for i, row := range matrix {
		for j, col := range row {
			if col == 'X' {
				xLocations = append(xLocations, Location{i, j})
			}
		}
	}

	fmt.Println("Occurence of 'X': ", len(xLocations))

	totalXMASPatterns := 0
	// for each occurance of 'X', find the 'XMAS' pattern in the matrix
	for _, xLoc := range xLocations {
		totalXMASPatterns += CountXMASPatternsGivenX(matrix, xLoc)
	}

	return totalXMASPatterns
}

func Part1(matrix [][]rune) {
	fmt.Println("CountAllXMASPatterns: ", CountAllXMASPatterns(matrix))
}

func CountAllCrossMAS(matrix [][]rune) int {
	// find the location of 'A' in the matrix
	aLocations := make([]Location, 0)
	for i, row := range matrix {
		for j, col := range row {
			if col == 'A' {
				aLocations = append(aLocations, Location{i, j})
			}
		}
	}
	fmt.Println("Occurence of 'A': ", len(aLocations))

	crossMasCount := 0
	for _, aLoc := range aLocations {
		if IsCrossMasAtA(matrix, aLoc) {
			crossMasCount++
		}
	}

	return crossMasCount
}

// given a matrix and an 'A' location, find two MAS in the shape of an X, e.g. the following is a valid CrossMAS
// M.S
// .A.
// M.S
func IsCrossMasAtA(matrix [][]rune, aLoc Location) bool {
	// the boundary value for row and column
	rowMin, colMin := 0, 0
	rowMax, colMax := len(matrix)-1, len(matrix[0])-1

	masks := [][]Location{
		{Location{aLoc.row - 1, aLoc.col - 1}, Location{aLoc.row, aLoc.col}, Location{aLoc.row + 1, aLoc.col + 1}},
		{Location{aLoc.row - 1, aLoc.col + 1}, Location{aLoc.row, aLoc.col}, Location{aLoc.row + 1, aLoc.col - 1}},
	}

	for _, mask := range masks {
		patternStr := ""

		for _, loc := range mask {
			// check if locations in the masks is within the boundary. If there is any location outside the boundary, we won't have a cross-MAS pattern, return false
			if loc.row < rowMin || loc.row > rowMax || loc.col < colMin || loc.col > colMax {
				return false
			} else {
				patternStr += string(matrix[loc.row][loc.col])
			}
		}

		if patternStr != "MAS" && patternStr != "SAM" {
			return false
		}
	}

	return true
}

func Part2(matrix [][]rune) {
	fmt.Println("Part2")
	fmt.Println("CountAllCrossMAS: ", CountAllCrossMAS(matrix))

}

func main() {

	// read input, line by line
	lines := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// scan input line by line
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Println("number of lines in input:", len(lines))

	matrix := LinesTo2dSlices(lines)

	Part1(matrix)

	Part2(matrix)

}
