package main

import (
	"fmt"
	"math"
)

/*
Given an array of positive numbers and a positive number ‘S’,
find the length of the smallest contiguous subarray whose sum is greater than or equal to ‘S’. Return 0, if no such subarray exists.
*/

func SmallestSubArrayWithGivenSum(sum int, arr []int) int {
	windowSum, windowStart := 0, 0
	result := math.Inf(1)
	for i := 0; i < len(arr); i++ {
		windowSum += arr[i]
		for windowSum >= sum {
			result = math.Min(result, float64(i-windowStart+1))
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	if result == math.Inf(1) {
		return -1
	}
	return int(result)
}
func SmallestSubArrayWithGivenSum2(sum int, arr []int) int {
	result, count, carry := math.Inf(1), 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= sum {
			return 1
		}
		carry += arr[i]
		count += 1
		if carry >= sum {
			result = math.Min(float64(count), float64(result))
			count = 1
			carry = arr[i]
		}
	}
	return int(result)
}

func main() {
	fmt.Println("#1: %d", SmallestSubArrayWithGivenSum(7, []int{2, 1, 5, 2, 3, 2}))
	fmt.Println("#1: %d", SmallestSubArrayWithGivenSum(7, []int{2, 1, 5, 2, 8}))
	fmt.Println("#1: %d", SmallestSubArrayWithGivenSum(8, []int{3, 4, 1, 1, 6}))

	fmt.Println("#2: %d", SmallestSubArrayWithGivenSum2(7, []int{2, 1, 5, 2, 3, 2}))
	fmt.Println("#2: %d", SmallestSubArrayWithGivenSum2(7, []int{2, 1, 5, 2, 8}))
	fmt.Println("#2: %d", SmallestSubArrayWithGivenSum2(8, []int{3, 4, 1, 1, 6}))
}
