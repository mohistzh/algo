package main

import (
	"fmt"
	"math"
)

/*
Given a list of intervals, merge all the overlapping intervals to produce a list that has only mutually exclusive intervals.
e.g.
Intervals: [[1,4], [2,5], [7,9]]
Output: [[1,5], [7,9]]
Explanation: Since the first two intervals [1,4] and [2,5] overlap, we merged them into one [1,5].
*/
// mergeIntervals Let's take the exmaple of two intervals('a' and 'b') such that a.start <= b.start.
// There are four possible scenarios
// 1. 'a' and 'b' do not overlap
// 2. some part of 'b' overlaps with 'a'
// 3. 'a' fully overlaps 'b'
// 4. 'b' fully overlaps 'a' but both have same start time
func mergeIntervals(input [][]int) [][]int {
	if len(input) < 2 {
		return input
	}
	var result [][]int
	start, end := input[0][0], input[0][1]
	for i := 1; i < len(input); i++ {
		interval := input[i]
		if interval[0] <= end { // overlapping intervals, adjust the 'end'
			end = int(math.Max(float64(end), float64(interval[1])))
		} else { // non-overlapping interval, add the previous interval and reset start and end
			result = append(result, []int{start, end})
			start = interval[0]
			end = interval[1]
		}
	}
	result = append(result, []int{start, end}) // append last interval
	return result
}

func main() {
	fmt.Println(mergeIntervals([][]int{{1, 4}, {2, 5}, {7, 9}}))
	fmt.Println(mergeIntervals([][]int{{2, 4}, {5, 9}, {6, 7}}))
	fmt.Println(mergeIntervals([][]int{{1, 4}, {2, 6}, {3, 5}}))
}
