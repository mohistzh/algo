package main

import (
	"fmt"
	"math"
)

/*
	Given two lists of intervals, find the intersection of these two lists.
	Each list consists of disjoint intervals sorted on their start time.
	e.g.
	Input: arr1=[[1, 3], [5, 6], [7, 9]], arr2=[[2, 3], [5, 7]]
	Output: [2, 3], [5, 6], [7, 7]
	Explanation: The output list contains the common intervals between the two lists.
*/
func intervalIntersection(intervalsA [][]int, intervalsB [][]int) [][]int {
	// c.start = max(a.start, b.start) c.end = min(a.end, b.end) if c.end < c.start then skip
	var intervalsC [][]int
	for i := 0; i < len(intervalsA); i++ {
		subArrayA := intervalsA[i]
		for j := 0; j < len(intervalsB); j++ {
			subArrayB := intervalsB[j]
			start := int(math.Max(float64(subArrayA[0]), float64(subArrayB[0])))
			end := int(math.Min(float64(subArrayA[1]), float64(subArrayB[1])))
			if end < start {
				continue
			} else {
				intervalsC = append(intervalsC, []int{start, end})
			}
		}
	}
	return intervalsC
}
func intervalIntersectionSolution2(intervalA [][]int, intervalB [][]int) [][]int {
	var intervalC [][]int
	indexA, indexB := 0, 0 // index of original interval array
	const start, end = 0, 1
	for indexA < len(intervalA) && indexB < len(intervalB) {
		// check if intervals overlap and intervalA startime lies within the other intervalB, vice versa
		isAOverlapsB := intervalA[indexA][start] >= intervalB[indexB][start] && intervalA[indexA][start] <= intervalB[indexB][end]
		isBOverlapsA := intervalB[indexB][start] >= intervalA[indexA][start] && intervalB[indexB][start] <= intervalA[indexA][end]

		if isAOverlapsB || isBOverlapsA {
			// intersection part
			maxStart := math.Max(float64(intervalA[indexA][start]), float64(intervalB[indexB][start]))
			minEnd := math.Min(float64(intervalA[indexA][end]), float64(intervalB[indexB][end]))
			intervalC = append(intervalC, []int{int(maxStart), int(minEnd)})
		}
		// move next from the interval which is finishing first
		if intervalA[indexA][end] < intervalB[indexB][end] {
			indexA++
		} else {
			indexB++
		}
	}
	return intervalC

}

func main() {
	fmt.Println(intervalIntersection([][]int{{1, 3}, {5, 6}, {7, 9}}, [][]int{{2, 3}, {5, 7}}))
	fmt.Println(intervalIntersection([][]int{{1, 3}, {5, 7}, {9, 12}}, [][]int{{5, 10}}))
	fmt.Println("Second solution")
	fmt.Println(intervalIntersectionSolution2([][]int{{1, 3}, {5, 6}, {7, 9}}, [][]int{{2, 3}, {5, 7}}))
	fmt.Println(intervalIntersectionSolution2([][]int{{1, 3}, {5, 7}, {9, 12}}, [][]int{{5, 10}}))
}
