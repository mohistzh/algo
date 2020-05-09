package main

import (
	"fmt"
)

/*
We are given an unsorted array containing ‘n+1’ numbers taken from the range 1 to ‘n’.
The array has only one duplicate but it can be repeated multiple times.
Find that duplicate number without using any extra space. You are, however, allowed to modify the input array.
*/
func findDuplicateNumber(input []int) int {
	i, n := 0, len(input)
	for i < n {
		if input[i] != i+1 {
			j := input[i] - 1         // find the correct index of input[i]
			if input[i] != input[j] { //we are try to place each number on its correct index.
				input[i], input[j] = input[j], input[i]
			} else {
				// Because there is only one duplicate number
				// if while swapping the number with its index both the numbers being swapped are same,
				// we've found our duplicate.
				return input[i]
			}
		} else {
			i++
		}
	}
	return -1
}

func main() {
	fmt.Println(findDuplicateNumber([]int{1, 4, 4, 3, 2}))
	fmt.Println(findDuplicateNumber([]int{2, 1, 3, 3, 5, 4}))
	fmt.Println(findDuplicateNumber([]int{2, 4, 1, 4, 4}))
}
