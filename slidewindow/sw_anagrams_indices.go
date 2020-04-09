package main

import (
	"fmt"
)

/*
	Given a string and a pattern, find all anagrams of the pattern in the given string.
	Write a function to return a list of starting indices of the anagrams of the pattern in the given string.
*/
func findStringAnagrams(input string, pattern string) []int {
	var result []int
	charFrequency := make(map[byte]int)
	matches, windowStart := 0, 0

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
			result = append(result, windowStart)
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
	return result

}

func main() {
	fmt.Println("Indices of Anagrams(abbcabc) of the pattern(abc) is", findStringAnagrams("abbcabc", "abc"))
	fmt.Println("Indices of Anagrams(ppqp) of the pattern(pq) is", findStringAnagrams("ppqp", "pq"))

}
