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

func SubsetSum(nums []int, sum int) int {
	return subsetSumRecursive(nums, len(nums), 0, sum, 0)
}

func subsetSumRecursive(nums []int, n int, index int, sum int, count int) int {
	// reached to the end
	if index == n {
		if sum == 0 {
			count++
		}
		return count
	}

	count = subsetSumRecursive(nums, n, index+1, sum-nums[index], count)
	count = subsetSumRecursive(nums, n, index+1, sum, count)
	return count
}

// SubsetSumTabulation tabulation solution
func SubsetSumTabulation(nums []int, sum int) int {
	n := len(nums)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, sum+1)
		dp[i][0] = 1
	}
	// fill first row
	for s := 1; s <= sum; s++ {
		state := 0
		if nums[0] == s {
			state = 1
		}
		dp[0][s] = state
	}
	for i := 1; i < n; i++ {
		for s := 1; s <= sum; s++ {
			dp[i%2][s] = dp[(i-1)%2][s]
			if s >= nums[i] {
				dp[i%2][s] += dp[(i-1)%2][s-nums[i]]
			}
		}
	}
	return dp[(n-1)%2][sum]
}

func main() {
	res := CountOfSubsetSum([]int{1, 1, 2, 3}, 4)
	fmt.Println(res)
	res = CountOfSubsetSum([]int{1, 2, 7, 1, 5}, 9)
	fmt.Println(res)
	nums := []int{1, 2, 7, 1, 5}
	res = SubsetSum(nums, 9)
	fmt.Println(res)
	res = SubsetSumTabulation(nums, 9)
	fmt.Println(res)
}
