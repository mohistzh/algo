package main

import (
	"fmt"
)

/*
Given an array with positive numbers and a target number, find all of its contiguous subarrays whose product is less than the target number.
*/
func subArrayWithProductLessThanTarget(input []int, target int) [][]int {
	var result [][]int
	product, left := 1, 0
	for right := 0; right < len(input); right++ {
		product *= input[right]
		for product >= target && left < len(input) {
			product /= input[left]
			left++
		}
		var tmpArray []int
		for i := right; i >= left; i-- {
			tmpArray = append(tmpArray, input[i])
			result = append(result, tmpArray)
		}
	}
	return result
}

func main() {
	fmt.Println("Target is 30, subarray is", subArrayWithProductLessThanTarget([]int{2, 5, 3, 10}, 30))
	fmt.Println("Target is 50, subarray is", subArrayWithProductLessThanTarget([]int{8, 2, 6, 5}, 50))
}
