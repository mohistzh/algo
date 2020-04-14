package main

import (
	"fmt"
)

/*
Given an unsorted array of numbers and a target ‘key’,
remove all instances of ‘key’ in-place and return the new length of the array.
*/
func removeAllKey(input []int, key int) int {
	fmt.Println("Original array", input)
	nextElement := 0
	for i := 0; i < len(input); i++ {
		if input[i] != key {
			input[nextElement] = input[i]
			nextElement++
		}
	}
	fmt.Println("Fixed array", input[0:nextElement])
	return nextElement
}

func main() {
	fmt.Println(removeAllKey([]int{3, 2, 3, 6, 3, 10, 9, 3}, 3))
	fmt.Println(removeAllKey([]int{2, 11, 2, 2, 1}, 2))
}
