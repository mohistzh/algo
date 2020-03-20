package main

import (
	"fmt"
)

/*
	Longest incresing subsequence, refer https://en.wikipedia.org/wiki/Longest_increasing_subsequence
*/

func LongestIncresingSubSequenceByDP(input []int) int {
	n := len(input)
	if n <= 0 {
		return 0
	}
	// result
	res := 1
	dp := make([]int, n) // to store state of the subsequence, we can use this array getting actual elements from the origin array
	for i := 0; i < n; i++ {
		// each item's subsequence is themself
		dp[i] = 1
		// iterate through each element and their subsequence(from current point to the left side part of the array)
		for j := 0; j < i; j++ {
			// found an incresing number
			if input[i] > input[j] {
				dp[i] = max(dp[i], dp[j]+1)
				res = max(dp[i], res) // always assign the maximum to the res
			}
		}
	}
	print(input, dp)
	return res
}

func print(input []int, dp []int) {
	var stack []int
	previous := -1
	for i := 0; i < len(dp); i++ {
		if dp[i] > previous {
			stack = append(stack, input[i])
			previous = dp[i]
		} else if dp[i] == previous {
			stack = stack[:len(stack)-1] //pop item from a simple stack
			stack = append(stack, input[i])
		}
	}
	fmt.Println("Origin array is: ")
	fmt.Printf("%v\n", input)
	fmt.Println("DP table is: ")
	fmt.Printf("%v\n", dp)
	fmt.Println("Subsequence is: ")
	for len(stack) > 0 {
		n := len(stack) - 1
		fmt.Printf("%v ", stack[n])
		stack = stack[:n]
	}
	fmt.Println()

}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	input1 := []int{
		10, 22, 9, 33, 21, 50, 41, 60, 80,
	}
	input2 := []int{
		0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15,
	}
	LongestIncresingSubSequenceByDP(input1)
	LongestIncresingSubSequenceByDP(input2)

}
