package main

import (
	"fmt"
)

/*
	Given a string and a pattern, find the smallest substring in the given string which has all the characters of the given pattern.
*/
func smallestWindowContainsSubString(input string, pattern string) string {
	windowStart, matched, substrStart, res := 0, 0, 0, ""
	minLength := len(input) + 1
	charFrequency := make(map[byte]int)

	for i := 0; i < len(pattern); i++ {
		charFrequency[pattern[i]]++
	}

	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		rightChar := input[windowEnd]
		if _, ok := charFrequency[rightChar]; ok {
			charFrequency[rightChar]--
			if charFrequency[rightChar] >= 0 {
				matched++
			}
		}

		for matched == len(pattern) {
			if minLength > (windowEnd - windowStart + 1) {
				minLength = windowEnd - windowStart + 1
				substrStart = windowStart
			}
			leftChar := input[windowStart]
			windowStart++

			if _, ok := charFrequency[leftChar]; ok {
				if charFrequency[leftChar] == 0 {
					matched--
				}
				charFrequency[leftChar]++
			}
		}
	}
	if minLength <= len(input) {
		res = input[substrStart : substrStart+minLength]
	}
	return res
}

func main() {
	fmt.Println("Smallest Window Containing SubString", smallestWindowContainsSubString("aabdec", "abc"))
	fmt.Println("Smallest Window Containing SubString", smallestWindowContainsSubString("abdabca", "abc"))
	fmt.Println("Smallest Window Containing SubString", smallestWindowContainsSubString("adcad", "abc"))
}
