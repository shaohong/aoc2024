package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetShortestPaths(t *testing.T) {
	start := numPadGraph.GetNode('2')
	end := numPadGraph.GetNode('9')
	got := numPadGraph.GetShortestPaths(start, end)
	expected := 3
	if len(got) != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	for _, p := range got {
		fmt.Println(p.String())
	}

	moveSequence := got[0].ToMoveSequence()
	fmt.Println(moveSequence.String())

}

func TestGetNumericPadMoveSequence(t *testing.T) {
	input := "A029A"
	got := GetMoveSequence(input, numPadGraph)
	expectedSequences := []string{"<A^A>^^AvvvA", "<A^A^>^AvvvA", "<A^A^^>AvvvA"}

	if len(got) != len(expectedSequences) {
		t.Errorf("Expected %v, got %v", expectedSequences, got)
	}

	for _, s := range got {
		fmt.Println(s.String())
		if !slices.Contains(expectedSequences, s.String()) {
			t.Errorf("Unexpected move sequence %v", s.String())
		}
	}
}

func TestGetDirectionalPadMoveSequence(t *testing.T) {
	input := "<A^A>^^AvvvA"
	got := GetMoveSequence(input, dirPadGraph)

	// assert there is no duplications in the result
	unique := make(map[string]bool)
	for _, s := range got {
		if unique[s.String()] {
			t.Errorf("Duplicate move sequence %v", s.String())
		} else {
			unique[s.String()] = true
		}
	}
	// expectedSequences := []string{"v<<A>>^A<A>AvA<^AA>A<vAAA>^A"}

	// if len(got) != len(expectedSequences) {
	// 	t.Errorf("Expected %v, got %v", expectedSequences, got)
	// }

	for _, s := range got {
		fmt.Println(s.String())
	}
}

func TestGetNumberFromCode(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"029A", 29},
		{"980A", 980},
		{"179A", 179},
	}

	for _, tc := range testCases {
		got := GetNumberFromCode(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, got)
		}
	}
}

func TestGetShortestMoveSequenceLength(t *testing.T) {
	input := "A029A"
	got := GetShortestMoveSequenceLength(input, 2)
	expected := uint64(68)
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestCodeComplexity(t *testing.T) {
	input := "029A"
	got := GetCodeComplexity(input, 2)
	expected := uint64(68 * 29)
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestGetShortestMoves(t *testing.T) {
	input := "029A"
	// human using the directional pad to control the robot arm on the numeric pad
	got := GetCodeCost(input, 0)
	expected := int64(12)
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	fmt.Println("")
	// one intermediate directional pad
	got = GetCodeCost(input, 1)
	expected = int64(28)
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	DumpMoveCostCache()

}
