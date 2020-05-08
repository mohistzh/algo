package main

import (
	"fmt"
)

/*
We are given an unsorted array containing numbers taken from the range 1 to ‘n’.
The array can have duplicates, which means some numbers will be missing. Find all those missing numbers.
e.g.

Input: [2, 3, 1, 8, 2, 3, 5, 1]
Output: 4, 6, 7
Explanation: The array should have all numbers from 1 to 8, due to duplicates 4, 6, and 7 are missing.
*/
func findMissingNumbers(input []int) []int {
	var result []int
	return result
}
func main() {
	fmt.Println(findMissingNumbers([]int{2, 3, 1, 8, 2, 3, 5, 1}))
	fmt.Println(findMissingNumbers([]int{2, 4, 1, 2}))
	fmt.Println(findMissingNumbers([]int{2, 3, 2, 1}))
}
