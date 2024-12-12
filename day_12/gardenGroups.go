// https://adventofcode.com/2024/day/12
package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func parseInput(data string) [][]rune {
	var gardenMap [][]rune

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		gardenMap = append(gardenMap, []rune(line))
	}

	return gardenMap
}

type Location struct {
	x int
	y int
}

type LocationGroup struct {
	groupID   string
	locations []Location
}

// Directions for moving up, down, left, right
var directions = [][2]int{
	{0, 1}, {1, 0},
	{0, -1}, {-1, 0},
}

// build a groupID based on the location and gardenMap
func BuildGroupID(gardenMap [][]rune, location Location) string {
	return fmt.Sprintf("%c_%d_%d", gardenMap[location.x][location.y], location.x, location.y)
}

// find the location groups in the garden. return a map of groupID to LocationGroup
// groupID is a string created by the coordinates of the first location in the group
func FindGroups(gardenMap [][]rune) map[string]*LocationGroup {

	groups := make(map[string]*LocationGroup)

	visitedPlots := make(map[Location]bool)

	queue := []Location{}

	for len(visitedPlots) < len(gardenMap)*len(gardenMap[0]) {

		// find the first unvisited plot
		firstUnvisitedPlot := findUnvisitedLocation(gardenMap, visitedPlots)

		// find the groupID of the first unvisited plot, add it to the groups map
		groupID := BuildGroupID(gardenMap, firstUnvisitedPlot)
		groups[groupID] = &LocationGroup{groupID, []Location{firstUnvisitedPlot}}
		visitedPlots[firstUnvisitedPlot] = true

		fmt.Println("firstUnvisitedPlot", firstUnvisitedPlot, "groupID", groupID)

		// do a BFS starting from the firstUnvisitedPlot, add all the connected plots to the group
		queue = append(queue, firstUnvisitedPlot)
		for len(queue) > 0 {
			// pop the first element from the queue
			currentLocation := queue[0]
			queue = queue[1:]

			// check the neighbors of the current location
			for _, dir := range directions {
				newLocation := Location{currentLocation.x + dir[0], currentLocation.y + dir[1]}
				if newLocation.x >= 0 && newLocation.x < len(gardenMap) && newLocation.y >= 0 && newLocation.y < len(gardenMap[0]) {

					// if this neighbor has not been visited before, and it has the same plant as the current location, add it to the group
					if _, ok := visitedPlots[newLocation]; !ok && gardenMap[newLocation.x][newLocation.y] == gardenMap[currentLocation.x][currentLocation.y] {
						// add the new location to the group
						groups[groupID].locations = append(groups[groupID].locations, newLocation)
						// mark the new location as visited
						visitedPlots[newLocation] = true
						// add the new location to the queue for further exploration
						queue = append(queue, newLocation)
					}

				}
			}
		}

	}

	return groups
}

func findUnvisitedLocation(gardenMap [][]rune, visitedPlots map[Location]bool) (firstUnvisitedPlot Location) {
	for i := 0; i < len(gardenMap); i++ {
		for j := 0; j < len(gardenMap[0]); j++ {
			if _, ok := visitedPlots[Location{i, j}]; !ok {
				firstUnvisitedPlot = Location{i, j}
				return
			}
		}
	}
	return
}

func CalculateRegionCost(gardenMap [][]rune, region *LocationGroup) (area int, perimeter int) {
	area = len(region.locations)

	perimeter = 4 * area
	for _, location := range region.locations {

		for _, dir := range directions {
			neighbor := Location{location.x + dir[0], location.y + dir[1]}
			if neighbor.x >= 0 && neighbor.x < len(gardenMap) && neighbor.y >= 0 || neighbor.y < len(gardenMap[0]) {
				if slices.Contains(region.locations, neighbor) {
					perimeter = perimeter - 1
				}
			}
		}

	}

	return
}

func side_count_update(current_type bool) int {
	if current_type {
		return 0
	}
	return 1
}

