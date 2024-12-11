package main

import (
	"testing"
)

func TestStoneMutation(t *testing.T) {
	tests := []struct {
		name     string
		number   uint64
		expected []uint64
	}{
		{"zero", 0, []uint64{1}},
		{"even", 1000, []uint64{10, 0}},
		{"odd", 12345, []uint64{12345 * 2024}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StoneMutation(tt.number)
			if len(got) != len(tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Fatalf("expected %v, got %v", tt.expected, got)
				}
			}
		})
	}
}

func TestGetNumberOfStonesAfterMutation(t *testing.T) {
	stones := []uint64{125, 17}
	stoneMap := SliceToMap(stones)

	want := uint64(22)
	got := GetNumberOfStonesAfterMutation(stoneMap, 6)
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}

	want = uint64(55312)
	got = GetNumberOfStonesAfterMutation(stoneMap, 25)
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
