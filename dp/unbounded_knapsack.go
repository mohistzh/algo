package main

import (
	"fmt"
	"math"
)

/*
	Given the weights and profits of 'N' items, we are asked to put these items in a knapsack which has a
	capacity 'C'. The goal is to get the maximum profit from the items in the knapsack. The only difference
	between the 0/1 knapsack problem and this problem is that we are allowed to use an unlimited quantity
	of an item.
*/
func solveUnboundedKnapsack(profits []int, weights []int, capacity int) int {
	return unboundedKnapsackRecursive(profits, weights, capacity, 0)
}
func unboundedKnapsackRecursive(profits []int, weights []int, capacity int, currentIndex int) int {
	// base conditions
	if currentIndex <= 0 || len(profits) == 0 || len(profits) != len(weights) || currentIndex >= len(profits) {
		return 0
	}
	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + unboundedKnapsackRecursive(profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}
	profit2 := unboundedKnapsackRecursive(profits, weights, capacity, currentIndex+1)
	return int(math.Max(float64(profit1), float64(profit2)))
}

func main() {
	profits := []int{15, 50, 60, 90}
	weights := []int{1, 3, 4, 5}
	fmt.Println(solveUnboundedKnapsack(profits, weights, 8))
}
