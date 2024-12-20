// https://adventofcode.com/2024/day/16
package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type FacingDirection int

const (
	east FacingDirection = iota
	north
	west
	south
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (d FacingDirection) rotationCost(newDirection FacingDirection) int {
	// They can also rotate clockwise or counterclockwise 90 degrees at a time (increasing their score by 1000 points).
	rotations := abs(int(newDirection) - int(d))
	switch rotations {
	case 0:
		return 0
	case 1, 3:
		return 1000
	case 2:
		return 2000

	}
	// panic("Invalid rotation")
	return 0
}

// Directions for moving up, down, left, right
var directions = [][2]int{
	{0, 1}, {1, 0},
	{0, -1}, {-1, 0},
}

type Location struct {
	row, col int
}

type Node struct {
	loc    Location
	symbol rune
}

func (n Node) String() string {
	return fmt.Sprintf("%v%c", n.loc, n.symbol)
}

type Graph struct {
	// adjacency list representation of the graph where each Node had a list of adjancent Nodes
	adj map[Node][]Node
}

func (g Graph) String() string {
	s := ""
	for node, neighbors := range g.adj {
		s += fmt.Sprintf("%s: ", node.String())
		for _, neighbor := range neighbors {
			s += fmt.Sprintf("%v ", neighbor)
		}
		s += "\n"
	}
	return s
}

func (g Graph) GetStartNode() Node {
	for node := range g.adj {
		if node.symbol == StartSymbol {
			return node
		}
	}
	panic("No start node found")
}

func (g Graph) GetEndNode() Node {
	for node := range g.adj {
		if node.symbol == EndSymbol {
			return node
		}
	}
	panic("No end node found")
}

func calculateNewFacingDirection(from Node, to Node) FacingDirection {
	dy, dx := to.loc.row-from.loc.row, to.loc.col-from.loc.col
	if dy == 0 {
		if dx > 0 {
			return east
		} else {
			return west
		}
	}
	if dx == 0 {
		if dy > 0 {
			return south
		} else {
			return north
		}
	}

	panic("Invalid move")
}

// calculate the cost of moving from 'from' node, facing 'facingDirection' to 'to' node
// returns the moving cost and the new facing direction
func GetMoveCost(from Node, facingDirection FacingDirection, to Node) (cost int, newFacingDirection FacingDirection) {

	newFacingDirection = calculateNewFacingDirection(from, to)
	cost = facingDirection.rotationCost(newFacingDirection) + 1

	return
}

func CalculatePathCost(path []Node) uint64 {
	// calculate the cost of the path
	// cost is the number of steps taken
	// if there is a
	currentNode := path[0]
	currentFacingDirection := east
	pathCost := uint64(0)
	for _, pathNode := range path[1:] {
		cost, newFacingDirection := GetMoveCost(currentNode, currentFacingDirection, pathNode)
		pathCost += uint64(cost)
		currentFacingDirection = newFacingDirection
		currentNode = pathNode
	}
	return pathCost
}

func (g Graph) GetAllPaths(start Node, end Node) [][]Node {
	// find all paths from start to end
	// use BFS to find all paths from start to end
	// use a queue to store the paths
	allPaths := make([][]Node, 0)

	queue := [][]Node{{start}}

	currPathLength := 0
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		if len(path) > currPathLength {
			currPathLength = len(path)
			fmt.Println("Current path length: ", currPathLength)
			fmt.Println("Queue length: ", len(queue))
		}
		last_node := path[len(path)-1]

		// find one path
		if last_node == end {
			fmt.Println("Found a path, with cost: ", CalculatePathCost(path))
			allPaths = append(allPaths, path)
		} else {
			for _, neighbor := range g.adj[last_node] {
				// check if the neighbor is already in the path
				// if so, skip it
				// if not, add it to the path and add the path to the queue
				if slices.Contains(path, neighbor) {
					continue
				}

				newPath := make([]Node, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	// return the paths
	return allPaths
}

const (
	EmptySymbol = '.'
	WallSymbol  = '#'
	StartSymbol = 'S'
	EndSymbol   = 'E'
)

func ParseInput(input string) Graph {
	// Parse the input string and construct the graph
	// where graph nodes are locations with the emptyymbol, startSymbol, endSymbol
	graph := Graph{make(map[Node][]Node)}
	grid := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	for y, row := range grid {
		for x, cell := range row {
			if cell == WallSymbol {
				continue
			}

			node := Node{Location{y, x}, cell}

			// check the neighbors of the current location
			neighbors := []Node{}
			for _, dir := range directions {
				ny, nx := y+dir[0], x+dir[1]
				if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[0]) && grid[ny][nx] != WallSymbol {
					neighbors = append(neighbors, Node{Location{ny, nx}, grid[ny][nx]})
				}
			}

			if _, ok := graph.adj[node]; !ok {
				graph.adj[node] = neighbors
			}
		}
	}

	// simply the graph by removing the ndos with just one neighbor, they can merged to very last node
	return graph
}

func (g Graph) FindMinCost(start Node, end Node) (minCost uint64, minPath []Node) {
	// use BFS to find the path with the minimum cost
	allPaths := g.GetAllPaths(start, end)
	minCost = 1<<63 - 1
	minPath = []Node{}
	for _, path := range allPaths {
		cost := CalculatePathCost(path)
		if cost < minCost {
			minCost = cost
			minPath = path
		}
	}
	return
}

type Path struct {
	nodes []Node
	cost  uint64
}

func (g Graph) Diijkstra(start Node, end Node) (minCost uint64, minPaths [][]Node) {
	// use Dijkstra's algorithm to find the path with the minimum cost
	// use a priority queue to store the paths
	// the priority is the cost of the path
	inf := uint64(1<<63 - 1)
	// saving the cost of reaching this Node from the start Node
	costAndPath := make(map[Node]Path)
	for node := range g.adj {
		costAndPath[node] = Path{[]Node{}, inf}
	}
	// initialize the cost of reaching the start Node to 0
	costAndPath[start] = Path{[]Node{start}, 0}

	// initialize the priority queue
	pq := []Path{costAndPath[start]}

	minCost = inf
	minPaths = [][]Node{}
	for len(pq) > 0 {
		// pop the path with the minimum cost
		path := pq[0]
		pq = pq[1:]

		// skip checking further if the cost is already higher than the minimum cost
		if minCost != inf && path.cost > minCost {
			break
		}

		node := path.nodes[len(path.nodes)-1]
		if path.cost <= costAndPath[node].cost {
			// find a shortest path to this node
			minPaths = append(minPaths, path.nodes)
			costAndPath[node] = path
		}

		// we reached exit
		if node == end {
			if minCost == inf {
				minCost = path.cost
				minPaths = append(minPaths, path.nodes)
			} else if path.cost == minCost {
				fmt.Println("saving extra path to: ", node)
				minPaths = append(minPaths, path.nodes)
			}
		}

		for _, neighbor := range g.adj[node] {

			// check if the neighbor is already in the path
			// if so, skip it
			// otherwise add it to the path and add the path to the queue
			if slices.Contains(path.nodes, neighbor) {
				continue
			}

			if neighbor == start {
				continue
			}

			newPath := make([]Node, len(path.nodes))
			copy(newPath, path.nodes)
			newPath = append(newPath, neighbor)
			newCost := CalculatePathCost(newPath)

			// neighbor had been calculated before, the new path is might be just one turn more then the previous path, and next move will make them equal
			if costAndPath[neighbor].cost != inf && newCost > 1000+costAndPath[neighbor].cost {
				continue
			}

			pq = append(pq, Path{newPath, newCost})

			// sort the priority queue
			slices.SortFunc(pq, func(a, b Path) int {
				return int(a.cost - b.cost)
			})
		}

	}

	// only return the ones that ends at 'EndSymbol'
	minPaths = slices.DeleteFunc(minPaths, func(path []Node) bool {
		return path[len(path)-1].symbol != EndSymbol
	})
	return
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	maze := ParseInput(string(data))
	fmt.Println(maze.String())

	// find the start and end nodes
	start := maze.GetStartNode()
	end := maze.GetEndNode()
	fmt.Println("Start node:", start)
	fmt.Println("End node:", end)
	// find the path with the minimum cost
	// minCost, minPath := maze.FindMinCost(start, end)
	minCost, minPaths := maze.Diijkstra(start, end)
	fmt.Println("Minimum cost:", minCost)
	// print the path
	bestSpots := map[Location]bool{}

	for i, path := range minPaths {
		fmt.Printf("Minimum path %d: %v \n", i, path)
		for _, node := range path {
			bestSpots[node.loc] = true
		}
	}

	fmt.Println("Number of best spots: ", len(bestSpots))

}
