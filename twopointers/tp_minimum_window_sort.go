package main

import (
	"fmt"
	"math"
)

/*
Given an array, find the length of the smallest subarray in it which when sorted will sort the whole array.
*/
func minimumWindowSort(input []int) int {
	low, high := 0, len(input)-1
	// find the first element of the array out of sorting order from the beginning
	for (low < len(input)-1) && input[low] <= input[low+1] {
		low++
	}
	if low == high { // if the array is sorted
		return 0
	}
	// find the first element of array out of sorting order from the end
	for (high > 0) && input[high] >= input[high-1] {
		high--
	}
	subArrayMax, subArrayMin := math.MinInt32, math.MaxInt32
	// find the maximum and minimum elements of the array
	for i := low; i <= high; i++ {
		subArrayMin = int(math.Min(float64(subArrayMin), float64(input[i])))
		subArrayMax = int(math.Max(float64(subArrayMax), float64(input[i])))
	}
	// extend the subarray to include any number which is bigger than the minimum of the subarray
	for low > 0 && input[low-1] > subArrayMin {
		low--
	}
	// extend the subarray to include any number which is smaller than the maximum of the subarray
	for high < len(input)-1 && input[high+1] < subArrayMax {
		high++
	}
	fmt.Println("sorting needed", input[low:high+1])
	return high - low + 1

}

func main() {
	fmt.Println(minimumWindowSort([]int{1, 2, 5, 3, 7, 10, 9, 12}))
	fmt.Println(minimumWindowSort([]int{1, 3, 2, 0, -1, 7, 10}))
	fmt.Println(minimumWindowSort([]int{1, 2, 3}))
	fmt.Println(minimumWindowSort([]int{3, 2, 1}))
}
