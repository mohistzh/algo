// Knuth-Morris-Pratt algorithm, also known as 'The needle in a haystack'
package main

import (
	"fmt"
)

/*
In computer science, the Knuth–Morris–Pratt string-searching algorithm (or KMP algorithm)
searches for occurrences of a "word" W within a main "text string" S by employing the observation that when a mismatch occurs,
the word itself embodies sufficient information to determine where the next match could begin, thus bypassing re-examination of previously matched characters.

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
*/
// indexOf S is a text to be searched, W is the word sought, return number of position
func indexOf(S string, W string) int {
	if len(S) < 1 {
		return -1
	}
	if len(W) < 1 {
		return 0
	}
	// indexOfS is the position of current character in S
	// indexOfW is the position of current character in W
	indexOfS, indexOfW, result := 0, 0, 0
	for indexOfS < len(S) {
		if W[indexOfW] == S[indexOfS] {
			indexOfS++
			indexOfW++
			if indexOfW == len(W) {

			}
		} else {

		}
	}

	return result

}

// O(m*n) solution, which m is the length of text, and n is the length of pattern
func bfSolution(text string, pattern string) int {
	if pattern == "" {
		return -1
	}
	textLen := len(text)
	patternLen := len(pattern)

	index := 0

	for index < (textLen - patternLen + 1) {
		match := true
		// internal interval
		for i := 0; i < patternLen; i++ {
			if pattern[i] != text[index+i] {
				match = false
				break
			}
		}
		if match {
			return index
		}
		index++
	}
	return -1
}

func main() {
	text := "doyouseeadoghere"
	pattern := "dog"
	fmt.Println("We found occurrence is:", bfSolution(text, pattern))
}
