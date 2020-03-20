package main

/*
	Edit distance problem
*/
import (
	"fmt"
)

// DP tabulation solution, bottom-top
func EditDistanceByTabulation(a string, b string) int {
	m, n := len(a), len(b)
	// use DP table to track status
	dp := make([][]int, m+1, n+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			// initial coloumns of first row, and rows of first column
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] // a and b have a same character, so just skip it.
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1 // choose a minimum value from insertion, deletion, substitution, + 1 means count for the operation
			}
		}
	}
	return dp[m][n]

}

// Recursively solution
func EditDistanceByRecursion(a string, b string, m int, n int) int {
	// base cases
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	// if the characters are the same, just skip it
	if a[m-1] == b[n-1] {
		return EditDistanceByRecursion(a, b, m-1, n-1)
	}
	return min(
		EditDistanceByRecursion(a, b, m, n-1),   // insertion
		EditDistanceByRecursion(a, b, m-1, n),   // deletion
		EditDistanceByRecursion(a, b, m-1, n-1), //substitution
	) + 1

}
func min(x int, y int, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	} else {
		return z
	}
}

func main() {
	a, b := "sunday", "saturday"
	distance := EditDistanceByTabulation(a, b)
	fmt.Printf("DP with tabulation solution. String '%v' and String '%v' \n", a, b)
	fmt.Println("Their edit distance is:", distance)
	fmt.Println()
	distance = EditDistanceByRecursion(a, b, len(a), len(b))
	fmt.Printf("DP with recursively solution. String '%v' and String '%v' \n", a, b)
	fmt.Println("Their edit distance is:", distance)

}
