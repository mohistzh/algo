package main

import (
	"fmt"
	"math"
)

/*
	Given a list of non-overlapping intervals sorted by their start time, insert a given interval at the correct position
	and merge all necessary intervals to produce a list that has only mutually exclusive intervals.
	e.g.
	Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,6]
	Output: [[1,3], [4,7], [8,12]]
	Explanation: After insertion, since [4,6] overlaps with [5,7], we merged them into one [4,7].
*/
func insertInterval(matrix [][]int, newInterval []int) [][]int {
	var merged [][]int
	index, start, end := 0, 0, 1
	// newInterval.start < matrix.end
	for index < len(matrix) && matrix[index][end] < newInterval[start] {
		merged = append(merged, []int{matrix[index][start], matrix[index][end]})
		index++
	}
	// merge all intervals that overlap with newInterval
	for index < len(matrix) && matrix[index][start] <= newInterval[end] {
		newInterval[start] = int(math.Min(float64(matrix[index][start]), float64(newInterval[start])))
		newInterval[end] = int(math.Max(float64(matrix[index][end]), float64(newInterval[end])))
		index++
	}

	merged = append(merged, newInterval)
	for index < len(matrix) {
		merged = append(merged, matrix[index]) // merge the left part
		index++
	}
	return merged
}

func main() {
	fmt.Println(insertInterval([][]int{{1, 3}, {5, 7}, {8, 12}}, []int{4, 6}))
	fmt.Println(insertInterval([][]int{{1, 3}, {5, 7}, {8, 12}}, []int{4, 10}))
	fmt.Println(insertInterval([][]int{{2, 3}, {5, 7}}, []int{1, 4}))
}
