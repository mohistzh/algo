package main

import (
	"fmt"
)

/*
	Given an unsorted array containing numbers and a number ‘k’, find the first ‘k’ missing positive numbers in the array.
	e.g.
	Input: [3, -1, 4, 5, 5], k=3
	Output: [1, 2, 6]
	Explanation: The smallest missing positive numbers are 1, 2 and 6.
*/
func findFirstKMissingNumbers(input []int, k int) []int {
	i, n := 0, len(input)-1

	for i < n {
		j := input[i] - 1
		if input[i] > 0 && input[i] <= n && input[i] != input[j] {
			input[i], input[j] = input[j], input[i]
		} else {
			i++
		}
	}
	var missingNumbers []int
	// no build-in set structure in Golang....
	extraNumbers := make(map[int]int)
	for i := 0; i < n; i++ {
		if len(missingNumbers) < k {
			if input[i] != i+1 {
				missingNumbers = append(missingNumbers, i+1)
				extraNumbers[input[i]] = 1
			}
		}
	}
	i = 1
	for len(missingNumbers) < k {
		candiateNumber := i + n
		if _, ok := extraNumbers[candiateNumber]; !ok {
			missingNumbers = append(missingNumbers, candiateNumber)
		}
		i++
	}

	fmt.Println(missingNumbers)
	return missingNumbers
}

func main() {
	findFirstKMissingNumbers([]int{3, -1, 4, 5, 5}, 3)
	findFirstKMissingNumbers([]int{2, 3, 4}, 3)
	findFirstKMissingNumbers([]int{-2, -3, 4}, 2)
}
