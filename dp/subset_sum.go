package main

import (
	"fmt"
)

/*
	Given a set of positive numbers, determine if there exists a subset whose sum is equal to given number 'S'
*/
func HasPartition(nums []int, target int) bool {
	memo := make(map[string]bool)
	return hasPartitionRecursive(nums, 0, 0, target, memo)
}

// hasPartitionRecursive Top-bottom with memoization approach
func hasPartitionRecursive(nums []int, index int, sum int, target int, memo map[string]bool) bool {
	current := fmt.Sprint(index, "", sum)
	if _, ok := memo[current]; ok {
		return memo[current]
	}
	if sum == target {
		return true
	}
	if sum > target || index >= len(nums) {
		return false
	}
	memo[current] = hasPartitionRecursive(nums, index+1, sum, target, memo) || hasPartitionRecursive(nums, index+1, sum+nums[index], target, memo)
	return memo[current]
}

// HasPartitionTabulation tabulation approach
func HasPartitionTabulation(nums []int, target int) bool {
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	for i := 1; i <= target; i++ {
		state := false
		if nums[0] == i {
			state = true
		}
		dp[0][i] = state
	}

	for i := 1; i < len(nums); i++ {
		for s := 1; s <= target; s++ {
			if dp[i-1][s] {
				dp[i][s] = dp[i-1][s]
			} else if s >= nums[i] {
				dp[i][s] = dp[i-1][s-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][target]

}

// HasPartitionTabulation2 space complexity costs O(target)
func HasPartitionTabulation2(nums []int, target int) bool {
	dp := make([][]bool, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	for i := 1; i <= target; i++ {
		state := false
		if nums[0] == i {
			state = true
		}
		dp[0][i] = state
	}

	for i := 1; i < len(nums); i++ {
		for s := 1; s <= target; s++ {
			if dp[(i-1)%2][s] {
				dp[i%2][s] = dp[(i-1)%2][s]
			} else if s >= nums[i] {
				dp[i%2][s] = dp[(i-1)%2][s-nums[i]]
			}
		}
	}

	return dp[(len(nums)-1)%2][target]

}

func main() {

	res := HasPartition([]int{1, 2, 7, 1, 5}, 10)
	fmt.Println(res)
	res = HasPartition([]int{1, 3, 4, 8}, 6)
	fmt.Println(res)
	res = HasPartitionTabulation([]int{1, 2, 7, 1, 5}, 10)
	fmt.Println(res)
	res = HasPartitionTabulation([]int{1, 3, 4, 8}, 6)
	fmt.Println(res)
	res = HasPartitionTabulation2([]int{1, 3, 4, 8}, 6)
	fmt.Println(res)
	res = HasPartitionTabulation2([]int{1, 2, 7, 1, 5}, 10)
	fmt.Println(res)
}
