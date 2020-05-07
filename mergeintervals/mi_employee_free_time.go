package main

import (
	"fmt"
	"math"
	"sort"
)

// findEmployeeFreeTime
/*
For ‘K’ employees, we are given a list of intervals representing the working hours of each employee.
Our goal is to find out if there is a free interval that is common to all employees.
You can assume that each list of employee working hours is sorted on the start time.
*/
func findEmployeeFreeTime(employees [][]int) [][]int {
	sortMatrixBySlice(employees) // sorting matrix by first element of array in matrix
	prevMaxEnd := math.MaxInt32
	var result [][]int
	const start, end = 0, 1
	employee := employees[start]
	for i := 1; i < len(employees); i++ {
		if prevMaxEnd < employees[i][start] { // comparing previous maximum end with current start
			result = append(result, []int{prevMaxEnd, employees[i][start]})
		}
		prevMaxEnd = int(math.Max(float64(employee[end]), float64(employees[i][end])))
		employee = employees[i]
	}
	return result
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
	fmt.Println(findEmployeeFreeTime([][]int{{1, 3}, {5, 6}, {2, 3}, {6, 8}}))
	fmt.Println(findEmployeeFreeTime([][]int{{1, 3}, {9, 12}, {2, 4}, {6, 8}}))
	fmt.Println(findEmployeeFreeTime([][]int{{1, 3}, {2, 4}, {3, 5}, {7, 9}}))
}
