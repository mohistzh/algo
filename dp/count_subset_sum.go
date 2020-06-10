package main

import (
	"fmt"
)

/*
	Given a set of positive numbers, find the total number of subsets whose sum is equal to a given number 'sum'
*/
func CountOfSubsetSum(nums []int, sum int) int {
	memo := make(map[string]int)
	return countOfSubsetSumRecursive(nums, sum, 0, memo)
}
func countOfSubsetSumRecursive(nums []int, sum int, currentIndex int, memo map[string]int) int {
	current := fmt.Sprint(currentIndex, "", sum)
	if _, ok := memo[current]; ok {
		return memo[current]
	}
	if sum == 0 {
		return 1
	}
	if currentIndex >= len(nums) {
		return 0
	}
	num1 := 0
	if nums[currentIndex] <= sum {
		num1 = countOfSubsetSumRecursive(nums, sum-nums[currentIndex], currentIndex+1, memo)
	}
	num2 := countOfSubsetSumRecursive(nums, sum, currentIndex+1, memo)
	memo[current] = num1 + num2
	return memo[current]
}

func main() {
	res := CountOfSubsetSum([]int{1, 1, 2, 3}, 4)
	fmt.Println(res)
	res = CountOfSubsetSum([]int{1, 2, 7, 1, 5}, 9)
	fmt.Println(res)
}
