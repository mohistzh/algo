package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
	Given an array of strings, group anagrams together.

	Note:
	1. All inputs will be in lowercase.
	2. The order of your output does not matter.
*/
func groupAnagrams(input []string) [][]string {
	auxMap := make(map[string][]string)
	var result [][]string
	for i := 0; i < len(input); i++ {
		str := input[i]
		key := sortString(str)
		if val, ok := auxMap[key]; ok {
			val = append(val, str)
			auxMap[key] = val
		} else {
			auxMap[key] = []string{str}
		}
	}

	for _, v := range auxMap {
		result = append(result, v)
	}
	return result
}

/*
	use ascii code calculation instead of sort string
*/
func groupAnagrams2(input []string) [][]string {
	auxMap := make(map[string][]string)
	var result [][]string
	for i := 0; i < len(input); i++ {
		str := input[i]
		// lower case
		letters := make([]byte, 26)
		for j := 0; j < len(str); j++ {
			letters[str[j]-'a']++
		}

		key := string(letters)

		if val, ok := auxMap[key]; ok {
			val = append(val, str)
			auxMap[key] = val
		} else {
			auxMap[key] = []string{str}
		}

	}
	for _, v := range auxMap {
		result = append(result, v)
	}
	return result

}

func sortString(str string) string {
	arr := strings.Split(str, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}

/*
	0. iterate the array one-time
	1. sorting each item
	2. use sorted item as map key to store anagrams

*/
func main() {
	sample1 := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	fmt.Println(groupAnagrams(sample1))
	fmt.Println()
	fmt.Println(groupAnagrams2(sample1))
}
