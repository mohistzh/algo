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
	return -1
}

func main() {
	fmt.Println(findSmallestMissingPositiveNumber([]int{-3, 1, 5, 4, 2}))
	fmt.Println(findSmallestMissingPositiveNumber([]int{3, -2, 0, 1, 2}))
	fmt.Println(findSmallestMissingPositiveNumber([]int{3, 2, 5, 1}))
}
