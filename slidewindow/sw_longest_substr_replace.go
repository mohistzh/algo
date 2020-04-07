package main

import (
	"fmt"
	"math"
)

/*
	Given a string with lowercase letters only,
	if you are allowed to replace no more than ‘k’ letters with any letter,
	find the length of the longest substring having the same letters after replacement.
	e.g.
	Input: String="aabccbb", k=2
	Output: 5
	Explanation: Replace the two 'c' with 'b' to have a longest repeating substring "bbbbb".
*/

func LongestSubStringReplace(input string, k int) int {
	res, windowStart, sameLetters := 0.0, 0, 0.0
	frequencyMap := make(map[byte]float64)
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		charRight := input[windowEnd]
		frequencyMap[charRight]++
		sameLetters = math.Max(sameLetters, frequencyMap[charRight])
		if int(windowEnd-windowStart+1-int(sameLetters)) > k {
			leftChar := input[windowStart]
			frequencyMap[leftChar]--
			windowStart++
		}
		res = math.Max(res, float64(windowEnd-windowStart+1))
	}

	return int(res)
}

func main() {
	fmt.Println("Longest substring with same letters after replacement:",
		LongestSubStringReplace("aabccbb", 2))
	fmt.Println("Longest substring with same letters after replacement:",
		LongestSubStringReplace("abbcb", 1))
	fmt.Println("Longest substring with same letters after replacement:",
		LongestSubStringReplace("abccde", 1))
}
