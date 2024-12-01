package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func simalarityScore(leftlocationIDs []int, rightlocationIDs []int) int64 {
	// calculate the similarity score between two slices of locationIDs
	// the similarity score is the sum of each number on the left times the number of occurences on the right

	// define a map to store the occurences of each locationID in rightlocationIDs
	occurences := make(map[int]int)
	for _, locationID := range rightlocationIDs {
		occurences[locationID]++
	}

	var simalarityScore int64 = 0
	// for each locationID in leftlocationIDs, calculate the similarity score
	for _, locationID := range leftlocationIDs {
		simalarityScore += int64(occurences[locationID] * locationID)
	}

	return simalarityScore
}

func main() {
	// two slices to store the left and right locationIDs
	var leftLocationIDs = make([]int, 0)
	var rightLocationIDs = make([]int, 0)

	// read from standard input until the end of the file
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var left, right int
		text := scanner.Text()
		fmt.Sscanln(text, &left, &right)

		leftLocationIDs = append(leftLocationIDs, left)
		rightLocationIDs = append(rightLocationIDs, right)
	}

	// sort the locationIDs
	slices.Sort(leftLocationIDs)
	slices.Sort(rightLocationIDs)

	// calculate the distance by sum over abs(leftLocationIDs[i] - rightLocationIDs[i])
	var distance int64 = 0
	// for each locationID in leftLocationIDs, calculate the distance to the corresponding locationID in rightLocationIDs
	for i := 0; i < len(leftLocationIDs); i++ {
		if leftLocationIDs[i] >= rightLocationIDs[i] {
			distance += int64(leftLocationIDs[i] - rightLocationIDs[i])
		} else {
			distance += int64(rightLocationIDs[i] - leftLocationIDs[i])
		}
	}
	// print the distance
	fmt.Println("distance: ", distance)

	// print similarity score
	fmt.Println("similarity score: ", simalarityScore(leftLocationIDs, rightLocationIDs))

}
