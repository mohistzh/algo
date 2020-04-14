package main

import (
	"fmt"
)

/*
	Given an array of sorted numbers and a target sum, find a pair in the array whose sum is equal to the given target.
*/
// because it's a two pointer practicing case, we don't use a hashmap approach
func findPairWithTargetSum(input []int, targetSum int) []int {
	result := []int{-1, -1}
	left, right := 0, len(input)-1
	for left < right {
		sum := input[left] + input[right]
		if sum == targetSum {
			result[0] = left
			result[1] = right
			break
		} else if sum > targetSum {
			right--
		} else {
			left++
		}
	}
	return result
}

func main() {
	fmt.Println(findPairWithTargetSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(findPairWithTargetSum([]int{2, 5, 9, 11}, 11))
}
