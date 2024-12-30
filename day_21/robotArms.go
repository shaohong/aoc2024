package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

var numericPad = [][]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'#', '0', 'A'},
}

var directionalPad = [][]rune{
	{'#', '^', 'A'},
	{'<', 'v', '>'},
}

var numPadGraph = ParseToGraph(numericPad)
var dirPadGraph = ParseToGraph(directionalPad)

type QueueElement struct {
	previousMoves MoveSequence
	remainingCode string
}

func GetNumberFromCode(code string) int {
	number, error := strconv.Atoi(string(code[:len(code)-1]))
	if error != nil {
		panic(error)
	}
	return number
}

func GetMoveSequence(code string, padGraph Graph) []MoveSequence {
	// robotic arm is aiming at 'A'
	// find the move sequence for the given code on directional pad
	allPossibleSequences := make([]MoveSequence, 0)
	queue := []QueueElement{}
	queue = append(queue, QueueElement{MoveSequence{}, code})

	// do a BFS to find all possible paths for the given code
	// each time we check one path segment (the first two characters in the code)
	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]

		if len(task.remainingCode) == 1 {
			// we have reached the end of the path
			allPossibleSequences = append(allPossibleSequences, task.previousMoves)
			continue
		}

		startNode := padGraph.GetNode(rune(task.remainingCode[0]))
		endNode := padGraph.GetNode(rune(task.remainingCode[1]))
		pathsInThisStep := padGraph.GetShortestPaths(startNode, endNode)

		// merge the previous path with the new paths via multiplication principle
		for _, path := range pathsInThisStep {
			// copy previous moves and then append newly found moves
			newSequence := make(MoveSequence, len(task.previousMoves))
			copy(newSequence, task.previousMoves)

			newSequence = append(newSequence, path.ToMoveSequence()...)
			// pressing the 'activate' button to simulate pressing the code completed in this step
			newSequence = append(newSequence, activate)
			queue = append(queue, QueueElement{newSequence, task.remainingCode[1:]})
		}

	}

	return allPossibleSequences
}

func GetShortestMoveSequenceLength(codePath string, numDirPads int) uint64 {
	// it's much easier to use python's itertools.product, itertools.permutation, ... to generate all possible move sequences

	// get robot 1 move sequences (to press the codes)
	robotArmMoveSequences := GetMoveSequence(codePath, numPadGraph)

	robotArmMoveSequencesStr := make([]string, len(robotArmMoveSequences))
	for i, sequence := range robotArmMoveSequences {
		robotArmMoveSequencesStr[i] = sequence.String()
	}

	for i := 0; i < numDirPads; i++ {
		newSequence := make([]MoveSequence, 0)
		for _, code := range robotArmMoveSequencesStr {
			newSequence = append(newSequence, GetMoveSequence("A"+code, dirPadGraph)...)
		}
		robotArmMoveSequencesStr = make([]string, len(newSequence))
		for i, sequence := range newSequence {
			robotArmMoveSequencesStr[i] = sequence.String()
		}
	}

	shortestSequence := ""
	longestSequence := ""
	// find the shortest move sequence
	for _, sequence := range robotArmMoveSequencesStr {
		if len(shortestSequence) == 0 || len(sequence) < len(shortestSequence) {
			shortestSequence = sequence
		}
		if len(longestSequence) == 0 || len(sequence) > len(longestSequence) {
			longestSequence = sequence
		}

	}

	fmt.Println("Shortest sequence: ", shortestSequence, " length: ", len(shortestSequence))
	fmt.Println("Longest sequence: ", longestSequence, " length: ", len(longestSequence))
	return uint64(len(shortestSequence))
}

// numDirPads is the number of intermediate directional pads
func GetCodeComplexity(code string, numDirPads int) uint64 {

	shortestSequenceLength := GetShortestMoveSequenceLength("A"+code, numDirPads)
	return uint64(GetNumberFromCode(code)) * shortestSequenceLength
}

