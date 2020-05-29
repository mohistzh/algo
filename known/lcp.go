// Longest common prefix

package main

import (
	"fmt"
	"math"
	"sort"
)

/*
	Given a set of strings, find the longest common prefix.
*/
func lcp(wordlist []string) string {
	size := len(wordlist)
	// base case
	if size == 0 {
		return ""
	}
	if size == 1 {
		return wordlist[0]
	}
	//sort the array of strings
	sort.Strings(wordlist)
	minimumLength := int(math.Min(float64(len(wordlist[0])), float64(len(wordlist[size-1]))))
	index := 0
	// becasue we sorted array first, so it's only need compare with first and last elements
	for index < minimumLength && wordlist[0][index] == wordlist[size-1][index] {
		index++
	}
	prefix := wordlist[0][0:index]
	return prefix
}

func main() {
	fmt.Println(lcp([]string{"apple", "ape", "april"}))
	fmt.Println(lcp([]string{"flower", "flow", "flight"}))
	fmt.Println(lcp([]string{"dog", "racecar", "car"}))
}
