package main

import (
	"fmt"
)

/*
	Length of the shortest supersequence  = (Sum of lengths of given two strings) -
                                        (Length of LCS of two given strings)
*/
func ShortestSuperSequence(a string, b string) int {
	m := len(a)
	n := len(b)
	l := lcs(a, b, m, n)
	return (m + n - l)
}

func lcs(a string, b string, m int, n int) int {
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
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	a, b := "AGGTAB", "GXTXAYB"
	scs := ShortestSuperSequence(a, b)
	fmt.Println(scs)

}
