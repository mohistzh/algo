package main

import (
	"fmt"
)

/*
We are given an array containing ‘n’ objects. Each object, when created, was assigned a unique number from 1 to ‘n’ based on their creation sequence.
This means that the object with sequence number ‘3’ was created just before the object with sequence number ‘4’.

Write a function to sort the objects in-place on their creation sequence number in O(n) and without any extra space.
For simplicity, let’s assume we are passed an integer array containing only the sequence numbers, though each number is actually an object.
*/
func cyclicSort(input []int) []int {
	i := 0
	for i < len(input) {
		j := input[i] - 1
		if input[i] != input[j] { // swap value, don't move pointer
			input[i], input[j] = input[j], input[i]
		} else {
			i++
		}
	}
	return input
}

func main() {
	fmt.Println(cyclicSort([]int{3, 1, 5, 4, 2}))
	fmt.Println(cyclicSort([]int{2, 6, 4, 3, 1, 5}))
	fmt.Println(cyclicSort([]int{1, 5, 6, 4, 3, 2}))
}
