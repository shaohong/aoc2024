// https://adventofcode.com/2024/day/22

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const pruner = uint(1<<24 - 1)

func nextSecret(currentSecret uint) uint {
	// Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
	currentSecret = pruner & ((currentSecret << 6) ^ currentSecret)

	// Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
	currentSecret = pruner & ((currentSecret >> 5) ^ currentSecret)

	// Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number
	return pruner & ((currentSecret << 11) ^ currentSecret)
}

func nthSecret(initialSecret uint, n int) uint {
	if n == 0 {
		return initialSecret
	} else {
		return nextSecret(nthSecret(initialSecret, n-1))
	}
}

type ChangeSequence struct {
	sequence [4]int
}

func main() {
	data, error := io.ReadAll(os.Stdin)

	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(data), "\n")
	sumOfSecrets := uint(0)
	for _, line := range lines {
		initialSecret := uint(0)
		fmt.Sscanf(line, "%d", &initialSecret)

		sumOfSecrets += nthSecret(initialSecret, 2000)

	}
	fmt.Println(sumOfSecrets)

	// part 2
	fmt.Println(strings.Repeat("-", 10), "part 2", strings.Repeat("-", 10))

	// generate buyer's price sequences
	priceSequences := make([][2001]uint, len(lines))

	for i, line := range lines {
		initialSecret := uint(0)
		fmt.Sscanf(line, "%d", &initialSecret)

		currentSecret := initialSecret
		priceSequences[i][0] = currentSecret % 10
		for j := 1; j <= 2000; j++ {
			currentSecret = nextSecret(currentSecret)
			priceSequences[i][j] = currentSecret % 10
		}
	}

	// here we save the sequence and first price, reached by that sequence, in buyer's price sequence
	sequeceAndFirstPrice := make(map[ChangeSequence][]uint)

	for i, priceSequence := range priceSequences {
		for j := 4; j < len(priceSequence); j++ {
			price := priceSequence[j]
			changes := ChangeSequence{}
			for k := 0; k < 4; k++ {
				changes.sequence[3-k] = int(priceSequence[j-k]) - int(priceSequence[j-k-1])
			}

			if _, ok := sequeceAndFirstPrice[changes]; !ok {
				sequeceAndFirstPrice[changes] = make([]uint, len(priceSequences))
			}
			// only save the first price for the sequence
			if sequeceAndFirstPrice[changes][i] == 0 {
				sequeceAndFirstPrice[changes][i] = price
			}
		}
	}

	// now find the sequence with the maximum total price (bananas)
	maxPrice := uint(0)
	bestSequence := ChangeSequence{}
	for k, v := range sequeceAndFirstPrice {
		// fmt.Printf("sequence: %v, first price: %v\n", k, v)
		totalPrice := uint(0)
		for _, price := range v {
			totalPrice += price
		}

		if totalPrice > maxPrice {
			maxPrice = totalPrice
			bestSequence = k
		}
	}

	fmt.Printf("sequence: %v, max price: %v\n", bestSequence, maxPrice)
}
