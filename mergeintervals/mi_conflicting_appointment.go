package main

import (
	"fmt"
)

/*
 Given an array of intervals representing 'N' appointments, find out if a person can attend all the appointments.
*/
func attendAllAppointments(appointments [][]int) bool {
	const start, end = 0, 1
	// i, j
	//a[i] overlaps a[j] appointments[i][start] <= appointments[j][end] && appointments[i][end] >= appointments[j][start]
	//a[j] overlaps a[i] appointments[j][start] <= appointments[i][end] && appointments[j][end] >= appointments[i][start]

	for i := 0; i < len(appointments); i++ {
		a := appointments[i]
		for j := i + 1; j < len(appointments); j++ {
			b := appointments[j]
			aOb := a[start] <= b[end] && a[end] >= b[start]
			bOa := b[start] <= a[end] && b[end] >= a[start]
			if aOb || bOa {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(attendAllAppointments([][]int{{1, 4}, {2, 5}, {7, 9}}))
	fmt.Println(attendAllAppointments([][]int{{6, 7}, {2, 4}, {8, 12}}))
	fmt.Println(attendAllAppointments([][]int{{4, 5}, {2, 3}, {3, 6}}))
}
