package main

import "fmt"

/*
	Given a set of intervals, find out if any two intervals overlap.
*/
func detectOverlap(matrix [][]int) bool {
	if len(matrix) < 2 {
		return false
	}
	end := matrix[0][1]

	for i := 1; i < len(matrix); i++ {
		intervals := matrix[i]
		if end > intervals[0] {
			return true
		}
		end = intervals[1]
	}
	return false
}

func main() {
	fmt.Println(detectOverlap([][]int{{1, 4}, {2, 5}, {7, 9}}))
	fmt.Println(detectOverlap([][]int{{2, 4}, {5, 9}}))
	fmt.Println(detectOverlap([][]int{{1, 4}, {2, 6}, {3, 7}}))
}
