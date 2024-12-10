// https://adventofcode.com/2024/day/10
package main

import (
	"container/list"
	"fmt"
	"slices"
)

type Location struct {
	x, y int
}

type Node struct {
	loc    Location
	height uint
}

type Graph struct {
	// adjacency list representation of the graph
	adj map[Node][]Node
}

func (g Graph) String() string {
	s := ""
	for node, neighbors := range g.adj {
		s += fmt.Sprintf("%v: ", node)
		for _, neighbor := range neighbors {
			s += fmt.Sprintf("%v ", neighbor)
		}
		s += "\n"
	}
	return s
}

func readInput() [][]uint {
	// read the input into a 2D array
	topomap := make([][]uint, 0)
	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		row := make([]uint, len(line))
		for i, c := range line {
			row[i] = uint(c - '0')
		}
		topomap = append(topomap, row)
	}
	return topomap
}

func MapToGraph(topomap [][]uint) Graph {
	// convert the topomap into a graph
	graph := Graph{make(map[Node][]Node)}
	for i, row := range topomap {
		for j, height := range row {
			node := Node{Location{i, j}, height}
			graph.adj[node] = make([]Node, 0)

			// posible edge is to a neighbor whose height is exactly 1 bigger than the current node
			// add the adjacent nodes

			if i > 0 { // moving up is possible
				if topomap[i-1][j] == height+1 {
					graph.adj[node] = append(graph.adj[node], Node{Location{i - 1, j}, topomap[i-1][j]})
				}
			}
			if i < len(topomap)-1 { // moving down
				if topomap[i+1][j] == height+1 {
					graph.adj[node] = append(graph.adj[node], Node{Location{i + 1, j}, topomap[i+1][j]})
				}
			}
			if j > 0 { // moving left
				if topomap[i][j-1] == height+1 {
					graph.adj[node] = append(graph.adj[node], Node{Location{i, j - 1}, topomap[i][j-1]})
				}
			}
			if j < len(row)-1 { // moving right
				if topomap[i][j+1] == height+1 {
					graph.adj[node] = append(graph.adj[node], Node{Location{i, j + 1}, topomap[i][j+1]})
				}
			}
		}
	}
	return graph
}

// do a BFS from the start node
// return top nodes (height of 9) visited
func (g Graph) BfsToTop(start Node) (topNodes []Location) {
	topNodes = make([]Location, 0)
	// breadth first search from the start node
	visited := make(map[Node]bool)
	queue := list.New()
	queue.PushBack(start)
	for queue.Len() > 0 {
		// pop the front node from the queue, visit it, then add its neighbors to the queue
		node := queue.Front().Value.(Node)
		queue.Remove(queue.Front())
		visited[node] = true

		if node.height == 9 {
			if !slices.Contains(topNodes, node.loc) {
				topNodes = append(topNodes, node.loc)
			}
		} else {
			for _, neighbor := range g.adj[node] {
				if !visited[neighbor] {
					queue.PushBack(neighbor)
				}
			}
		}
	}
	return topNodes
}

func (g Graph) CountPathstoTop(start Node) int {
	numberOfPathsToTop := 0

	visited := make(map[Node]bool)
	queue := list.New()
	queue.PushBack(start)
	for queue.Len() > 0 {
		// pop the front node from the queue, visit it, then add its neighbors to the queue
		node := queue.Front().Value.(Node)
		queue.Remove(queue.Front())
		visited[node] = true

		if node.height == 9 {
			numberOfPathsToTop += 1
		} else {
			for _, neighbor := range g.adj[node] {
				if !visited[neighbor] {
					queue.PushBack(neighbor)
				}
			}
		}
	}
	return numberOfPathsToTop
}

func main() {
	fmt.Println("Day 10: Hiking Trail")

	// read the input into a 2D array
	topomap := readInput()

	// convert the topomap into a graph
	graph := MapToGraph(topomap)

	fmt.Println("topomap: ", graph.String())

	startNode := Node{Location{0, 2}, topomap[0][2]}
	topNodes := graph.BfsToTop(startNode)
	// print the top nodes
	fmt.Println("start from", startNode, ", reachable tops:", topNodes)

	// find all the trail heads, i.e. the nodes with height 0
	trailHeads := make([]Node, 0)
	// find the trail heads
	for node, _ := range graph.adj {
		if node.height == 0 {
			trailHeads = append(trailHeads, node)
		}
	}

	totalScores := 0
	for _, trailHead := range trailHeads {
		topNodes := graph.BfsToTop(trailHead)
		totalScores += len(topNodes)
	}
	fmt.Println("Number of paths to the top nodes:", totalScores)

	fmt.Println("----- Part 2 -----")
	totalRatings := 0
	for _, trailHead := range trailHeads {
		numPathsToTop := graph.CountPathstoTop(trailHead)
		totalRatings += numPathsToTop
	}
	fmt.Println("Total number of paths to the top nodes:", totalRatings)
}
