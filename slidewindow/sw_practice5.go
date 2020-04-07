package main

import (
	"fmt"
	"math"
)

/**
Given a string, find the length of the longest substring which has no repeating characters.

e.g.
Input: String="aabccbb"
Output: 3
Explanation: The longest substring without any repeating characters is "abc".
**/
func NonRepeatSubString(input string) int {
	res, windowStart := 0.0, 0.0
	charIndexMap := make(map[byte]int)

	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		charRight := input[windowEnd]
		if _, ok := charIndexMap[charRight]; ok {
			windowStart = math.Max(windowStart, float64(charIndexMap[charRight]+1))
		}
		charIndexMap[charRight] = windowEnd

		res = math.Max(res, float64(windowEnd)-windowStart+1.0)
	}
	return int(res)

}

func main() {
	fmt.Println("Length of the longest substring is", NonRepeatSubString("aabccbb"))
	fmt.Println("Length of the longest substring is", NonRepeatSubString("abbbb"))
	fmt.Println("Length of the longest substring is", NonRepeatSubString("abccde"))

}
