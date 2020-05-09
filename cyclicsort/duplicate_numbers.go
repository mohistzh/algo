package main

import (
	"fmt"
)

/*
We are given an unsorted array containing ‘n’ numbers taken from the range 1 to ‘n’.
The array has some duplicates, find all the duplicate numbers without using any extra space.
*/
func findDuplicateNumbers(input []int) []int {
	i, n := 0, len(input)
	var result []int
	for i < n {
		if input[i] != i+1 {
			j := input[i] - 1
			if input[i] != input[j] {
				input[i], input[j] = input[j], input[i]
			} else {
				result = append(result, input[i])
				i++
			}
		} else {
			i++
		}

	}
	return result
}

func main() {
	fmt.Println(findDuplicateNumbers([]int{3, 4, 4, 5, 5}))
	fmt.Println(findDuplicateNumbers([]int{5, 4, 7, 2, 3, 5, 3}))
}
