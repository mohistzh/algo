package main

import (
	"fmt"
)

/*
	Given an array of sorted numbers and a target sum, find a pair in the array whose sum is equal to the given target.
*/
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

// hashmap approach
func findPairWithTargetSumUsingHashMap(input []int, targetSum int) []int {
	indicesMap := make(map[int]int)
	result := []int{-1, -1}
	for i, num := range input {
		if _, ok := indicesMap[targetSum-num]; ok {
			result[0] = indicesMap[targetSum-num]
			result[1] = i
			break
		} else {
			indicesMap[num] = i
		}
	}
	return result
}

func main() {
	fmt.Println(findPairWithTargetSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(findPairWithTargetSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(findPairWithTargetSumUsingHashMap([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(findPairWithTargetSumUsingHashMap([]int{2, 5, 9, 11}, 11))
}
