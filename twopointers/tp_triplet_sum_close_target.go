package main

import (
	"fmt"
	"math"
	"sort"
)

//Given an array of unsorted numbers and a target number, find a triplet in the array whose sum is as close to the target number as possible,
//return the sum of the triplet. If there are more than one such triplet, return the sum of the triplet with the smallest sum.
// e.g.
// Input: [-2, 0, 1, 2], target=2
// Output: 1
// Explanation: The triplet [-2, 1, 2] has the closest sum to the target.
func tripletSumCloseToTarget(input []int, targetSum int) int {
	sort.Ints(input)
	smallestDiff := math.MaxInt32
	for i := 0; i < len(input)-2; i++ {
		left := i + 1
		right := len(input) - 1
		for left < right {
			targetDiff := targetSum - input[i] - input[left] - input[right]
			// we've found a triplet with an exact sum
			if targetDiff == 0 {
				return targetSum - targetDiff
			}
			if targetDiff < smallestDiff {
				smallestDiff = targetDiff //always save the closest and biggest difference
			}
			// we need a triplet with a bigger sum
			if targetDiff > 0 {
				left++
			} else {
				right--
			}
		}

	}
	return targetSum - abs(smallestDiff)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
Time complexity
Sorting the array will take O(N* logN)O(N∗logN). Overall searchTriplet() will take O(N * logN + N^2), which is asymptotically equivalent to O(N^2)).

Space complexity #
The space complexity of the above algorithm will be O(N) which is required for sorting.
*/
func main() {
	fmt.Println(tripletSumCloseToTarget([]int{-2, 0, 1, 2}, 2))
	fmt.Println(tripletSumCloseToTarget([]int{-3, -1, 1, 2}, 1))
	fmt.Println(tripletSumCloseToTarget([]int{1, 0, 1, 1}, 100))
}
