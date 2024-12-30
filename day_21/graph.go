package main

import "fmt"

const WallSymbol = '#'

type Move int

const (
	left Move = iota
	right
	up
	down
	activate
)

func (m Move) String() string {
	return [...]string{"<", ">", "^", "v", "A"}[m]
}

type MoveSequence []Move

func (m MoveSequence) String() string {
	result := ""
	for _, move := range m {
		result += move.String()
	}
	return result
}

func DxDyToMove(dx, dy int) Move {
	switch dx {
	case 0:
		switch dy {
		case 1:
			return down
		case -1:
			return up
		}
	case 1:
		return right
	case -1:
		return left
	}

	panic(fmt.Sprintf("invalid dx, dy: %d, %d", dx, dy))
}

var directions = [][2]int{
	{0, 1}, {1, 0},
	{0, -1}, {-1, 0},
}

type Location struct {
	x int
	y int
}

type Node struct {
	loc    Location
	symbol rune
}

type Path []Node

func (n Node) String() string {
	return fmt.Sprintf("(%d, %d)%c", n.loc.x, n.loc.y, n.symbol)
}

type Graph struct {
	// adjacency list representation of the graph
	adj       map[Node][]Node
	pathCache map[Node]map[Node][]Path
}

func (g Graph) GetPathFromCache(start Node, end Node) []Path {
	if g.pathCache == nil {
		g.pathCache = make(map[Node]map[Node][]Path)
	}
	if g.pathCache[start] == nil {
		g.pathCache[start] = make(map[Node][]Path)
	}
	return g.pathCache[start][end]
}

func (g Graph) AddPathToCache(start Node, end Node, paths []Path) {
	if g.pathCache == nil {
		g.pathCache = make(map[Node]map[Node][]Path)
	}
	if g.pathCache[start] == nil {
		g.pathCache[start] = make(map[Node][]Path)
	}
	g.pathCache[start][end] = paths
}

func (g Graph) String() string {
	result := ""
	for node, neighbors := range g.adj {
		result += fmt.Sprintf("%s: ", node.String())
		for _, neighbor := range neighbors {
			result += fmt.Sprintf("%s ", neighbor.String())
		}
		result += "\n"
	}
	return result
}

// parse a 2d array of runes into a graph
func ParseToGraph(input [][]rune) Graph {
	adj := make(map[Node][]Node)
	for y, row := range input {
		for x, symbol := range row {
			if symbol == WallSymbol {
				continue
			}
			node := Node{Location{x, y}, symbol}
			adj[node] = make([]Node, 0)
			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]
				if ny >= 0 && ny < len(input) && nx >= 0 && nx < len(input[0]) {
					adj[node] = append(adj[node], Node{Location{nx, ny}, input[ny][nx]})
				}
			}
		}
	}

	return Graph{adj, nil}
}

func (g Graph) GetNode(symbol rune) Node {
	for node := range g.adj {
		if node.symbol == symbol {
			return node
		}
	}
	panic("No node found")
}

// check if a node is in a path
func NodeInPath(node Node, path Path) bool {
	for _, n := range path {
		if n.loc == node.loc {
			return true
		}
	}
	return false
}

// get the shortest paths (can be multiple) from start to end
func (g Graph) GetShortestPaths(start Node, end Node) []Path {
	cached := g.GetPathFromCache(start, end)
	if cached != nil {
		return cached
	}

	// breadth-first search
	queue := make([]Path, 0)
	queue = append(queue, Path{start})

	var shortestPaths []Path

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		if len(shortestPaths) > 0 && len(path) > len(shortestPaths[0]) {
			break
		}

		lastNode := path[len(path)-1]
		if lastNode == end {
			shortestPaths = append(shortestPaths, path)
			continue
		}

		for _, neighbor := range g.adj[lastNode] {
			if !NodeInPath(neighbor, path) {
				newPath := make(Path, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	// save the result to cache
	g.AddPathToCache(start, end, shortestPaths)
	return shortestPaths
}

func (p Path) Length() int {
	return len(p) - 1
}

func (p Path) String() string {
	result := ""
	for _, node := range p {
		result += fmt.Sprintf("%c", node.symbol)
	}
	return result
}

func (p Path) ToMoveSequence() MoveSequence {
	result := make(MoveSequence, 0)
	for i := 0; i < p.Length(); i++ {
		dx, dy := p[i+1].loc.x-p[i].loc.x, p[i+1].loc.y-p[i].loc.y
		result = append(result, DxDyToMove(dx, dy))
	}
	return result
}
