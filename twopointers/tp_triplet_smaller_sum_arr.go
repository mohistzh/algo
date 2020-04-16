package main

import (
	"fmt"
	"sort"
)

/*
Given an array arr of unsorted numbers and a target sum, count all triplets in it such that arr[i] + arr[j] + arr[k] < target where i, j, and k are three different indices.
 Write a function to return the list of all such triplets instead of the count.
*/
func tripletWithSmallerSumArr(input []int, targetSum int) [][]int {
	sort.Ints(input)
	var triplets [][]int
	for i := 0; i < len(input)-2; i++ {
		// X + Y + Z < TargetSum == Y + Z == TargetSum - X
		fillinTriplet(input, targetSum-input[i], i, &triplets)
	}
	return triplets
}
func fillinTriplet(input []int, targetSum int, index int, triplets *[][]int) {
	left, right := index+1, len(input)-1
	for left < right {
		if input[left]+input[right] < targetSum {
			// iterater through backwards
			for i := right; i > left; i-- {
				*triplets = append(*triplets, []int{input[index], input[left], input[i]})
			}
			left++
		} else {
			right--
		}
	}
}

func main() {
	fmt.Println(tripletWithSmallerSumArr([]int{-1, 0, 2, 3}, 3))
	fmt.Println(tripletWithSmallerSumArr([]int{-1, 4, 2, 1, 3}, 5))
}
