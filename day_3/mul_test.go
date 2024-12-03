package main

import (
	"testing"
)

func TestPatternRecognition(t *testing.T) {
	line := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	got := thePattern.FindAllString(line, -1)

	// Expected matches
	expected := []string{
		"mul(2,4)",
		"mul(5,5)",
		"mul(11,8)",
		"mul(8,5)",
	}
	if len(got) != len(expected) {
		t.Errorf("Expected %d matches, got %d", len(expected), len(got))
	}

	for i, match := range got {
		if match != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], match)
		}
	}
}

func TestEvaluateMul(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected int
	}{
		{"mul(2,4)", 8},
		{"mul(5,5)", 25},
		{"mul(11,8)", 88},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := EvaluateMul(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func TestEvaluateLine(t *testing.T) {
	line := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := int64(48)
	result := EvalulateLineWithStateMachine(line)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
