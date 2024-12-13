package main

import (
	"testing"
)

func TestSolvePuzzle(t *testing.T) {
	// create a puzzle
	puzzle := Puzzle{
		buttonA: [2]uint64{94, 34},
		buttonB: [2]uint64{22, 67},
		prize:   [2]uint64{8400, 5400},
	}
	// solve the puzzle
	got := puzzle.Solve()
	want := [2]uint64{80, 40}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	t.Run("TestPuzzlWithNosolution", func(t *testing.T) {
		puzzle2 := Puzzle{
			buttonA: [2]uint64{26, 66},
			buttonB: [2]uint64{67, 21},
			prize:   [2]uint64{12748, 12176},
		}
		got := puzzle2.Solve()
		want := [2]uint64{0, 0}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
