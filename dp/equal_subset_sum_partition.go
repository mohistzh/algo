package main

import (
	"fmt"
)

/*
Given a set of positive numbers, find if we can partition it into two subsets such that the sum of elements into both the subsets is equal.
*/
// CanPartition brute-force solution
func CanPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	return canPartitionRecursive(nums, sum/2, len(nums))
}

/*
	Let recursive function (nums, sum/2, n) be the function that returns true if there is a subset of nums[0..n-1] with sum equal to sum/2
	The recursive problem can be divided into two subproblems
	a) recursive func without considering last element (reducing n to n-1)
	b) recursive func considering the last element (reducing sum/2 by nums[n-1] and n to n-1)
	If any of the above subproblems return true, then return true.
	recursiveFunc(nums, sum/2, n) = recursiveFunc(nums, sum/2, n-1) || recursiveFunc(nums, sum/2 - nums[n-1], n-1)
*/
func canPartitionRecursive(nums []int, sum int, index int) bool {
	if sum == 0 {
		return true
	}
	if sum != 0 && index == 0 {
		return false
	}
	if nums[index-1] > sum {
		return canPartitionRecursive(nums, sum, index-1)
	}
	return canPartitionRecursive(nums, sum, index-1) || canPartitionRecursive(nums, sum-nums[index-1], index-1)
}

func main() {
	res := CanPartition([]int{1, 2, 3, 4})
	fmt.Println(res)
}
