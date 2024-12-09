// https://adventofcode.com/2024/day/8
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	x, y int
}

// Recursive finds and returns the greatest common divisor of a given integer.
func GCD(a, b int) int {
	if a == b {
		return a
	}

	if a < b {
		a, b = b, a
	}

	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindAllColinearPoints(A, B Coordinate, maxX, maxY int) (colinearPoints []Coordinate) {
	// Calculate the differences
	dx := B.x - A.x
	dy := B.y - A.y

	// Calculate the greatest common divisor
	gcd := GCD(abs(dx), abs(dy))

	// Normalize the differences
	dx /= int(gcd)
	dy /= int(gcd)

	// Generate all colinear points of with A and B, within the bounds
	// C = A + lambda * (dx, dy)
	colinearPoints = make([]Coordinate, 0, maxX*maxY)
	for x, y := A.x, A.y; x < maxX && y < maxY && x >= 0 && y >= 0; x, y = x+dx, y+dy {
		colinearPoints = append(colinearPoints, Coordinate{x, y})
	}

	for x, y := A.x, A.y; x < maxX && y < maxY && x >= 0 && y >= 0; x, y = x-dx, y-dy {
		colinearPoints = append(colinearPoints, Coordinate{x, y})
	}
	return colinearPoints
}

func isAlphaNumeric(c rune) bool {
	// Check if the byte value falls within the range of alphanumeric characters
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func ParseAntennaMap(data string) [][]rune {
	antennaMap := [][]rune{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		antennaMap = append(antennaMap, []rune(line))
	}

	return antennaMap
}

// find the coordinates of all radars
func GetRadarLocations(antennaMap [][]rune) map[rune][]Coordinate {
	radarLocations := map[rune][]Coordinate{}
	for i, row := range antennaMap {
		for j, cell := range row {
			if isAlphaNumeric(cell) {
				if _, ok := radarLocations[cell]; !ok {
					radarLocations[cell] = []Coordinate{}
				}
				radarLocations[cell] = append(radarLocations[cell], Coordinate{i, j})
			}
		}
	}
	return radarLocations
}

func FindAntiNodesPart1(antennaMap [][]rune) (antiNodes map[Coordinate]bool) {
	radarLocations := GetRadarLocations(antennaMap)

	antiNodes = make(map[Coordinate]bool)

	for _, locations := range radarLocations {
		// fmt.Printf("find anti nodes for radar type %c \n", radarType)
		if len(locations) == 1 {
			continue
		}

		for _, locationA := range locations {
			for _, locationB := range locations {

				if locationA == locationB {
					continue
				}

				node1_x := locationA.x - (locationB.x - locationA.x)
				node1_y := locationA.y - (locationB.y - locationA.y)
				if node1_x >= 0 && node1_x < len(antennaMap) && node1_y >= 0 && node1_y < len(antennaMap[0]) {
					antiNodes[Coordinate{node1_x, node1_y}] = true
				}

				node2_x := locationB.x - (locationA.x - locationB.x)
				node2_y := locationB.y - (locationA.y - locationB.y)
				if node2_x >= 0 && node2_x < len(antennaMap) && node2_y >= 0 && node2_y < len(antennaMap[0]) {
					antiNodes[Coordinate{node2_x, node2_y}] = true
				}
			}
		}
	}

	return
}

func FindAntiNodesPart2(antennaMap [][]rune) (antiNodes map[Coordinate]bool) {
	radarLocations := GetRadarLocations(antennaMap)

	antiNodes = make(map[Coordinate]bool)

	for _, locations := range radarLocations {
		// fmt.Printf("find anti nodes for radar type %c \n", radarType)
		if len(locations) == 1 {
			continue
		}

		for _, locationA := range locations {
			for _, locationB := range locations {

				if locationA == locationB {
					continue
				}

				colinearPoints := FindAllColinearPoints(locationA, locationB, len(antennaMap), len(antennaMap[0]))
				for _, point := range colinearPoints {
					if point.x >= 0 && point.x < len(antennaMap) && point.y >= 0 && point.y < len(antennaMap[0]) {
						antiNodes[point] = true
					}
				}
			}
		}
	}

	return
}

func main() {
	// read the input to 2d slice of runes
	data := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		data += line + "\n"
	}
	antennaMap := ParseAntennaMap(data)
	// print the input
	fmt.Println("Antenna size, rows:", len(antennaMap), ", columns:", len(antennaMap[0]))

	// find the coordinates of all radars
	radarLocations := GetRadarLocations(antennaMap)
	fmt.Println("Radar types", len(radarLocations))

	totalRadarCount := 0
	for _, locations := range radarLocations {
		totalRadarCount += len(locations)
	}
	fmt.Println("Total radars", totalRadarCount)

	antiNodes1 := FindAntiNodesPart1(antennaMap)
	fmt.Println("Antinodes count", len(antiNodes1))

	antiNodes2 := FindAntiNodesPart2(antennaMap)
	fmt.Println("Antinodes count in Part2:", len(antiNodes2))
}
