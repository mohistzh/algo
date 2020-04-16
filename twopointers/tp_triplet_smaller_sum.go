package main

import (
	"fmt"
	"sort"
)

/*
Given an array arr of unsorted numbers and a target sum, count all triplets in it such that arr[i] + arr[j] + arr[k] < target where i, j,
and k are three different indices. Write a function to return the count of such triplets.

	X + Y + Z < targetSum equals to Y + Z == targetSum - X
*/
func tripletWithSmallerSum(input []int, targetSum int) int {
	sort.Ints(input)
	count := 0
	for i := 0; i < len(input)-2; i++ {
		count += findPair(input, targetSum-input[i], i)
	}
	return count
}
func findPair(input []int, targetSum int, index int) int {
	count := 0
	left, right := index+1, len(input)-1
	for left < right {
		if input[left]+input[right] < targetSum { // found the triplet
			// since input[right] >= input[left], thereforce, we can replace input[right] by any numberr between left and right
			// to get a sum less than the target sum
			count += right - left
			left++
		} else {
			right-- // we need a pair with a smaller sum
		}
	}
	return count
}

func main() {
	fmt.Println(tripletWithSmallerSum([]int{-1, 0, 2, 3}, 3))
	fmt.Println(tripletWithSmallerSum([]int{-1, 4, 2, 1, 3}, 5))
}
