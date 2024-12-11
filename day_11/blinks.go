// https://adventofcode.com/2024/day/11
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type MutationResult struct {
	stoneNumber     uint64
	newStoneNumbers []uint64
}

// this table saves the mutation result of any number seen so far
var mutationTable = map[uint64]MutationResult{}

func StoneMutation(number uint64) []uint64 {

	// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
	if number == 0 {
		return []uint64{1}
	}

	// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
	// The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone.
	// (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
	numberStr := fmt.Sprintf("%d", number)
	if len(numberStr)%2 == 0 {
		leftStone, _ := strconv.Atoi(numberStr[:len(numberStr)/2])
		ringStone, _ := strconv.Atoi(numberStr[len(numberStr)/2:])
		return []uint64{uint64(leftStone), uint64(ringStone)}
	}

	// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
	return []uint64{number * 2024}
}

// mutate a list of stones to a new list of stones
func MutateStones(stones []uint64) (result []uint64) {
	for _, stone := range stones {
		// test if the stone mutation is already in the table
		_, ok := mutationTable[stone]
		if !ok {
			mutationTable[stone] = MutationResult{stone, StoneMutation(stone)}
		}
		result = append(result, mutationTable[stone].newStoneNumbers...)
	}
	return
}

// input and out are maps of stone numbers and their count
func MutateStoneMap(stones map[uint64]uint64) (result map[uint64]uint64) {
	result = make(map[uint64]uint64)
	for stone, count := range stones {
		// the mutation and caching part
		if _, ok := mutationTable[stone]; !ok {
			mutationTable[stone] = MutationResult{stone, StoneMutation(stone)}
		}

		// the saving part
		for _, newStone := range mutationTable[stone].newStoneNumbers {
			if _, ok := result[newStone]; !ok {
				result[newStone] = count
			} else {
				result[newStone] += count
			}
		}
	}

	return
}

func SliceToMap(stones []uint64) (result map[uint64]uint64) {
	result = make(map[uint64]uint64)
	for _, stone := range stones {
		if _, ok := result[stone]; !ok {
			result[stone] = 1
		} else {
			result[stone] += 1
		}
	}
	return
}

func GetNumberOfStonesAfterMutation(stoneMap map[uint64]uint64, numberOfBlinks int) uint64 {
	for i := 1; i <= numberOfBlinks; i++ {
		stoneMap = MutateStoneMap(stoneMap)
	}

	stoneCounts := uint64(0)
	for _, count := range stoneMap {
		stoneCounts += count
	}

	return stoneCounts
}

func main() {
	input := "965842 9159 3372473 311 0 6 86213 48"
	stoneStrs := strings.Fields(input)
	stoneInts := make([]uint64, len(stoneStrs))
	for i, stoneStr := range stoneStrs {
		x, _ := strconv.Atoi(stoneStr)
		stoneInts[i] = uint64(x)
	}

	stoneIntMap := SliceToMap(stoneInts)

	numberOfBlinks := 75
	start := time.Now()
	numberOfStones := GetNumberOfStonesAfterMutation(stoneIntMap, numberOfBlinks)
	fmt.Println("after", numberOfBlinks, " blinks, stoneCounts:", numberOfStones, "Elapsed time:", time.Since(start))

}
