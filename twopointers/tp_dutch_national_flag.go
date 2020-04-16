package main

import (
	"fmt"
)

/*
Given an array containing 0s, 1s and 2s, sort the array in-place. You should treat numbers of the array as objects, hence, we can’t count 0s, 1s, and 2s to recreate the array.

The flag of the Netherlands consists of three colors: red, white and blue; and since our input array also consists of three different numbers that is why it is called Dutch National Flag problem.
*/
// move 0s before left, move 2s after right, keep 1s stay in middle
func dutchNationalFlagProblem(input []int) []int {
	index, left, right := 0, 0, len(input)-1
	for index <= right {
		if input[index] == 0 {
			input[index], input[left] = input[left], input[index] // move 0s before left
			left++
			index++
		} else if input[index] == 1 {
			index++
		} else if input[index] == 2 {
			input[index], input[right] = input[right], input[index] // move 2s after right
			right--
		} else {
			panic("invalid input")
		}
	}
	return input
}

func main() {
	fmt.Println(dutchNationalFlagProblem([]int{1, 0, 2, 1, 0}))
	fmt.Println(dutchNationalFlagProblem([]int{2, 2, 0, 1, 2, 0}))
}
