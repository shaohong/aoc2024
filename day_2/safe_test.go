package main

import (
	"fmt"
	"testing"
)

// To run one specific case: `go test -v -run TestEdgeCase`
func TestEdgeCase(t *testing.T) {
	levels := []int{31, 34, 32, 30, 28, 27, 24, 22}
	fmt.Println("testing levels: ", levels)
	want := true
	got := IsSafeAfterDamping(levels)
	assertCorrectMessage(t, got, want, levels)
}

func TestUnsafeAfterDamping(t *testing.T) {

	t.Run("Unsafe after damping", func(t *testing.T) {
		unsafeLevels := [][]int{
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{3, 4, 7, 9, 8, 9, 9},
		}

		for _, levels := range unsafeLevels {
			want := false
			got := IsSafeAfterDamping(levels)
			assertCorrectMessage(t, got, want, levels)
		}
	})
}

func TestIsSafeAfterDamping(t *testing.T) {

	t.Run("dampening example", func(t *testing.T) {
		levels := []int{1, 3, 2, 4, 5}
		want := true
		got := IsSafeAfterDamping(levels)
		assertCorrectMessage(t, got, want, levels)
	})

	t.Run("all increasing case", func(t *testing.T) {
		levels := []int{1, 2, 3, 4, 8}
		want := true
		got := IsSafeAfterDamping(levels)
		assertCorrectMessage(t, got, want, levels)

	})

	t.Run("all decreasing case", func(t *testing.T) {
		levels := []int{8, 4, 3, 2, 1}
		want := true
		got := IsSafeAfterDamping(levels)
		assertCorrectMessage(t, got, want, levels)
	})

	t.Run("equal levels", func(t *testing.T) {
		levels := []int{1, 2, 3, 3, 4}
		want := true
		got := IsSafeAfterDamping(levels)
		assertCorrectMessage(t, got, want, levels)
	})

	t.Run("more cases", func(t *testing.T) {
		edgeCases := [][]int{
			{8, 6, 4, 4, 1},
			{48, 46, 47, 49, 51, 54, 56},
			{1, 1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5, 5},
			{5, 1, 2, 3, 4, 5},
			{1, 4, 3, 2, 1},
			{1, 6, 7, 8, 9},
			{1, 2, 3, 4, 3},
			{9, 8, 7, 6, 7},
			{7, 10, 8, 10, 11},
			{29, 28, 27, 25, 26, 25, 22, 20}, // we can remove the first 25

			{90, 89, 86, 84, 83, 79},
			{97, 96, 93, 91, 85},
			{29, 26, 24, 25, 21},
			{36, 37, 40, 43, 47},
			{43, 44, 47, 48, 49, 54},
			{35, 33, 31, 29, 27, 25, 22, 18},
			{77, 76, 73, 70, 64},
			{68, 65, 69, 72, 74, 77, 80, 83}, // we can remove 65
			{37, 40, 42, 43, 44, 47, 51},
			{70, 73, 76, 79, 86},
			{75, 77, 72, 70, 69}, // we can remove 77
			{7, 10, 8, 10, 11},   // we can remove the fist 10
		}

		for _, levels := range edgeCases {
			want := true
			fmt.Println("testing edge case: ", levels)
			got := IsSafeAfterDamping(levels)
			assertCorrectMessage(t, got, want, levels)
		}

	})

}

func assertCorrectMessage(t testing.TB, got, want bool, levels []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t for levels %v \n ", got, want, levels)
	}
}
