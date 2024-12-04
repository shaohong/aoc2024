package main

import (
	"strings"
	"testing"
)

func TestLinesTo2dSlices(t *testing.T) {
	lines := []string{"abc", "def", "ghi"}

	want := [][]rune{{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'}}

	got := LinesTo2dSlices(lines)

	if len(got) != len(want) {
		t.Errorf("Lengths don't match")
	}
	for i := range want {
		for j := range want[i] {
			if want[i][j] != got[i][j] {
				t.Errorf("want[%d][%d] != got[%d][%d]", i, j, i, j)
			}
		}
	}

}

func TestCountXMASPatterns(t *testing.T) {
	block := `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`
	matrix := LinesTo2dSlices(strings.Split(block, "\n"))

	t.Run("CountXMASPatterns given an X location", func(t *testing.T) {
		// the last X location
		xLoc := Location{len(matrix) - 1, len(matrix[0]) - 1}

		want := 2
		got := CountXMASPatternsGivenX(matrix, xLoc)
		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})

	t.Run("Count All XMAS Patterns", func(t *testing.T) {
		want := 18
		got := CountAllXMASPatterns(matrix)
		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}

func TestIsCrossMasAtA(t *testing.T) {
	block := `M.S
.A.
M.S`
	matrix := LinesTo2dSlices(strings.Split(block, "\n"))
	want := true
	got := IsCrossMasAtA(matrix, Location{1, 1})
	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}

}

func TestCountAllCrossMas(t *testing.T) {
	block := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

	matrix := LinesTo2dSlices(strings.Split(block, "\n"))
	want := 9
	got := CountAllCrossMAS(matrix)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
