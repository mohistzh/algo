package main

import (
	"fmt"
)

/*
	Given a value N, if we want to make change for N cents, and we have infinite supply of each of S = { S1, S2, .. , Sm} valued coins,
	how many ways can we make the change? The order of coins doesn’t matter.
	For example, for N = 4 and S = {1,2,3}, there are four solutions: {1,1,1,1},{1,1,2},{2,2},{1,3}.
	So output should be 4. For N = 10 and S = {2, 5, 3, 6}, there are five solutions: {2,2,2,2,2}, {2,2,3,3}, {2,2,6}, {2,3,5} and {5,5}.
	So the output should be 5.
*/

func CoinChangeRecursiveSolution(coins []int, n int, amount int) int {
	if amount == 0 {
		return 1
	}
	if amount < 0 {
		return 0
	}
	if n <= 0 && amount >= 1 {
		return 0
	}
	// f(coins, m, n) = f(coins, m-1, n) + f(coins, m, N-coins[m-1])
	return CoinChangeRecursiveSolution(coins, n-1, amount) + CoinChangeRecursiveSolution(coins, n, amount-coins[n-1])
}
func CoinChangeDPSolution(coins []int, n int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 3}
	res := CoinChangeRecursiveSolution(coins, len(coins), 4)
	fmt.Println(res)
	res = CoinChangeDPSolution(coins, len(coins), 4)
	fmt.Println(res)
}
