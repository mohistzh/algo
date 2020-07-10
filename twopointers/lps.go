// Longest Palindromic Substring
package main

import "fmt"

/**
	O(n^2) time complexity
	Iterate whole array token O(n)
	Separing two pointers cost O(n)
**/
func lpsTwoPointers(input string) string {
	var result string
	for i := 0; i < len(input); i++ {
		oddString := palindrome(input, i, i)
		evenString := palindrome(input, i, i+1)

		if len(result) < len(oddString) {
			result = oddString
		}
		if len(result) < len(evenString) {
			result = evenString
		}
	}
	return result
}

func palindrome(input string, left int, right int) string {
	for left >= 0 && right < len(input) && input[left] == input[right] {
		left--
		right++
	}
	return input[left+1 : right]

}

func main() {
	fmt.Println(lpsTwoPointers("abba"))
	fmt.Println(lpsTwoPointers("cbbd"))
}
