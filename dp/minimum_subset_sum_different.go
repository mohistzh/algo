package main

import (
	"fmt"
	"math"
)

/*
	Given a set of positive numbers, partition the set into two subsets with a minimum different between their subset sums
*/
func MinimumSubSetSumDifferent(nums []int) int {

	return minimumSubSetSumDifferentRecursive(nums, 0, 0, 0)
}

// naive approach, it costs time complexity O(2^n) where n represents numbers of nums array
// space complexity costs O(n) due to store the recursion stack
func minimumSubSetSumDifferentRecursive(nums []int, index int, sum1 int, sum2 int) int {
	// base case
	if index == len(nums) {
		return intAbs(sum1 - sum2)
	}
	diff1 := minimumSubSetSumDifferentRecursive(nums, index+1, sum1+nums[index], sum2)
	diff2 := minimumSubSetSumDifferentRecursive(nums, index+1, sum1, sum2+nums[index])

	return int(math.Min(float64(diff1), float64(diff2)))
}

// lazy to convert data type
func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	res := MinimumSubSetSumDifferent([]int{1, 2, 3, 9})
	fmt.Println(res)
	res = MinimumSubSetSumDifferent([]int{1, 3, 100, 4})
	fmt.Println(res)
}
