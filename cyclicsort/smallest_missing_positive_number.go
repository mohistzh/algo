package main

import (
	"fmt"
)

/*
	Given an unsorted array containing numbers, find the smallest missing positive number in it.
	e.g.
	Input: [-3, 1, 5, 4, 2]
	Output: 3
	Explanation: The smallest missing positive number is '3'
*/
func findSmallestMissingPositiveNumber(input []int) int {
	i, n := 0, len(input)
	for i < n {
		j := input[i] - 1
		if input[i] > 0 && input[i] <= n && input[i] != input[j] {
			input[i], input[j] = input[j], input[i]
		} else {
			i++
		}
	}
	for i := 0; i < n; i++ {
		if input[i] != i+1 { // find smallest positive missing number
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(findSmallestMissingPositiveNumber([]int{-3, 1, 5, 4, 2}))
	fmt.Println(findSmallestMissingPositiveNumber([]int{3, -2, 0, 1, 2}))
	fmt.Println(findSmallestMissingPositiveNumber([]int{3, 2, 5, 1}))
}
