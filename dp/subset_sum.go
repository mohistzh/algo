package main

import (
	"fmt"
)

/*
	Given a set of positive numbers, determine if there exists a subset whose sum is equal to given number 'S'
*/
func HasPartition(nums []int, target int) bool {
	return hasPartitionRecursive(nums, 0, 0, target)
}
func hasPartitionRecursive(nums []int, index int, sum int, target int) bool {
	if sum == target {
		return true
	}
	if sum > target || index >= len(nums) {
		return false
	}

	return hasPartitionRecursive(nums, index+1, sum, target) || hasPartitionRecursive(nums, index+1, sum+nums[index], target)
}
func main() {

	res := HasPartition([]int{1, 2, 7, 1, 5}, 10)
	fmt.Println(res)
	res = HasPartition([]int{1, 3, 4, 8}, 6)
	fmt.Println(res)
}
