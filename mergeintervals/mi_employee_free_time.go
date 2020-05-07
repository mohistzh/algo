package main

import (
	"fmt"
	"sort"
)

// findEmployeeFreeTime
/*
For ‘K’ employees, we are given a list of intervals representing the working hours of each employee.
Our goal is to find out if there is a free interval that is common to all employees.
You can assume that each list of employee working hours is sorted on the start time.
*/
func findEmployeeFreeTime(schedule [][]int) [][]int {
	sortMatrixBySlice(schedule) // sorting matrix by first element of array in matrix
	fmt.Println(schedule)
	return nil
}

// sort 2d array by slice
func sortMatrixBySlice(matrix [][]int) {
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
	findEmployeeFreeTime([][]int{{1, 3}, {5, 6}, {2, 3}, {6, 8}})
	findEmployeeFreeTime([][]int{{1, 3}, {9, 12}, {2, 4}, {6, 8}})
	findEmployeeFreeTime([][]int{{1, 3}, {2, 4}, {3, 5}, {7, 9}})
}
