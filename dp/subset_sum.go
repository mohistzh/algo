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
func main() {

	res := HasPartition([]int{1, 2, 7, 1, 5}, 10)
	fmt.Println(res)
	res = HasPartition([]int{1, 3, 4, 8}, 6)
	fmt.Println(res)
}
