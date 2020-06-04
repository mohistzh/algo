package main

/*
	Knapsack problem, refer https://en.wikipedia.org/wiki/Knapsack_problem
*/
import (
	"fmt"
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
	fmt.Println(dp)
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
}
