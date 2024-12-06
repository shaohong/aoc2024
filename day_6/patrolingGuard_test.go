package main

import (
	"strings"
	"testing"
)

func TestPatrol(t *testing.T) {
	mapStr := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	matrix := make([][]rune, 0)
	for _, line := range strings.Split(mapStr, "\n") {
		matrix = append(matrix, []rune(line))
	}

	patrolRoute, _ := Patrol(matrix)

	want := 41
	get := len(GetUniqueLocations(patrolRoute))

	if want != get {
		t.Errorf("want %d, get %d", want, get)
	}
}

func TestGetPatrolLoopOpportunities(t *testing.T) {
	mapStr := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	matrix := make([][]rune, 0)
	for _, line := range strings.Split(mapStr, "\n") {
		matrix = append(matrix, []rune(line))
	}

	t.Run("GetLoopOpportunitiesBF", func(t *testing.T) {

		want := 6
		get := GetPatrolLoopOpportunities(matrix)

		if want != get {
			t.Errorf("want %d, get %d", want, get)
		}
	})

}