// calculate the horizontal sides of the region
func CalculateHorizontalSides(gardenMap [][]rune, region *LocationGroup) int {

	// a horizontal side is created when the plot is in the region and the plot above it is not in the region, call it type_1_side, (or vice versa call it type_2_side)
	// side ends when the plot and the one above it is in the same region

	// plots outside the garden boundary are treated as not	in the region

	total_sides := 0

	for row := 0; row <= len(gardenMap); row++ {
		type_1_side := 0
		type_2_side := 0
		type_1_side_ongoing := false
		type_2_side_ongoing := false

		for col := 0; col < len(gardenMap[0]); col++ {
			currentPlot := Location{row, col}
			upperPlot := Location{row - 1, col}

			// check for type_1_side
			if slices.Contains(region.locations, currentPlot) && !slices.Contains(region.locations, upperPlot) {
				// we may or may not update the type_1_side count
				type_1_side += side_count_update(type_1_side_ongoing)
				type_1_side_ongoing = true
				type_2_side_ongoing = false
				continue
			}

			// check for type_2_side
			if !slices.Contains(region.locations, currentPlot) && slices.Contains(region.locations, upperPlot) {
				// we may or may not update the type_2_side count
				type_2_side += side_count_update(type_2_side_ongoing)
				type_2_side_ongoing = true
				type_1_side_ongoing = false
				continue
			}

			// if both plots are in the region or outside the region, we stop the ongoing side
			type_1_side_ongoing = false
			type_2_side_ongoing = false
		}

		// fmt.Println("row", row, "type_1_side", type_1_side, "type_2_side", type_2_side)

		total_sides += type_1_side + type_2_side
	}

	return total_sides

}

func CalculateVerticalSides(gardenMap [][]rune, region *LocationGroup) int {

	// a vertical side is created when the plot is in the region and the plot to the left of it is not in the region, call it type_1_side, (or vice versa call it type_2_side)
	// side ends when the plot and the one to the left of it is in the same region

	// plots outside the garden boundary are treated as not	in the region

	total_sides := 0

	for col := 0; col <= len(gardenMap[0]); col++ {
		type_1_side := 0
		type_2_side := 0
		type_1_side_ongoing := false
		type_2_side_ongoing := false

		for row := 0; row < len(gardenMap); row++ {
			currentPlot := Location{row, col}
			leftPlot := Location{row, col - 1}

			// check for type_1_side
			if slices.Contains(region.locations, currentPlot) && !slices.Contains(region.locations, leftPlot) {
				type_1_side += side_count_update(type_1_side_ongoing)
				type_1_side_ongoing = true
				type_2_side_ongoing = false
				continue
			}

			// check for type_2_side
			if !slices.Contains(region.locations, currentPlot) && slices.Contains(region.locations, leftPlot) {
				type_2_side += side_count_update(type_2_side_ongoing)
				type_2_side_ongoing = true
				type_1_side_ongoing = false
				continue
			}

			type_1_side_ongoing = false
			type_2_side_ongoing = false
		}

		total_sides += type_1_side + type_2_side
	}

	return total_sides
}

func CalculateSides(gardenMap [][]rune, region *LocationGroup) int {
	return CalculateHorizontalSides(gardenMap, region) + CalculateVerticalSides(gardenMap, region)
}

func CalculatePricePart2(gardenMap [][]rune) int {
	plotGroups := FindGroups(gardenMap)

	totalCost := 0
	for _, group := range plotGroups {
		area := len(group.locations)
		horizontalSides := CalculateHorizontalSides(gardenMap, group)
		verticalSides := CalculateVerticalSides(gardenMap, group)

		sides := horizontalSides + verticalSides
		fmt.Println("group", group.groupID, "area", area, "sides", sides, "horizontal:", horizontalSides, "verticalSides", verticalSides)
		totalCost += area * sides
	}

	return totalCost
}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	// read the input into a 2D array
	gardenMap := parseInput(string(data))

	// fmt.Println(gardenMap)

	// plotGroups := FindGroups(gardenMap)

	// totalCost := 0
	// for _, group := range plotGroups {
	// 	area, perimeter := CalculateRegionCost(gardenMap, group)
	// 	fmt.Println("area", area, "perimeter", perimeter)
	// 	totalCost += area * perimeter
	// }

	// fmt.Println("totalCost", totalCost)

	fmt.Println("totalCostPart2", CalculatePricePart2(gardenMap))

}
