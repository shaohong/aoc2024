package main

import (
	"slices"

	"github.com/hashicorp/go-set"
)

type Graph struct {
	adjList map[string][]string
}

func (g *Graph) AddEdge(node1, node2 string) {
	if g.adjList == nil {
		g.adjList = make(map[string][]string)
	}
	g.adjList[node1] = append(g.adjList[node1], node2)
	g.adjList[node2] = append(g.adjList[node2], node1)
}

func (g *Graph) String() string {
	s := ""
	for k, v := range g.adjList {
		s += k + " -> "
		for _, val := range v {
			s += val + " "
		}
		s += "\n"
	}
	return s
}

func (g *Graph) FindTriplets() [][]string {
	strongComponents := [][]string{}

	for k, v := range g.adjList {
		// find strongly connected triplets for node k
		// for each node, n, in the adjacency list of node k, see if adj[n] contains k
		for _, n := range v {
			for _, m := range g.adjList[n] {
				if m == k {
					continue
				}
				if slices.Contains(g.adjList[m], k) { // k->n->m-k
					newComponent := []string{k, n, m}
					slices.Sort(newComponent)

					// check if the new component is already in the list
					found := false
					for _, component := range strongComponents {
						if slices.Equal(component, newComponent) {
							found = true
							break
						}
					}
					if !found {
						strongComponents = append(strongComponents, newComponent)
					}
				}
			}
		}
	}

	return strongComponents
}

func (g *Graph) FindLargestFullyConnnectedComponent() *set.Set[string] {
	visited := set.New[string](0)
	largestComponent := set.New[string](0)
	for k := range g.adjList {
		if visited.Contains(k) {
			continue
		}

		// seed the current component with node k
		currentComponent := set.From([]string{k})
		visited.Insert(k)

		// check all the nodes connected to k
		for _, n := range g.adjList[k] {
			// if node n is connected with every component in currentComponent, add n to the current component
			fullyConnected := true
			for _, c := range currentComponent.Slice() {
				if !slices.Contains(g.adjList[n], c) {
					fullyConnected = false
					break
				}
			}
			if fullyConnected {
				currentComponent.Insert(n)
				visited.Insert(n)
			}
		}

		if currentComponent.Size() > largestComponent.Size() {
			largestComponent = currentComponent
		}

	}

	return largestComponent
}
