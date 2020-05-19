package main

import (
	"fmt"
)

func findMaximumGiftValue(input [][]int) int {
	row, col := len(input), len(input[0])
	if row == 0 || col == 0 {
		return -1
	}
	dp := make([][]int, row)
	for i := 0; i < col; i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			left, up, max := 0, 0, 0
			if i > 0 {
				up = dp[i-1][j]
			}
			if j > 0 {
				left = dp[i][j-1]
			}
			if left > up {
				max = left
			} else {
				max = up
			}
			dp[i][j] = max + input[i][j]
		}
	}
	return dp[row-1][col-1]
}

func main() {
	input := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(findMaximumGiftValue(input))
}
