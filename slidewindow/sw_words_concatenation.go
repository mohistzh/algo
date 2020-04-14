package main

import (
	"fmt"
)

/*
	Given a string and a list of words, find all the starting indices of substrings in the given string that are a concatenation of
	all the given words exactly once without any overlapping of words. It is given that all words are of the same length.

	e.g.
	Input: String="catfoxcat", Words=["cat", "fox"]
	Output: [0, 3]
	Explanation: The two substring containing both the words are "catfox" & "foxcat".
*/
func findWordConcatenation(input string, words []string) []int {
	wordFrequency := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordFrequency[words[i]]++
	}
	var result []int
	wordsCount := len(words)
	wordLength := len(words[0])
	length := (len(input) - wordsCount*wordLength) + 1
	for i := 0; i < length; i++ {
		wordSeenMap := make(map[string]int)
		for j := 0; j < wordsCount; j++ {
			nextWordIndex := i + j*wordLength
			// Get the next word from the string
			word := input[nextWordIndex : nextWordIndex+wordLength]
			// if we don't need this word, just let it go
			if _, ok := wordFrequency[word]; !ok {
				break
			}
			wordSeenMap[word]++
			// No need to process further if the word has higher frequency than required
			if wordSeenMap[word] > wordFrequency[word] {
				break
			}
			// store index if we have found all the words
			if (j + 1) == wordsCount {
				result = append(result, i)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(findWordConcatenation("catfoxcat", []string{"cat", "fox"}))
	fmt.Println(findWordConcatenation("catcatfoxfox", []string{"cat", "fox"}))
}
