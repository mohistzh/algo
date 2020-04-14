package main

import (
	"fmt"
	"math"
)

/*
	Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).
	e.g.
	Input: S = "ADOBECODEBANC", T = "ABC"
	Output: "BANC"

*/

func minimumWindowSubString(input string, pattern string) string {
	windowStart, letterMissing := 0, len(pattern)
	window := []int{math.MinInt16, math.MaxInt16}
	letterSeenMap, letterNeededMap := make(map[byte]int), make(map[byte]int)
	result := ""
	for i := 0; i < len(pattern); i++ {
		letterSeenMap[pattern[i]] = 0
		letterNeededMap[pattern[i]]++
	}
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		rightChar := input[windowEnd]
		if _, ok := letterNeededMap[rightChar]; ok {
			letterSeenMap[rightChar]++
			if letterNeededMap[rightChar] >= letterSeenMap[rightChar] {
				letterMissing--
			}
		}
		for letterMissing == 0 {
			if (windowEnd - windowStart) < (window[1] - window[0]) {
				window[0] = windowStart
				window[1] = windowEnd
			}
			leftChar := input[windowStart]
			if _, ok := letterNeededMap[leftChar]; ok {
				letterSeenMap[leftChar]--
				if letterSeenMap[leftChar] < letterNeededMap[leftChar] {
					letterMissing++
				}
			}
			windowStart++
		}
	}
	if window[0] > -1 {
		result = input[window[0] : window[1]+1]
	}
	return result

}
func main() {
	fmt.Println("Minimum window in 'ADOBECODEBANC' contains all characters in 'ABC':", minimumWindowSubString("ADOBECODEBANC", "ABC"))
	fmt.Println("Minimum window in 'aa' contains all characters in 'aa':", minimumWindowSubString("aa", "aa"))
}
