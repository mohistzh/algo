package main

import (
	"fmt"
	"sort"
)

// Meeting meeting structure
type Meeting struct {
	start int
	end   int
}

/*
Given a list of intervals representing the start and end time of 'N' meetings.
find the minimum number of rooms required to hold all the meetings.
e.g.
Meetings: [[1, 4], [2, 5], [7, 9]]
Output: 2
Explanation: Since [1, 4] and [2, 5] overlap, we need two rooms to hold these two meetings.
[7, 9] can occur in any of the two rooms later.
*/

func minimumMeetingRooms(meetings [][]int) int {
	sortMeetings(meetings)
	fmt.Println(meetings)
	return -1
}
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
