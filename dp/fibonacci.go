package main

import (
	"fmt"
)

/*
	In this approach, we try to solve the bigger problem by recursively finding the solution to
	smaller subproblems. Whenever we solve a subproblem, we cache its result so that we don't end up solving
	it repeatedly if it's called multiple times. Instead, we can just return the saved result.
	This technique of storing the results of already solved subproblems is called Memoization.
*/
func fibDP1(n int) int {
	memoize := make([]int, n+1)
	return memoization(n, memoize)
}
func memoization(n int, memoize []int) int {
	if n < 2 {
		return n
	}
	if memoize[n] != 0 {
		return memoize[n]
	}
	memoize[n] = memoization(n-1, memoize) + memoization(n-2, memoize)
	return memoize[n]
}

func fibDP2(n int) int {
	return tabulation(n)
}
func tabulation(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println("Memoization Approach")
	fmt.Println("5th Fibonacci is:", fibDP1(5))
	fmt.Println("6th Fibonacci is:", fibDP1(6))
	fmt.Println("7th Fibonacci is:", fibDP1(7))
	fmt.Println("Tabulation Approach")
	fmt.Println("5th Fibonacci is:", fibDP2(5))
	fmt.Println("6th Fibonacci is:", fibDP2(6))
	fmt.Println("7th Fibonacci is:", fibDP2(7))

}
