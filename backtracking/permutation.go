package main

import (
	"fmt"
)

func Permute(input string, left int, right int) {
	if left == right {
		fmt.Println(input)
	} else {
		for i := left; i <= right; i++ {
			input = swap(input, left, i)
			Permute(input, left+1, right)
			input = swap(input, left, i)
		}
	}
}

func swap(str string, i int, j int) string {
	alph := ([]byte)(str)
	tmp := alph[i]
	alph[i] = alph[j]
	alph[j] = tmp
	return string(alph)

}

func main() {
	str := "abc"
	Permute(str, 0, len(str)-1)
}
