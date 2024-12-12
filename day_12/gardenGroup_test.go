package main

import (
	"fmt"
	"testing"
)

func TestFindGroups(t *testing.T) {
	inputStr := `AAAA
BBCD
BBCC
EEEC`
	gardenMap := parseInput(inputStr)

	fmt.Println("gardenMap", gardenMap)

	if len(gardenMap) != 4 {
		t.Errorf("Expected 4 rows, got %d", len(gardenMap))
	}

	if len(gardenMap[0]) != 4 {
		t.Errorf("Expected 4 columns, got %d", len(gardenMap[0]))
	}

	if gardenMap[0][0] != 'A' {
		t.Errorf("Expected A, got %c", gardenMap[0][0])
	}

	// fmt.Println(gardenMap)
	groups := FindGroups(gardenMap)

	if len(groups) != 5 {
		t.Errorf("Expected 5 groups, got %d", len(groups))
	}

	group1 := groups["A_0_0"]
	expectedArea, expectedPerimeter := 4, 10
	gotArea, gotPerimeter := CalculateRegionCost(gardenMap, group1)
	if gotArea != expectedArea || gotPerimeter != expectedPerimeter {
		t.Errorf("Expected area %d and perimeter %d, got %d and %d", expectedArea, expectedPerimeter, gotArea, gotPerimeter)
	}

	t.Run("CalculateVerticalSides", func(t *testing.T) {
		// CalculateHorizontalSides
		expectedHorizontalSides := 2
		got := CalculateHorizontalSides(gardenMap, group1)
		if got != expectedHorizontalSides {
			t.Errorf("Expected %d horizontal sides, got %d", expectedHorizontalSides, got)
		}
	})

	t.Run("CalculateVerticalSides", func(t *testing.T) {
		// CalculateVerticalSides
		expectedVerticalSides := 2
		got := CalculateVerticalSides(gardenMap, group1)
		if got != expectedVerticalSides {
			t.Errorf("Expected %d vertical sides, got %d", expectedVerticalSides, got)
		}
	})

	t.Run("CalculateSides", func(t *testing.T) {
		// CalculateSize
		groupC := groups["C_1_2"]
		expected := 8
		got := CalculateSides(gardenMap, groupC)
		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})

}

func TestFindGroups2(t *testing.T) {
	inputStr := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	gardenMap := parseInput(inputStr)

	groups := FindGroups(gardenMap)

	if len(groups) != 11 {
		t.Errorf("Expected 4 groups, got %d", len(groups))
	}

	group1 := groups["I_0_4"]
	expectedArea, expectedPerimeter := 4, 8
	gotArea, gotPerimeter := CalculateRegionCost(gardenMap, group1)
	if gotArea != expectedArea || gotPerimeter != expectedPerimeter {
		t.Errorf("Expected area %d and perimeter %d, got %d and %d", expectedArea, expectedPerimeter, gotArea, gotPerimeter)
	}
}

func TestCalculatePricePart2(t *testing.T) {
	inputStr := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
	gardenMap := parseInput(inputStr)

	expected := 368
	got := CalculatePricePart2(gardenMap)
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestCalculatePricePart2_larger(t *testing.T) {
	inputStr := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	gardenMap := parseInput(inputStr)

	expected := 1206
	got := CalculatePricePart2(gardenMap)
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
