package main

import (
	"testing"
)

func TestNextSecret(t *testing.T) {
	tests := []struct {
		input    uint
		expected uint
	}{
		{123, 15887950},
		{15887950, 16495136},
		{16495136, 527345},
		{527345, 704524},
		{7753432, 5908254},
	}

	for _, tt := range tests {
		got := nextSecret(tt.input)
		if got != tt.expected {
			t.Errorf("case %v, Expected %v, got %v", tt.input, tt.expected, got)
		}
	}
}

func TestNthSecret(t *testing.T) {
	tests := []struct {
		input    uint
		expected uint
	}{
		{1, 8685429},
		{10, 4700978},
		{100, 15273692},
		{2024, 8667524},
	}

	for _, tt := range tests {
		got := nthSecret(tt.input, 2000)
		if got != tt.expected {
			t.Errorf("case %v, Expected %v, got %v", tt.input, tt.expected, got)
		}
	}
}
