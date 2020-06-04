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

/*
	Tabulation is opposite of the top-down approach and avoids recursion. In this approach, we solve the problem "bottom-up"
	(i.e. by solving all the related subproblems first). This is typically done by filling up an n-dimensional table.
	Based on the results in the table, the solution to the top/original problem is then computed.

	Tabulation is the opposite of Memoization, as in Memoization we solve the problem and maintain a map of already solved subproblems.
	In other words, in memoization, we do it top-down in the sense that we solve the top problem first(which typically recurses down to solve the subproblems).
*/
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
