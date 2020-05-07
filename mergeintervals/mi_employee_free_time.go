package main

// findEmployeeFreeTime
/*
For ‘K’ employees, we are given a list of intervals representing the working hours of each employee.
Our goal is to find out if there is a free interval that is common to all employees.
You can assume that each list of employee working hours is sorted on the start time.
*/
func findEmployeeFreeTime(schedule [][]int) [][]int {
	return nil
}

func main() {
	findEmployeeFreeTime([][]int{{1, 3}, {5, 6}, {2, 3}, {6, 8}})
	findEmployeeFreeTime([][]int{{1, 3}, {9, 12}, {2, 4}, {6, 8}})
	findEmployeeFreeTime([][]int{{1, 3}, {2, 4}, {3, 5}, {7, 9}})
}
