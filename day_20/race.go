// https://adventofcode.com/2024/day/20

// there are some duplications from the day_16 for the data structure

package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Location struct {
	y int
	x int
}

// Directions for moving up, down, left, right
var directions = [][2]int{
	{0, 1}, {1, 0},
	{0, -1}, {-1, 0},
}

type Node struct {
	loc    Location
	symbol rune
}

func (n Node) String() string {
	return fmt.Sprintf("%v%c", n.loc, n.symbol)
}

type Path []Node

type Grid struct {
	nodes     [][]rune
	startNode Node
	endNode   Node
}

const (
	EmptySymbol = '.'
	WallSymbol  = '#'
	StartSymbol = 'S'
	EndSymbol   = 'E'
)

func (g Grid) GetShortestPath() Path {
	start := g.startNode
	end := g.endNode

	result := Path{}

	queue := []Path{{start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		lastNode := path[len(path)-1]
		if lastNode == end {
			return path
		}

		for _, dir := range directions {
			ny, nx := lastNode.loc.y+dir[0], lastNode.loc.x+dir[1]
			if ny > 0 && ny < len(g.nodes)-1 && nx > 0 && nx < len(g.nodes[0])-1 {
				neighbor := Node{Location{ny, nx}, g.nodes[ny][nx]}
				if neighbor.symbol != WallSymbol && !slices.Contains(path, neighbor) {

					newPath := make([]Node, len(path))
					copy(newPath, path)
					newPath = append(newPath, neighbor)
					queue = append(queue, newPath)

				}
			}
		}
	}

	return result
}

func (g Grid) Clone() Grid {
	nodes := make([][]rune, len(g.nodes))
	for i, row := range g.nodes {
		nodes[i] = make([]rune, len(row))
		copy(nodes[i], row)
	}

	return Grid{nodes, g.startNode, g.endNode}
}

func ParseInput(input string) Grid {
	grid := Grid{nodes: [][]rune{}}
	for _, line := range strings.Split(input, "\n") {
		grid.nodes = append(grid.nodes, []rune(line))
	}

	for y, row := range grid.nodes {
		for x, cell := range row {
			if cell == StartSymbol {
				grid.startNode = Node{Location{y, x}, StartSymbol}
				continue
			}

			if cell == EndSymbol {
				grid.endNode = Node{Location{y, x}, EndSymbol}
				continue
			}
		}
	}
	return grid
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(a, b Location) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func FindShortCuts(track Path, cheatLength int) map[int]int {
	// shortCuts saves the  timesavings and count
	shortCuts := make(map[int]int)

	// see if any node on the track is reachable via a cheating short cut
	for i := 0; i < len(track)-cheatLength; i++ {
		for j := i + 1; j < len(track); j++ {
			distance := manhattanDistance(track[i].loc, track[j].loc)
			if distance <= cheatLength {
				// node j on the track is reachable via a cheating short cut
				saving := j - i - distance
				if saving > 0 {
					shortCuts[saving] = shortCuts[saving] + 1
				}
			}
		}
	}

	return shortCuts
}

func main() {
	// Read the input
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	grid := ParseInput(string(data))

	track := grid.GetShortestPath()
	fmt.Println("lengh of track: ", len(track))
	// fmt.Println("track: ", track)
	fmt.Println("time to reach the end: ", len(track)-1)

	cheatLength := 2
	shortCuts := FindShortCuts(track, cheatLength)
	fmt.Println("shortCuts: ", shortCuts)

	goodShortcuts := 0
	goodShortcutThreshold := 100
	for timeSaving, count := range shortCuts {
		if timeSaving >= goodShortcutThreshold {
			goodShortcuts += count

		}
	}

	fmt.Println("# of goodShortcuts: ", goodShortcuts)

	fmt.Println("--- part 2 ---")
	cheatLength = 20
	shortCuts = FindShortCuts(track, cheatLength)
	goodShortcuts = 0
	goodShortcutThreshold = 100
	for timeSaving, count := range shortCuts {
		if timeSaving >= goodShortcutThreshold {
			goodShortcuts += count

		}
	}
	fmt.Println("# of goodShortcuts: ", goodShortcuts)

}
