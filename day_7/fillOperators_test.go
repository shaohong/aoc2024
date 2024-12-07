package main

import (
	"testing"
)

func TestIsOperatorProblemSolvable(t *testing.T) {

	// test cases
	testCases := []struct {
		operatorProblem OperatorProblem
		expected        bool
	}{
		{ParseOperatorProblem("190: 10 19"), true},
		{ParseOperatorProblem("3267: 81 40 27"), true},
		{ParseOperatorProblem("83: 17 5"), false},
		{ParseOperatorProblem("156: 15 6"), true},
		{ParseOperatorProblem("7290: 6 8 6 15"), true},
		{ParseOperatorProblem("161011: 16 10 13"), false},
		{ParseOperatorProblem("192: 17 8 14"), true},
		{ParseOperatorProblem("21037: 9 7 18 13"), false},
		{ParseOperatorProblem("292: 11 6 16 20"), true},
	}

	for _, tc := range testCases {
		tc_name := tc.operatorProblem.String()
		t.Run(tc_name, func(t *testing.T) {
			result := IsOperatorProblemSolvable(tc.operatorProblem)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}

}
