package main

/*
 Longest common subsequence, refer https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
*/
import (
	"fmt"
)

//  O(2^n)  time complexity
func LongestCommonSubSequenceNaive(a string, b string, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if a[m-1] == b[n-1] {
		return 1 + LongestCommonSubSequenceNaive(a, b, m-1, n-1)
	} else {
		// gussing properteis of the question,
		return max(
			LongestCommonSubSequenceNaive(a, b, m, n-1),
			LongestCommonSubSequenceNaive(a, b, m-1, n),
		)
	}
}

// O(mn) time complexity
func LongestCommonSubSequenceDP(a string, b string) int {
	m, n := len(a), len(b)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]

}
func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func main() {
	a, b := "AGGTAB", "GXTXAYB"
	fmt.Println("Naive solution: ", LongestCommonSubSequenceNaive(a, b, len(a), len(b)))
	fmt.Println("DP solution: ", LongestCommonSubSequenceDP(a, b))
}
