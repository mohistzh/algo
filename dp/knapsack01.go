package main

/*
	Knapsack problem, refer https://en.wikipedia.org/wiki/Knapsack_problem
*/
import (
	"fmt"
	"math"
)

// KnapsackNaiveSolution Brute-force solution which recursive all passible value
func KnapsackNaiveSolution(capacity int, weight []int, value []int, n int) int {
	if n == 0 || capacity == 0 {
		return 0
	}
	if weight[n-1] > capacity {
		return KnapsackNaiveSolution(capacity, weight, value, n-1)
	}
	return max(value[n-1]+
		KnapsackNaiveSolution(capacity-weight[n-1], weight, value, n-1),
		KnapsackNaiveSolution(capacity, weight, value, n-1),
	)

}

// KnapsackMemoizationSolution Top-down DP with Memoization
func KnapsackMemoizationSolution(capacity int, weight []int, value []int) int {
	memo := make([][]int, len(value))
	for i := 0; i < len(value); i++ {
		memo[i] = make([]int, capacity+1)
	}
	return knapsackMemoization(memo, weight, value, capacity, 0)
}

func knapsackMemoization(memo [][]int, weight []int, value []int, capacity int, currentIndex int) int {
	// base condition
	if capacity <= 0 || currentIndex >= len(weight) {
		return 0
	}
	// memoization
	if memo[currentIndex][capacity] != 0 {
		return memo[currentIndex][capacity]
	}
	profitA := 0
	if weight[currentIndex] <= capacity {
		profitA = value[currentIndex] + knapsackMemoization(memo, weight, value, capacity-weight[currentIndex], currentIndex+1)
	}

	profitB := knapsackMemoization(memo, weight, value, capacity, currentIndex+1)
	memo[currentIndex][capacity] = int(math.Max(float64(profitA), float64(profitB)))
	return memo[currentIndex][capacity]

}

//KnapsackDPSolution Dynamic Programming solution
func KnapsackDPSolution(capacity int, weight []int, value []int, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if weight[i-1] <= j {
				dp[i][j] = max(value[i-1]+dp[i-1][j-weight[i-1]], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}

		}
	}
	return dp[n][capacity]

}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	value := []int{60, 100, 120}
	weight := []int{10, 20, 30}
	capacity := 50
	res := KnapsackNaiveSolution(capacity, weight, value, len(value))
	fmt.Println(res)
	res = KnapsackDPSolution(capacity, weight, value, len(value))
	fmt.Println(res)

	res = KnapsackMemoizationSolution(capacity, weight, value)
	fmt.Println(res)
}
