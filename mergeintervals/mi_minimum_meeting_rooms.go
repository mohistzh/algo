package main

import (
	"fmt"
)

/*
Given a list of intervals representing the start and end time of 'N' meetings.
find the minimum number of rooms required to hold all the meetings.
e.g.
Meetings: [[1, 4], [2, 5], [7, 9]]
Output: 2
Explanation: Since [1, 4] and [2, 5] overlap, we need two rooms to hold these two meetings.
[7, 9] can occur in any of the two rooms later.
*/
func minimumMeetingRooms(meetingRooms [][]int) int {
	const start, end = 0, 1
	roomsRequired := 1
	for i := 0; i < len(meetingRooms); i++ {
		a := meetingRooms[i]
		for j := i + 1; j < len(meetingRooms); j++ {
			b := meetingRooms[j]
			aOb := a[start] <= b[end] && a[end] >= b[start]
			bOa := b[start] <= a[end] && b[end] >= a[start]

			if aOb || bOa {
				roomsRequired++
			}
		}
	}
	return roomsRequired
}

func main() {
	fmt.Println(minimumMeetingRooms([][]int{{1, 4}, {2, 5}, {7, 9}}))
	fmt.Println(minimumMeetingRooms([][]int{{6, 7}, {2, 4}, {8, 12}}))
	fmt.Println(minimumMeetingRooms([][]int{{1, 4}, {2, 3}, {3, 6}}))
	fmt.Println(minimumMeetingRooms([][]int{{4, 5}, {2, 3}, {2, 4}, {3, 5}}))
}
