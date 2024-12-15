package main

import (
	"fmt"
	"testing"
)

func TestMoveRobot(t *testing.T) {
	robot := Robot{location: Location{2, 4}, velocity: Velocity{2, -3}}
	xLimit, yLimit := 11, 7

	testCases := []struct {
		time     int
		expected Location
	}{
		{1, Location{4, 1}},
		{2, Location{6, 5}},
		{3, Location{8, 2}},
		{4, Location{10, 6}},
		{5, Location{1, 3}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("after %d Seconds", tc.time), func(t *testing.T) {
			got := robot.move(tc.time, xLimit, yLimit)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}

}
