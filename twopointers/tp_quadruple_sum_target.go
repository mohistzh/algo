package main

import (
	"fmt"
	"sort"
)

/*
Given an array of unsorted numbers and a target number,
find all unique quadruplets in it, whose sum is equal to the target number.
e.g.
Input: [4, 1, 2, -1, 1, -3], target=1
Output: [-3, -1, 1, 4], [-3, 1, 1, 2]
Explanation: Both the quadruplets add up to the target.
*/
func quadrupleSumToTarget(input []int, targetSum int) [][]int {
	// w + x + y + z = target
	sort.Ints(input)
	var quadrauplets [][]int
	for i := 0; i < len(input)-3; i++ {
		if i > 0 && input[i] == input[i-1] { // avoid duplicate quadruplets
			continue
		}
		for j := i + 1; j < len(input)-2; j++ { // avoid duplicate quadruplets
			if j > i+1 && input[j] == input[j-1] {
				continue
			}
			findQuadruplets(input, targetSum, i, j, &quadrauplets)
		}

	}
	return quadrauplets

}

// we use two motivable pointers to iterate through the array
func findQuadruplets(input []int, targetSum int, first int, second int, quadrauplets *[][]int) {
	left := second + 1
	right := len(input) - 1
	for left < right {
		currentSum := input[first] + input[second] + input[left] + input[right]
		if currentSum == targetSum {
			*quadrauplets = append(*quadrauplets, []int{input[first], input[second], input[left], input[right]})
			left++
			right--
			for left < right && input[left] == input[left-1] { // skip the same element to avoid duplicate quadruplets
				left++
			}
			for left < right && input[right] == input[right+1] { // skip the same element to avoid duplicate quadruplets
				right--
			}
		} else if currentSum > targetSum {
			right--
		} else {
			left++
		}
	}
}

func main() {
	fmt.Println(quadrupleSumToTarget([]int{4, 1, 2, -1, 1, -3}, 1))
	fmt.Println(quadrupleSumToTarget([]int{2, 0, -1, 1, -2, 2}, 2))

}
