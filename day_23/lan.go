// https://adventofcode.com/2024/day/23

package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func part1(input string) {
	g := Graph{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		ids := strings.Split(line, "-")
		g.AddEdge(ids[0], ids[1])

	}
	triplets := g.FindTriplets()

	interesectingTriplets := [][]string{}
	for _, triplet := range triplets {

		for _, node := range triplet {
			if strings.HasPrefix(node, "t") {
				interesectingTriplets = append(interesectingTriplets, triplet)
				break
			}
		}
	}
	for i, triplet := range interesectingTriplets {
		fmt.Println(i, triplet)
	}
}

func part2(input string) {
	fmt.Println("Part 2")

	g := Graph{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		ids := strings.Split(line, "-")
		g.AddEdge(ids[0], ids[1])

	}

	largestComponent := g.FindLargestFullyConnnectedComponent()
	fmt.Println("Largest fully connected component", largestComponent.Size())

	sorted := largestComponent.Slice()
	slices.Sort(sorted)
	outputStr := strings.Join(sorted, ",")
	fmt.Println(outputStr)
}

func AddToConnectedCompoenents(stronglyConnected [][]string, newComponent []string) [][]string {
	found := false
	for _, component := range stronglyConnected {
		if slices.Equal(component, newComponent) {
			found = true
			break
		}
	}
	if !found {
		stronglyConnected = append(stronglyConnected, newComponent)
	}
	return stronglyConnected
}

func main() {
	data, error := io.ReadAll(os.Stdin)

	if error != nil {
		panic(error)
	}

	part1(string(data))

	part2(string(data))

}