var MoveCostCache = make(map[string]int64)

func GetMoveCostCacheKey(from rune, to rune, levels int) string {
	return string(from) + "-" + string(to) + "-" + strconv.Itoa(levels)
}

func GetMoveCostFromCache(from rune, to rune, levels int) int64 {
	key := GetMoveCostCacheKey(from, to, levels)
	if cost, ok := MoveCostCache[key]; ok {
		return cost
	}
	return -1
}

func SetMoveCostToCache(from rune, to rune, levels int, cost int64) {
	key := GetMoveCostCacheKey(from, to, levels)
	MoveCostCache[key] = cost
}

func DumpMoveCostCache() {
	fmt.Println("MoveCostCache:")
	for key, value := range MoveCostCache {
		fmt.Println(key, value)
	}
}

var numericPadChars = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func GetMoveCost(from rune, to rune, levels int) int64 {
	// check if the cost is already in the cache
	cost := GetMoveCostFromCache(from, to, levels)
	if cost != -1 {
		return int64(cost)
	}

	useDirectionalPad := true
	if slices.Contains(numericPadChars, from) || slices.Contains(numericPadChars, to) {
		useDirectionalPad = false
	}
	var paths []Path
	if useDirectionalPad {
		paths = dirPadGraph.GetShortestPaths(dirPadGraph.GetNode(from), dirPadGraph.GetNode(to))
	} else {
		paths = numPadGraph.GetShortestPaths(numPadGraph.GetNode(from), numPadGraph.GetNode(to))
	}

	if levels == 0 {
		// no intermediate directional pad
		fmt.Println("level", levels, "Move from ", string(from), " to ", string(to), " cost: ", len(paths[0]))
		SetMoveCostToCache(from, to, levels, int64(len(paths[0])))
		return int64(len(paths[0]))
	}

	// recursively calculate the cost of moving from 'from' to 'to' with 'levels-1' intermediate directional pads
	minCost := int64(1 << 60)
	for _, path := range paths {
		// start to use the intermediate robot and directional pad
		sequenceStr := "A" + path.ToMoveSequence().String() + "A"
		cost := int64(0)
		for i := 0; i < len(sequenceStr)-1; i++ {
			cost += GetMoveCost(rune(sequenceStr[i]), rune(sequenceStr[i+1]), levels-1)
		}
		if cost < minCost {
			minCost = cost
		}
	}

	// save the cost to cache
	SetMoveCostToCache(from, to, levels, minCost)
	return minCost

}

func GetCodeCost(numericCode string, levels int) int64 {

	codePath := "A" + numericCode
	cost := int64(0)
	for i := 0; i < len(codePath)-1; i++ {
		cost += GetMoveCost(rune(codePath[i]), rune(codePath[i+1]), levels)
	}

	return cost
}

func ParseInput(input string) []string {
	return strings.Split(input, "\n")
}

func main() {
	data, _ := io.ReadAll(os.Stdin)

	codes := ParseInput(string(data))

	totalCodeComplexity := uint64(0)
	numDirPads := 2
	for _, code := range codes {
		fmt.Println("Code: ", code)
		// totalCodeComplexity += GetCodeComplexity(code, numDirPads)
		cost := GetCodeCost(code, numDirPads)
		totalCodeComplexity += uint64(cost) * uint64(GetNumberFromCode(code))

	}
	fmt.Println("Total code complexity: ", totalCodeComplexity)

	fmt.Println("--- part 2 ---")
	numDirPads = 25
	complexity := uint64(0)
	for _, code := range codes {
		fmt.Println("Code: ", code)
		cost := GetCodeCost(code, numDirPads)
		fmt.Println("Code cost: ", cost)
		complexity += uint64(cost) * uint64(GetNumberFromCode(code))
	}
	fmt.Println("Total code complexity: ", complexity)
}
