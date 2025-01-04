// https://adventofcode.com/2024/day/24

package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type NodeType int

const (
	input NodeType = iota
	compute
)

type Node struct {
	id        string
	nodeType  NodeType
	value     int
	operation string    // AND, OR XOR
	inputs    [2]string // 2 inputs node ids
}

func (n *Node) String() string {
	if n.nodeType == input {
		return fmt.Sprintf("%s = %d", n.id, n.value)
	}
	return fmt.Sprintf("%s = %s %s %s, value = %d", n.id, n.inputs[0], n.operation, n.inputs[1], n.value)
}

type Circuit struct {
	nodes map[string]*Node
}

func (c *Circuit) String() string {
	result := ""
	for _, node := range c.nodes {
		result += node.String() + "\n"
	}
	return result
}

func (c *Circuit) Compute(nodeID string) int {
	node := c.nodes[nodeID]
	if node.nodeType == input {
		return node.value
	}

	if node.nodeType == compute && node.value != -1 {
		return node.value
	}

	input1 := c.Compute(node.inputs[0])
	input2 := c.Compute(node.inputs[1])

	switch node.operation {
	case "AND":
		node.value = input1 & input2
	case "OR":
		node.value = input1 | input2
	case "XOR":
		node.value = input1 ^ input2
	default:
		panic("Invalid operation " + node.operation)
	}

	return node.value
}

func (c *Circuit) ComputeAll() {
	for _, node := range c.nodes {
		c.Compute(node.id)
	}
}

func (c *Circuit) GetOutput() int {
	c.ComputeAll()

	// get all outputs from nodes names starting with 'z'
	outputNodeValues := make(map[string]int)
	outputNodeIDs := make([]string, 0)
	for _, node := range c.nodes {
		if strings.HasPrefix(node.id, "z") {
			outputNodeIDs = append(outputNodeIDs, node.id)
			outputNodeValues[node.id] = node.value
		}
	}

	slices.Sort(outputNodeIDs)
	outputAsStr := ""
	for i := len(outputNodeIDs) - 1; i >= 0; i-- {
		outputAsStr += fmt.Sprintf("%d", outputNodeValues[outputNodeIDs[i]])
	}

	result, _ := strconv.ParseInt(outputAsStr, 2, 0)
	return int(result)
}

func part2(c *Circuit) []string {
	// reference https://www.bytesizego.com/blog/aoc-day24-golang

	// 1. If the output of a gate is z, then the operation has to be XOR unless it is the last bit.
	// 2. If the output of a gate is not z and the inputs are not x, y then it has to be AND / OR, but not XOR.
	// 3. If you have a XOR gate with inputs x, y, there must be another XOR gate with this gate as an input.
	// Search through all gates for an XOR-gate with this gate as an input; if it does not exist, your (original) XOR gate is faulty.
	// 4. Similarly, if you have an AND-gate, there must be an OR-gate with this gate as an input. If that gate doesn't exist, the original AND gate is faulty., unless it's the first sum (x00 + y00)

	faltyGates := []string{}

	for _, node := range c.nodes {

		// 1. If the output of a gate is z, then the operation has to be XOR unless it is the last bit.
		if strings.HasPrefix(node.id, "z") {
			if node.operation != "XOR" {
				if node.id != "z45" {
					fmt.Printf("Node %s is faulty, per rule 1\n", node.id)
					faltyGates = append(faltyGates, node.id)
				}
			}
			continue
		}

		// 2. If the output of a gate is not z and the inputs are not x, y then it has to be AND / OR, but not XOR.
		if !strings.HasPrefix(node.inputs[0], "x") && !strings.HasPrefix(node.inputs[1], "x") && !strings.HasPrefix(node.inputs[0], "y") && !strings.HasPrefix(node.inputs[1], "y") {
			if !strings.HasPrefix(node.id, "z") {
				if node.operation == "XOR" {
					fmt.Printf("Node %s is faulty, per rule 2\n", node.id)
					faltyGates = append(faltyGates, node.id)
				}
			}
			continue
		}

		// 3. If you have a XOR gate with inputs x, y, there must be another XOR gate with this gate as an input.
		if node.operation == "XOR" && (strings.HasPrefix(node.inputs[0], "x") || strings.HasPrefix(node.inputs[1], "x") || strings.HasPrefix(node.inputs[0], "y") || strings.HasPrefix(node.inputs[1], "y")) {
			found := false
			for _, otherNode := range c.nodes {
				if otherNode.operation == "XOR" && (otherNode.inputs[0] == node.id || otherNode.inputs[1] == node.id) {
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("Node %s is faulty, per rule 3\n", node.id)
				faltyGates = append(faltyGates, node.id)
			}
			continue
		}

		// 4. Similarly, if you have an AND-gate (of x and y), there must be an OR-gate with this gate as an input. If that gate doesn't exist, the original AND gate is faulty.
		// only exception is x00 + y00
		if node.operation == "AND" && (strings.HasPrefix(node.inputs[0], "x") || strings.HasPrefix(node.inputs[1], "x") || strings.HasPrefix(node.inputs[0], "y") || strings.HasPrefix(node.inputs[1], "y")) {
			if node.inputs[0] == "x00" && node.inputs[1] == "y00" {
				continue
			}

			found := false
			for _, otherNode := range c.nodes {
				if otherNode.operation == "OR" && (otherNode.inputs[0] == node.id || otherNode.inputs[1] == node.id) {
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("Node %s is faulty, per rule 4\n", node.id)
				faltyGates = append(faltyGates, node.id)
			}
		}

	}

	return faltyGates
}

func ParseInput(data string) Circuit {
	sections := strings.Split(data, "\n\n")

	inputNodesData := strings.Split(sections[0], "\n")

	computeNodesData := strings.Split(sections[1], "\n")

	circuit := Circuit{nodes: make(map[string]*Node)}

	for _, nodeData := range inputNodesData {
		parts := strings.Split(nodeData, ":")
		nodeID := strings.TrimSpace(parts[0])
		nodeValue, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		circuit.nodes[nodeID] = &Node{id: nodeID, nodeType: input, value: nodeValue}
	}

	for _, nodeData := range computeNodesData {
		parts := strings.Split(nodeData, " ")
		inputNode1 := strings.TrimSpace(parts[0])
		operator := strings.TrimSpace(parts[1])
		inputNode2 := strings.TrimSpace(parts[2])
		nodeID := strings.TrimSpace(parts[4])

		circuit.nodes[nodeID] = &Node{id: nodeID, value: -1, nodeType: compute, operation: operator, inputs: [2]string{inputNode1, inputNode2}}
	}

	return circuit
}

func main() {
	fmt.Println("Day 24: Crossed Wires")

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	circuit := ParseInput(string(data))
	fmt.Println(circuit.String())

	fmt.Println("Output of the circuit:", circuit.GetOutput())

	fmt.Println(strings.Repeat("=", 20), "part 2", strings.Repeat("=", 20))
	circuit = ParseInput(string(data))
	faultyGates := part2(&circuit)
	slices.Sort(faultyGates)
	fmt.Println(strings.Join(faultyGates, ","))
}
