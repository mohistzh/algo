package main

import (
	"fmt"
	"sort"
)

//	Given an array of unsorted numbers, find all unique triplets in it that add up to zero.
//	e.g.
//	Input: [-3, 0, 1, 2, -1, 1, -2]
//	Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
//	Explanation: There are four unique triplets whose sum is equal to zero.
func FindTriplets(input []int) [][]int {
	var triplets [][]int
	sort.Ints(input)
	for i := 0; i < len(input); i++ {
		// skip same element to avoid duplicate triplets
		if i > 0 && input[i] == input[i-1] {
			continue
		}
		searchTriplets(input, -input[i], i+1, &triplets)
	}
	return triplets
}

func searchTriplets(input []int, targetSum int, left int, triplets *[][]int) {
	right := len(input) - 1
	for left < right {
		leftVal := input[left]
		rightVal := input[right]
		currentSum := leftVal + rightVal
		if currentSum == targetSum { // found the triplet
			*triplets = append(*triplets, []int{-targetSum, leftVal, rightVal})
			left++
			right--
			for left < right && input[left] == input[left-1] {
				left++ // skip same element to avoid duplicate triplets
			}
			for left < right && input[right] == input[right+1] {
				right-- // skip same element to avoid duplicate triplets
			}
		} else if targetSum > currentSum {
			left++ // we need a pair with a bigger sum
		} else {
			right-- // we need a pair with a smaller sum
		}
	}
}

func main() {
	fmt.Println("Unique triplets in it that add up to zero:", FindTriplets([]int{-3, 0, 1, 2, -1, 1, -2}))
}
