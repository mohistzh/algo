package main

import (
	"fmt"
	"math"
)

/*
	Given an array of positive numbers and a positive number ‘k’,
	find the maximum sum of any contiguous subarray of size ‘k’.
*/

func MaximumSubArrayOfSizeK(k int, arr []int) int {
	if len(arr) < k {
		return -1
	}
	res, windowSum, windowStart := 0.0, 0, 0
	for i := 0; i < len(arr); i++ {
		windowSum += arr[i]
		if i >= k-1 {
			res = math.Max(float64(res), float64(windowSum))
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	return int(res)
}

func main() {
	fmt.Println("Subarray with maximum sum is", MaximumSubArrayOfSizeK(3, []int{2, 1, 5, 1, 3, 2}))
	fmt.Println("Subarray with maximum sum is", MaximumSubArrayOfSizeK(2, []int{2, 3, 4, 1, 5}))
}
