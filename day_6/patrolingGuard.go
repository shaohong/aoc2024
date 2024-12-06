// https://adventofcode.com/2024/day/6
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const Obstacle = '#'
const Guard = '^' // Guard's initial symbol

type Location struct {
	row, col int
}

// a visiting record saves the location and the visiting direction
type VisitingRecord struct {
	location  Location
	direction rune
}

var nextDirections = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

// from the current guardLocation, along the guardDirection, patrol until meet an obstacle, return the visited location and the next guardLocation and direction
func PatrolInOneDirection(matrix [][]rune, guardLocation Location, guardDirection rune) (visitedLocations []Location, nextLocation Location, nextDirection rune) {
	visitedLocations = make([]Location, 0)
	visitedLocations = append(visitedLocations, guardLocation)
	nextLocation = guardLocation
	nextDirection = guardDirection
	switch guardDirection {
	case '^': // up
		for i := guardLocation.row - 1; i >= 0; i-- {
			if matrix[i][guardLocation.col] != Obstacle {
				nextLocation = Location{i, guardLocation.col}
				visitedLocations = append(visitedLocations, nextLocation)
			} else {
				// found obstacle, change direction
				nextDirection = nextDirections[guardDirection]

				// // pop out the last visitedLocation because guard will start from there next time and in the nextDirection
				// visitedLocations = visitedLocations[:len(visitedLocations)-1]
				break
			}
		}
	case '>': // move right
		for j := guardLocation.col + 1; j < len(matrix[guardLocation.row]); j++ {
			if matrix[guardLocation.row][j] != Obstacle {
				nextLocation = Location{guardLocation.row, j}
				visitedLocations = append(visitedLocations, nextLocation)
			} else {
				nextDirection = nextDirections[guardDirection]
				// visitedLocations = visitedLocations[:len(visitedLocations)-1]
				break
			}
		}
	case 'v': // move down
		for i := guardLocation.row + 1; i < len(matrix); i++ {
			if matrix[i][guardLocation.col] != Obstacle {
				nextLocation = Location{i, guardLocation.col}
				visitedLocations = append(visitedLocations, nextLocation)
			} else {
				nextDirection = nextDirections[guardDirection]
				// visitedLocations = visitedLocations[:len(visitedLocations)-1]
				break
			}
		}
	case '<': // move left
		for j := guardLocation.col - 1; j >= 0; j-- {
			if matrix[guardLocation.row][j] != Obstacle {
				nextLocation = Location{guardLocation.row, j}
				visitedLocations = append(visitedLocations, nextLocation)
			} else {
				nextDirection = nextDirections[guardDirection]
				// visitedLocations = visitedLocations[:len(visitedLocations)-1]
				break
			}
		}
	}

	return
}

// guard patrols the map. Return the locations visited by the guard, (location is the key and patroling direction is the value)
func Patrol(matrix [][]rune) (patrolRoute []VisitingRecord, loopFormed bool) {

	patrolRoute = make([]VisitingRecord, 0)
	patrolRecords := make(map[VisitingRecord]rune)

	loopFormed = false

	// guard's initial direction
	guardDirection := '^'

	// find the guard's initial location
	var guardLocation Location
	for i, row := range matrix {
		for j, cell := range row {
			if cell == Guard {
				guardLocation = Location{i, j}
				break
			}
		}
	}

	// patrol the map
	for {
		visitedLocations, nextLocation, nextDirection := PatrolInOneDirection(matrix, guardLocation, guardDirection)

		// save the patrolled locations
		for _, location := range visitedLocations {
			newRecord := VisitingRecord{location, guardDirection}
			patrolRoute = append(patrolRoute, newRecord)

			// if the same record appeared before, we have a loop!
			if _, ok := patrolRecords[newRecord]; ok {
				loopFormed = true
				return
			} else {
				patrolRecords[newRecord] = guardDirection
			}
		}

		if nextDirection == guardDirection {
			// no change of direction! guard is at boundary, done
			return
		}

		// update guard's location and direction for next round of patrolling
		guardDirection = nextDirection
		guardLocation = nextLocation
	}

}

func GetUniqueLocations(allpatrolledLocations []VisitingRecord) []Location {
	uniqueLocations := make([]Location, 0, len(allpatrolledLocations))
	for _, visitingRecord := range allpatrolledLocations {
		if !slices.Contains(uniqueLocations, visitingRecord.location) {
			uniqueLocations = append(uniqueLocations, visitingRecord.location)
		}
	}
	return uniqueLocations
}

func CopyMatrix(matrix [][]rune) [][]rune {
	newMatrix := make([][]rune, len(matrix))
	for i, row := range matrix {
		newMatrix[i] = make([]rune, len(row))
		copy(newMatrix[i], row)
	}
	return newMatrix
}

func GetPatrolLoopOpportunities(matrix [][]rune) int {

	patrolRoute, _ := Patrol(matrix)

	result := 0
	// find all candidate obstacle locations
	// candidaeObstacleLocations should only be on the original Patrol Route, so that the patrol route is changed to possibly forming a loop
	allpatrolledLocations := GetUniqueLocations(patrolRoute)

	for _, obstacleLocation := range allpatrolledLocations[1:] { // exclude the guard's starting position
		// create a new matrix with the obstacle
		newMatrix := CopyMatrix(matrix)
		newMatrix[obstacleLocation.row][obstacleLocation.col] = Obstacle

		// patrol the new matrix
		_, loopFormed := Patrol(newMatrix)

		if loopFormed {
			result++
		}
	}

	return result
}

func main() {
	// read the input into 2d matrix
	matrix := make([][]rune, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	patrolRoutes, _ := Patrol(matrix)

	// print the number of patrolled locations
	fmt.Println("unique patrolled locations", len(GetUniqueLocations(patrolRoutes)))

	fmt.Println("candidate locations, if guard change direction there, to form loop", GetPatrolLoopOpportunities(matrix))

}
