package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"slices"
	"strings"
)

// ReadFileToBytes reads the contents of a file into a byte array.
func ReadFileToBytes(filePath string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type Location struct {
	x int
	y int
}

// Directions for moving up, down, left, right
var directions = [][2]int{
	{0, 1}, {1, 0},
	{0, -1}, {-1, 0},
}

func getCurrptedLocations(data string) []Location {
	var corrupted_locations []Location
	lines := strings.Split(data, "\n")

	for i := 0; i < len(lines); i++ {
		// Generate a random location
		var x, y int
		fmt.Sscanf(lines[i], "%d,%d", &x, &y)
		corrupted_locations = append(corrupted_locations, Location{x, y})
	}
	return corrupted_locations
}

func FindShortestExitPath(grid_size int, corrupted_locations []Location) int {
	startLocation := Location{0, 0}
	endLocation := Location{grid_size - 1, grid_size - 1}

	// use BFS to find the shortest path from startLocation to endLocation
	// location in the corrupted_locations list are blocked
	visited := map[Location]bool{}
	distance := map[Location]int{}
	distance[startLocation] = 0
	queue := []Location{startLocation}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == endLocation {
			return distance[endLocation]
		}

		for _, direction := range directions {
			next := Location{current.x + direction[0], current.y + direction[1]}
			if next.x < 0 || next.x >= grid_size || next.y < 0 || next.y >= grid_size {
				continue
			}
			if slices.Contains(corrupted_locations, next) {
				continue
			}

			if visited[next] {
				continue
			}

			visited[next] = true
			distance[next] = distance[current] + 1
			queue = append(queue, next)
		}
	}

	return -1
}

func GetConnectedComponents(grid_size int, blockedLocations []Location) map[Location][]Location {
	fmt.Println("Getting connected components: len(blockedLocations) ", len(blockedLocations))
	connectedComponents := map[Location][]Location{}

	for i := 0; i < grid_size; i++ {
		for j := 0; j < grid_size; j++ {
			currentLocation := Location{j, i}

			// skip blockers
			if slices.Contains(blockedLocations, currentLocation) {
				continue
			}

			// see if current location's neighbors are already in some component
			foundExistingComponent := false
			existingComponentID := Location{-1, -1}
			for _, direction := range directions {
				neighbor := Location{currentLocation.x + direction[0], currentLocation.y + direction[1]}

				if neighbor.x < 0 || neighbor.x >= grid_size || neighbor.y < 0 || neighbor.y >= grid_size {
					continue
				}
				if slices.Contains(blockedLocations, neighbor) {
					continue
				}

				for componentID, locations := range connectedComponents {
					if slices.Contains(locations, neighbor) {
						if !foundExistingComponent {
							foundExistingComponent = true
							existingComponentID = componentID
							// add current location to the existing component
							connectedComponents[existingComponentID] = append(connectedComponents[existingComponentID], currentLocation)
						} else { // already found one existing component, need to merge components
							if existingComponentID != componentID {
								fmt.Println("Merging components: ", existingComponentID, " and ", componentID)
								connectedComponents[existingComponentID] = append(connectedComponents[existingComponentID], connectedComponents[componentID]...)
								delete(connectedComponents, componentID)
							}

						}
					}
				}
			}

			if !foundExistingComponent {
				// adding a new component
				connectedComponents[currentLocation] = []Location{currentLocation}
			}
		}
	}
	return connectedComponents

}

func Part2(grid_size int, allCorruptions []Location) (firstBlocker Location) {

	// given the grid
	for corruptionLength := 1025; corruptionLength < len(allCorruptions); corruptionLength++ {
		fmt.Println("corruptionLength: ", corruptionLength)
		shortedSteps := FindShortestExitPath(grid_size, allCorruptions[:corruptionLength])
		if shortedSteps == -1 {
			return allCorruptions[corruptionLength-1]
		}
	}

	return Location{-1, -1}
}

func main() {
	inputFilePath := "input.txt"
	number_of_corrupted_locations := 1024
	// number_of_corrupted_locations := 12
	grid_size := 71
	data, err := ReadFileToBytes(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	corrupted_locations := getCurrptedLocations(string(data))

	shortest_exit_path := FindShortestExitPath(grid_size, corrupted_locations[:number_of_corrupted_locations])

	fmt.Println("Shortest exit path length: ", shortest_exit_path)

	fmt.Println("--- part 2 ---")
	firstBlocker := Part2(grid_size, corrupted_locations)
	fmt.Println("First blocker: ", firstBlocker)
}
