package main

import (
	"fmt"
)

/*
 We are given an array containing ‘n’ distinct numbers taken from the range 0 to ‘n’.
 Since the array has only ‘n’ numbers out of the total ‘n+1’ numbers, find the missing number.
*/
func findMissingNumber(input []int) int {
	i, n := 0, len(input)
	for i < n {
		j := input[i]
		if j < n && j != input[j] { // if the index and value wasn't match, then swap them.
			input[i], input[j] = input[j], input[i]
		} else {
			i++
		}
	}
	for i := 0; i < n; i++ {
		if input[i] != i {
			return i
		}
	}
	return -1
}
func main() {
	input1 := []int{4, 0, 3, 1}
	input2 := []int{8, 3, 5, 2, 4, 6, 0, 1}
	fmt.Println(input1, findMissingNumber(input1))
	fmt.Println(input2, findMissingNumber(input2))
}
