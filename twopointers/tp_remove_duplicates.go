package main

import (
	"fmt"
)

/*
Given an array of sorted numbers, remove all duplicates from it. You should not use any extra space;
after removing the duplicates in-place return the new length of the array.
*/
func removeDuplicates(input []int) int {
	fmt.Println("Original array", input)
	index, nextNonDuplicates := 1, 1
	for index < len(input) {
		if input[nextNonDuplicates-1] != input[index] {
			input[nextNonDuplicates] = input[index]
			nextNonDuplicates++
		}
		index++
	}
	fmt.Println("Removed duplicates", input[0:nextNonDuplicates])
	return nextNonDuplicates
}

func main() {
	fmt.Println(removeDuplicates([]int{2, 3, 3, 3, 6, 9, 9}))
}
