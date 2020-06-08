package main

import (
	"fmt"
)

// CanPartition
// Given a set of positive numbers, find if we can partition it into two subsets such that the sum of elements
// into both the subsets is equal.
func CanPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	return canPartitionRecursive(nums, sum/2, 0)
}

func canPartitionRecursive(nums []int, sum int, currentIndex int) bool {
	if sum == 0 {
		return true
	}
	if len(nums) == 0 || currentIndex >= len(nums) {
		return false
	}
	if nums[currentIndex] <= sum {
		if canPartitionRecursive(nums, sum-nums[currentIndex], currentIndex+1) {
			return true
		}
	}
	return canPartitionRecursive(nums, sum, currentIndex+1)
}

func main() {
	res := CanPartition([]int{1, 2, 3, 4})
	fmt.Println(res)
}
