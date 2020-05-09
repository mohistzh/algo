package main

import (
	"fmt"
)

/*
 We are given an unsorted array containing ‘n’ numbers taken from the range 1 to ‘n’.
 The array originally contained all the numbers from 1 to ‘n’,
 but due to a data error, one of the numbers got duplicated which also resulted in one number going missing.
 Find both these numbers.
*/
func findCorruptPair(input []int) []int {
	var pair []int
	i, n := 0, len(input)
	// swapping items to the relatively correct places
	for i < n {
		j := input[i] - 1
		if input[i] != input[j] {
			input[i], input[j] = input[j], input[i]
		} else {
			i++
		}
	}
	// finding out missing numbers and duplicate numbers
	for i := 0; i < n; i++ {
		if i+1 != input[i] {
			pair = append(pair, input[i], i+1)
		}
	}
	return pair
}

func main() {
	fmt.Println(findCorruptPair([]int{3, 1, 2, 5, 2}))
	fmt.Println(findCorruptPair([]int{3, 1, 2, 3, 6, 4}))
}
