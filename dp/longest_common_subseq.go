package main

/*
 Longest common subsequence
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
func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func main() {
	a, b := "AGGTAB", "GXTXAYB"
	fmt.Println(LongestCommonSubSequenceNaive(a, b, len(a), len(b)))
}
