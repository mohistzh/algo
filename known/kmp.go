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

/*
	Use aux[] to store the index. Unlike naive solution, where we shift the world W by one and
	compare all characters at each shift, we use a value from aux[] to dicide the next characters
	to be matched. No need to match characters that we know will match anyway.
*/
func createAux(W string) []int {
	aux := make([]int, len(W))
	// current index to be processed
	index := 1
	// index of the cell before which prefix is equal to the suffix
	m := 0
	for index < len(W) {
		if W[index] == W[m] {
			m++
			aux[index] = m
			index++
		} else if W[index] != W[m] && m != 0 {
			m = aux[m-1]
		} else {
			aux[index] = 0
			index++
		}
	}
	return aux
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
