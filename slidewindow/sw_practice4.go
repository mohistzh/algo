package main

import (
	"fmt"
	"math"
)

/*
Given a string, find the length of the longest substring in it with no more than K distinct characters.
*/

func LongestSubStringKDistinct(str string, k int) int {
	frequency := make(map[byte]int)
	res, windowStart := 0.0, 0
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChart := str[windowEnd]
		if _, ok := frequency[rightChart]; ok {
			frequency[rightChart] = 0
		}
		frequency[rightChart]++
		for len(frequency) > k {
			leftChar := str[windowStart]
			frequency[leftChar]--
			if frequency[leftChar] == 0 {
				delete(frequency, leftChar)
			}
			windowStart++
		}
		res = math.Max(res, float64(windowEnd-windowStart+1))
	}
	return int(res)
}

func main() {
	fmt.Println("Longest SubString with K distinct", LongestSubStringKDistinct("araaci", 2))
	fmt.Println("Longest SubString with K distinct", LongestSubStringKDistinct("araaci", 1))
	fmt.Println("Longest SubString with K distinct", LongestSubStringKDistinct("cbbebi", 3))
}
