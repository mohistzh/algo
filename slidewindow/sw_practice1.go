package main

import (
	"fmt"
)

/*
	Given an array, find the average of all contiguous subarrays of size ‘K’ in it.
	O(N) time complexity
*/

func solution(K int, input []int) []float64 {
	if K > len(input) {
		return nil
	}
	result := make([]float64, 0)
	windowSum, windowStart := 0.0, 0
	for i := 0; i < len(input); i++ {
		windowSum += float64(input[i])
		if i >= (K - 1) {
			result = append(result, float64(windowSum/float64(K)))
			windowSum -= float64(input[windowStart])
			windowStart++
		}
	}
	return result
}

func main() {
	K := 5
	input := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	result := solution(K, input)
	fmt.Println(result)
}
