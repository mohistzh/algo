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

// CanPartitionTabulationSolution bottom-up solution
func CanPartitionTabulationSolution(nums []int) bool {
	sum, n := 0, len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	// With only one number, we can form a subset only when the required sum is equal to its value
	for s := 1; s <= sum; s++ {
		t := true
		if s != nums[0] {
			t = false
		}
		dp[0][s] = t
	}

	for r := 1; r < n; r++ {
		for s := 1; s <= sum; s++ {
			if dp[r-1][s] {
				dp[r][s] = dp[r-1][s]
			} else if s >= nums[r] {
				dp[r][s] = dp[r-1][s-nums[r]]
			}
		}
	}
	return dp[n-1][sum]
}

func main() {
	res := CanPartition([]int{1, 2, 3, 4})
	fmt.Println(res)
	res = CanPartition([]int{1, 5, 7})
	fmt.Println(res)
	res = CanPartitionTabulationSolution([]int{1, 2, 3, 4})
	fmt.Println(res)
}
