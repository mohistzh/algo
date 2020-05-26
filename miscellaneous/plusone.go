package main

import (
	"fmt"
)

/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.
*/
func plusOne(nums []int) []int {
	if nums == nil || len(nums) < 1 {
		return nil
	}
	index := 0
	for index = len(nums) - 1; index >= 0; index-- {
		// e.g. [9, 9]
		if nums[index] == 9 {
			nums[index] = 0
		} else {
			nums[index]++
			break
		}
	}
	if index < 0 {
		result := make([]int, len(nums)+1)
		result[0] = 1
		return result
	} else {
		return nums
	}
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9}))

}
