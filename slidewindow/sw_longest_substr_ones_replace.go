package main

import (
	"fmt"
	"math"
)

/*
	Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s,
	find the length of the longest contiguous subarray having all 1s.

	e.g.

	Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
	Output: 6
	Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.
*/
func LongestSubArrayWithOnesAfterReplace(input []int, k int) int {
	res, maxOnes, windowStart := 0.0, 0, 0
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		if input[windowEnd] == 1 {
			maxOnes++
		}
		if (windowEnd - windowStart + 1 - maxOnes) > k {
			if input[windowStart] == 1 {
				maxOnes--
			}
			windowStart++
		}
		res = math.Max(float64(res), float64(windowEnd-windowStart+1))
	}
	return int(res)
}

func main() {

	fmt.Println("Longest Subarray with ones after replacement:",
		LongestSubArrayWithOnesAfterReplace([]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2))

	fmt.Println("Longest Subarray with ones after replacement:",
		LongestSubArrayWithOnesAfterReplace([]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 3))
}
