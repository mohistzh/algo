package main

import (
	"fmt"
)

/*
Given two strings containing backspaces (identified by the character ‘#’), check if the two strings are equal.
e.g.
Input: str1="xp#", str2="xyz##"
Output: true
Explanation: After applying backspaces the strings become "x" and "x" respectively.
In "xyz##", the first '#' removes the character 'z' and the second '#' removes the character 'y'.
*/
func compareStringsContainsBackspaces(first string, second string) bool {
	firstIndex, secondIndex := len(first)-1, len(second)-1
	for firstIndex > 0 || secondIndex > 0 {
		i1 := getNextValidIndexOfString(first, firstIndex)
		i2 := getNextValidIndexOfString(second, secondIndex)
		if i1 < 0 && i2 < 0 { // both indices reached the end of strings
			return true
		}
		if i1 < 0 || i2 < 0 { // reached the end of one of the strings
			return false
		}
		if first[i1] != second[i2] { // characters are not equal
			return false
		}
		firstIndex = i1 - 1
		secondIndex = i2 - 1
	}

	return true
}

// detecting backspace and to find a valid index of the input string
func getNextValidIndexOfString(input string, index int) int {
	backSpacesCount := 0
	for index >= 0 {
		if input[index] == '#' { // found a backspace
			backSpacesCount++
		} else if backSpacesCount > 0 { // index before backspace symbol is not a valid index
			backSpacesCount--
		} else {
			break
		}
		index--
	}
	return index
}

func main() {
	fmt.Println(compareStringsContainsBackspaces("xy#z", "xzz#"))
	fmt.Println(compareStringsContainsBackspaces("xy#z", "xyz#"))
	fmt.Println(compareStringsContainsBackspaces("xp#", "xyz##"))
	fmt.Println(compareStringsContainsBackspaces("xywrrmp", "xywrrmu#p"))
}
