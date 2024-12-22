package main

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func program(A uint64) (outputs []int) {
	for ok := true; ok; ok = (A != 0) {
		B := (A % 8) ^ 3
		C := A >> B // so C is A shifted right by 0 ~ 7 bits, the relevant bits in A in this whole computation is the last 10 binary digits
		A = A / 8   // A >> 3

		B = B ^ 5
		B = B ^ C //(only last 3 binary digits matters)

		outputs = append(outputs, int(B%8))
	}

	return outputs
}

func last10DigitsOutput() map[int][]string {
	// result is a mapping of the first output to the list of As which can generates that output

	result := make(map[int][]string)
	// the last 10 digits of A determines the first output number
	for a := 0b0000000000; a < 0b10000000000; a++ {
		outputs := program(uint64(a))
		if _, ok := result[int(outputs[0])]; !ok {
			result[int(outputs[0])] = []string{}
		}
		result[int(outputs[0])] = append(result[int(outputs[0])], fmt.Sprintf("%010b", a))
	}

	for _, candidates := range result {
		slices.Sort(candidates)
	}
	return result
}

func solveForA(outputs []int) (A uint64) {

	fmt.Println("solve for A with outputs: ", outputs)

	outputToCandidates := last10DigitsOutput()

	candidateStrings := make([]string, 0)
	for i := 0; i < len(outputs); i++ {
		fmt.Printf("checking output[%d]: %d\n", i, outputs[i])
		if len(candidateStrings) == 0 {
			// this should only happen when hendling the first output
			candidateStrings = append(candidateStrings, outputToCandidates[int(outputs[i])]...)
			continue
		}

		currentCandidates := outputToCandidates[int(outputs[i])]

		newCandidates := make([]string, 0)
		// new candidates are the intersection of the current candidates and previous candidates
		// the intersection is done by checking the last 7 digits the current candidate with the first 7 digits  of previous candidates
		for _, candidate := range currentCandidates {
			for _, prevCandidate := range candidateStrings {
				if strings.HasPrefix(prevCandidate, candidate[3:]) {
					newCandidates = append(newCandidates, fmt.Sprintf("%s%s", candidate[0:3], prevCandidate))
				}
			}
		}

		candidateStrings = newCandidates
		fmt.Printf("New we have %d candidates of length %d\n", len(candidateStrings), len(candidateStrings[0]))

	}

	// sort the candidates
	slices.Sort(candidateStrings)

	for _, candidate := range candidateStrings {
		fmt.Println(candidate)
		// convert binary string to decimal
		fmt.Sscanf(candidate, "%b", &A)
		currentOutputs := program(A)
		if reflect.DeepEqual(currentOutputs, outputs) {
			fmt.Printf("A: %d\n", A)
			return A
		}
	}

	panic("no solution found")
}

// func main() {
// 	fmt.Println("Day 17 Part 2")

// 	outputs := make([]uint64, 0)
// 	for _, i := range strings.Split("2,4,1,3,7,5,0,3,1,5,4,4,5,5,3,0", ",") {
// 		outputs = append(outputs, uint64(i[0]-'0'))
// 	}
// 	fmt.Println(outputs)

// 	A := solveForA(outputs)
// 	fmt.Printf("solution for A: %d\n", A)
// }
