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

// MinimumSubSetSumDifferentMemoization based on naive solution, add a cache mechanism
func MinimumSubSetSumDifferentMemoization(nums []int) int {
	memo := make(map[string]int)
	return minimumSubSetSumDifferentMemoization(nums, 0, 0, 0, memo)
}
func minimumSubSetSumDifferentMemoization(nums []int, index int, sum1 int, sum2 int, memo map[string]int) int {
	current := fmt.Sprint(index, "", sum1)
	if index == len(nums) {
		return intAbs(sum1 - sum2)
	}
	if _, ok := memo[current]; !ok {
		diff1 := minimumSubSetSumDifferentMemoization(nums, index+1, sum1+nums[index], sum2, memo)
		diff2 := minimumSubSetSumDifferentMemoization(nums, index+1, sum1, sum2+nums[index], memo)
		memo[current] = int(math.Min(float64(diff1), float64(diff2)))
	}
	return memo[current]
}

// MinimumSubSetSumDifferentTabulation  bottom-up solution
func MinimumSubSetSumDifferentTabulation(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
		total += nums[i]
	}
	dp := make([][]bool, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]bool, total/2+1)
		dp[i][0] = true
	}

	for s := 1; s <= total/2; s++ {
		state := false
		if nums[0] == s {
			state = true
		}
		dp[0][s] = state
	}
	for i := 1; i < n; i++ {
		for s := 1; s <= total/2; s++ {
			if dp[(i-1)%2][s] {
				dp[i%2][s] = dp[(i-1)%2][s]
			} else if s >= nums[i] {
				dp[i%2][s] = dp[(i-1)%2][s-nums[i]]
			}
		}
	}
	sum1 := 0
	for i := total / 2; i >= 0; i-- {
		if dp[(n-1)%2][i] {
			sum1 = i
			break
		}
	}
	sum2 := total - sum1
	return intAbs(sum1 - sum2)

}

// lazy to convert data type
func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("Naive solution")
	res := MinimumSubSetSumDifferent([]int{1, 2, 3, 9})
	fmt.Println("{1, 2, 3, 9}", res)
	res = MinimumSubSetSumDifferent([]int{1, 3, 100, 4})
	fmt.Println("{1, 3, 100, 4}", res)

	fmt.Println()
	fmt.Println("Memoization solution")
	res = MinimumSubSetSumDifferentMemoization([]int{1, 2, 3, 9})
	fmt.Println("{1, 2, 3, 9}s", res)
	res = MinimumSubSetSumDifferentMemoization([]int{1, 3, 100, 4})
	fmt.Println("{1, 3, 100, 4}", res)

	fmt.Println()
	fmt.Println("Tabulation solution")
	res = MinimumSubSetSumDifferentTabulation([]int{1, 2, 3, 9})
	fmt.Println("{1, 2, 3, 9}", res)
	res = MinimumSubSetSumDifferentTabulation([]int{1, 3, 100, 4})
	fmt.Println("{1, 3, 100, 4}", res)
}
