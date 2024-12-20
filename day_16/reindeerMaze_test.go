package main

import (
	"fmt"
	"testing"
)

func TestGetAllPaths(t *testing.T) {
	// test case from the problem
	maze := ParseInput(`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`)

	paths := maze.GetAllPaths(maze.GetStartNode(), maze.GetEndNode())

	// check if the number of paths is correct
	if !(len(paths) > 0) {
		t.Errorf("Expected 4 paths, got %d", len(paths))
	}
	fmt.Println("Number of paths:", len(paths))

	for i, path := range paths {
		fmt.Printf("Path %d: %v\n", i, path)
	}
}

func TestRotationCost(t *testing.T) {
	got := east.rotationCost(north)
	expected := 1000
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}

	if north.rotationCost(south) != 2000 {
		t.Errorf("Expected 2000, got %d", north.rotationCost(south))
	}

	// test the rotation cost of the same direction
	if north.rotationCost(north) != 0 {
		t.Errorf("Expected 0, got %d", north.rotationCost(north))
	}

	if north.rotationCost(west) != 1000 {
		t.Errorf("Expected 1000, got %d", north.rotationCost(east))
	}
}

func TestCalculateNewFacingDirection(t *testing.T) {
	fromNode := Node{Location{row: 1, col: 1}, '.'}
	toNode := Node{Location{row: 1, col: 2}, '.'}
	// test the rotation cost of the same direction
	if calculateNewFacingDirection(fromNode, toNode) != east {
		t.Errorf("Expected east, got %v", calculateNewFacingDirection(fromNode, toNode))
	}

	toNode = Node{Location{row: 2, col: 1}, '.'}
	if calculateNewFacingDirection(fromNode, toNode) != south {
		t.Errorf("Expected south, got %v", calculateNewFacingDirection(fromNode, toNode))
	}

	toNode = Node{Location{row: 0, col: 1}, '.'}
	if calculateNewFacingDirection(fromNode, toNode) != north {
		t.Errorf("Expected north, got %v", calculateNewFacingDirection(fromNode, toNode))
	}

	toNode = Node{Location{row: 1, col: 0}, '.'}
	if calculateNewFacingDirection(fromNode, toNode) != west {
		t.Errorf("Expected west, got %v", calculateNewFacingDirection(fromNode, toNode))
	}
}

func TestCalculateMinCostPath(t *testing.T) {
	maze := ParseInput(`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`)

	start := maze.GetStartNode()
	end := maze.GetEndNode()
	// find the path with the minimum cost
	minCost, minPath := maze.FindMinCost(start, end)
	expectedCost := 7036
	fmt.Printf("Minimum cost: %d\n", minCost)
	fmt.Printf("Minimum cost path: %v\n", minPath)
	if minCost != uint64(expectedCost) {
		t.Errorf("Expected %d, got %d", expectedCost, minCost)
	}

}

func TestDijkstra(t *testing.T) {
	maze := ParseInput(`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`)

	start := maze.GetStartNode()
	end := maze.GetEndNode()
	got, minPaths := maze.Diijkstra(start, end)
	// find the path with the minimum cost
	expectedCost := 7036
	if got != uint64(expectedCost) {
		t.Errorf("Expected %d, got %d", expectedCost, got)
	}

	fmt.Printf("minPaths: %v", minPaths)

}
