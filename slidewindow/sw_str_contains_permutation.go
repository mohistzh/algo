package main

import (
	"fmt"
)

/*
	Given a string and a pattern, find out if the string contains any permutation of the pattern.
*/
func findPermutation(input string, pattern string) bool {
	windowStart, matches, charFrequency := 0, 0, make(map[byte]int)

	for i := 0; i < len(pattern); i++ {
		charFrequency[pattern[i]]++
	}
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		rightChar := input[windowEnd]
		if _, ok := charFrequency[rightChar]; ok {
			charFrequency[rightChar]--
			if charFrequency[rightChar] == 0 {
				matches++
			}
		}
		if matches == len(charFrequency) {
			return true
		}
		if windowEnd >= len(pattern)-1 {
			leftChar := input[windowStart]
			windowStart++
			if _, ok := charFrequency[leftChar]; ok {
				if charFrequency[leftChar] == 0 {
					matches--
				}
				charFrequency[leftChar]++
			}
		}

	}
	return false
}

func main() {
	fmt.Println("Check if the string oidbcaf contains any permutation of the pattern abc:", findPermutation("oidbcaf", "abc"))
	fmt.Println("Check if the string odicf contains any permutation of the pattern dc:", findPermutation("odicf", "dc"))
}
