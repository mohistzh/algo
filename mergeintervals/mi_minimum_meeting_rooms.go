package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

const start, end = 0, 1

// MinHeap to find out the smallest meeting end time
// Minheap restrictions are !mh.Less(j, i) for 0 <= i < mh.Len() and 2*i+1 <= j <= 2*i+2 and j < mh.Len()
type MinHeap [][]int

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i][end] > mh[j][end]
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

// Push add an item
func (mh *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*mh = append(*mh, x.([]int))
}

// Pop popup an item
func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

// Peak get the Peak element
func (mh *MinHeap) Peak() interface{} {
	old := *mh
	n := old.Len()
	if n < 1 {
		return []int{-1, -1}
	}
	return old[n-1]
}

/*
Given a list of intervals representing the start and end time of 'N' meetings.
find the minimum number of rooms required to hold all the meetings.
e.g.
Meetings: [[1, 4], [2, 5], [7, 9]]
Output: 2
Explanation: Since [1, 4] and [2, 5] overlap, we need two rooms to hold these two meetings.
[7, 9] can occur in any of the two rooms later.

Similar problems:
1. Given a list of intervals, find the point where the maximum number of intervals overlap.
2. Given a list of intervals representing the arrival and departure times of trains to a train station,
	our goal is to find the minimum number of platforms required for the train station so that no train has to wait.
*/

func minimumMeetingRooms(meetings [][]int) int {
	sortMeetings(meetings)
	minRooms := 0
	minHeap := &MinHeap{}
	for i := 0; i < len(meetings); i++ {
		meeting := meetings[i]
		minHeapPeekEnd := minHeap.Peak().([]int)[1]
		if minHeapPeekEnd != -1 {
			for minHeap.Len() > 0 && meeting[start] >= minHeapPeekEnd { // remove all meetings that had ended
				heap.Pop(minHeap)
			}
		}
		heap.Push(minHeap, meeting)
		heap.Fix(minHeap, minHeap.Len()-1)
		// all active meetings are in the min heap, so we need for all of them
		minRooms = int(math.Max(float64(minRooms), float64(minHeap.Len())))
	}
	return minRooms
}

// sort 2d array by slice
func sortMeetings(matrix [][]int) {
	sort.Slice(matrix[:], func(i, j int) bool {
		for x := range matrix[i] {
			if matrix[i][x] == matrix[j][x] {
				continue
			}
			return matrix[i][x] < matrix[j][x]
		}
		return false
	})
}

func main() {
	fmt.Println(minimumMeetingRooms([][]int{{1, 4}, {2, 5}, {7, 9}}))
	fmt.Println(minimumMeetingRooms([][]int{{6, 7}, {2, 4}, {8, 12}}))
	fmt.Println(minimumMeetingRooms([][]int{{1, 4}, {2, 3}, {3, 6}}))
	fmt.Println(minimumMeetingRooms([][]int{{4, 5}, {2, 3}, {2, 4}, {3, 5}}))
}
